package normal

import (
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/sirupsen/logrus"
)

//节点配置
type Node struct {
	Address     string      //节点地址 127.0.0.1:6379
	Password    string      //密码
	MaxIdle     int         //最大空闲连接数
	MaxActive   int         //最大的激活连接数，同时最多有N个连接：0=不限制连接数
	IdleTimeout int         //空闲连接等待时间，超过此时间后，空闲连接将被关闭：0=超时也不断开空闲连接
	DbNumber    int         //连接库编号
	pool        *redis.Pool //连接池
}

//Init 初始化连接池
func (n *Node) Init() (err error) {
	var (
		pw, dbnumber redis.DialOption
		conn         redis.Conn
	)

	logrus.Infof("connect redis ...")
	defer func() {
		if err == nil {
			logrus.Infof("connect redis success [%s]", n.Address)
		}
	}()

	pw = redis.DialPassword(n.Password)
	dbnumber = redis.DialDatabase(n.DbNumber)

	//检查
	if conn, err = redis.Dial("tcp", n.Address, pw, dbnumber); err != nil {
		return
	}
	defer conn.Close()

	//连接池
	n.pool = &redis.Pool{
		MaxIdle:     n.MaxIdle,
		IdleTimeout: time.Duration(n.IdleTimeout) * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", n.Address, pw, dbnumber)
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}

	return
}

//输出连接
func (n *Node) Get() redis.Conn {
	return n.pool.Get()
}

//Close
func (n *Node) Close() {
	if n.pool != nil {
		n.pool.Close()
	}
}
