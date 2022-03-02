package sentinel

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/gomodule/redigo/redis"
	"github.com/sirupsen/logrus"

	"git.internal.yunify.com/benchmark/benchmark/pkg/cache/redis/normal"
)

//redis-sentinel
type Sentinel struct {
	monitor             string         //监听集群名称
	masterAddressChange chan bool      //主服地址
	sentinelNodes       []*normal.Node //连接sentinel的配置
	masterNode          *normal.Node   //连接redis主服务配置
	cancel              context.CancelFunc
	ctx                 context.Context
}

func New(monitor string, address ...string) *Sentinel {
	s := &Sentinel{
		monitor:             monitor,
		masterAddressChange: make(chan bool),
		masterNode:          &normal.Node{},
	}

	for _, v := range address {
		s.setSentinelNodes(&normal.Node{Address: v})
	}

	s.ctx, s.cancel = context.WithCancel(context.TODO())

	return s
}

//WithMasterConfig 配置连接redis主服务参数
func (s *Sentinel) WithMasterConfig(option *normal.Node) *Sentinel {
	s.masterNode = option
	return s
}

//Init Sentinel初始化
func (s *Sentinel) Init() (err error) {
	//切换主库
	if err = s.switchMaster(); err != nil {
		return fmt.Errorf("sentinel init err:%v", err)
	}

	//监听主库切换
	go s.listen()
	return
}

//Close 退出
func (s *Sentinel) Close() {
	if s.cancel != nil {
		s.cancel()
	}

	return
}

//输出连接
func (this *Sentinel) Get() redis.Conn {
	return this.masterNode.Get()
}

//配置详细的节点信息
func (this *Sentinel) setSentinelNodes(nodes ...*normal.Node) *Sentinel {
	readOrCreateNode := func(node *normal.Node) {
		for k := range this.sentinelNodes {
			if this.sentinelNodes[k].Address == node.Address {
				//更新
				this.sentinelNodes[k] = node
				return
			}
		}

		//添加
		this.sentinelNodes = append(this.sentinelNodes, node)
		return
	}

	for k := range nodes {
		readOrCreateNode(nodes[k])
	}

	return this

}

//通过sentinel设置主服务地址
func (s *Sentinel) switchMaster() (err error) {
	var address string
	if s.masterNode == nil {
		s.masterNode = new(normal.Node)
	}

	//获取主库地址
	if address, err = s.masterAddress(); err != nil {
		return
	}

	//地址未变更
	if s.masterNode.Address == address {
		return nil
	}

	//重新初始化连接池
	s.masterNode.Address = address
	if err = s.masterNode.Init(); err != nil {
		return
	}
	return
}

//监听主库切换
func (s *Sentinel) listen() {
	//为每个sentinel client 创建subscribe
	s.createSubscriberForAllClient()

	for {
		select {
		case <-s.ctx.Done():
			return
		case _, ok := <-s.masterAddressChange:
			if !ok {
				return
			}

			//主服切换了
			if err := s.switchMaster(); err != nil {
				log.Println("主服切换了,重新生成客户端：", err)
			}
		}
	}
}

//获取redis主服务地址
func (this *Sentinel) masterAddress() (address string, err error) {

	for _, v := range this.sentinelNodes {
		if address, err = this.getMasterAddrByName(v.Address, v.Password, this.monitor); err == nil {
			return
		}

		logrus.Infof("get master adderss fail [%v,%v]", v.Address, err)
	}

	err = errors.New("sentinel unavailable")

	return
}

//获取主服名称
func (this *Sentinel) getMasterAddrByName(address, password, monitor string) (masterAddress string, err error) {
	var (
		conn  redis.Conn
		reply interface{}
		ss    []interface{}
	)

	if conn, err = redis.Dial("tcp", address); err != nil {
		return
	}

	defer conn.Close()

	//验证redis密码
	if password != "" {
		if _, err = conn.Do("AUTH", password); err != nil {
			return
		}
	}

	if reply, err = conn.Do("SENTINEL", "get-master-addr-by-name", monitor); err != nil {
		return
	}

	if ss, err = redis.Values(reply, err); err != nil {
		return
	}

	masterAddress = fmt.Sprintf("%s:%s", ss[0], ss[1])

	return

}

//为sentinel客户端对应的redis-sentinel 创建一个监听
func (s *Sentinel) createSubscriberForAllClient() {
	for _, node := range s.sentinelNodes {
		subscriber, err := NewSubscriber(s.ctx, node.Address, node.Password)
		if err != nil {
			logrus.Infof("connection sentinel fail err:%v", err)
			continue
		}

		if err := subscriber.Subscribe("+switch-master", func(channel, massage string) {
			monitor, address := s.parseAddress(massage)
			if monitor == s.monitor && s.masterNode.Address != address {
				//主服发生了切换
				s.masterAddressChange <- true
			}
		}); err != nil {
			logrus.Infof("sentinel subscribe fail err:%v", err)
			continue
		}
	}

	return
}

//解析订阅到的主从切换的消息
func (this *Sentinel) parseAddress(msg string) (monitor string, address string) {
	msgArr := strings.Split(msg, " ")

	if len(msgArr) < 3 {
		return
	}

	return msgArr[0], fmt.Sprintf("%s:%s", msgArr[len(msgArr)-2], msgArr[len(msgArr)-1])

}
