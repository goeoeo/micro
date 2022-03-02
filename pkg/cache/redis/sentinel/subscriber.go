package sentinel

import (
	"context"
	"fmt"
	"time"
	"unsafe"

	"github.com/gomodule/redigo/redis"
	"github.com/sirupsen/logrus"
)

type (
	//订阅回调函数
	SubscribeCallback func(channel, massage string)

	Subscriber struct {
		cancel context.CancelFunc
		client *redis.PubSubConn
		cbMap  map[string]SubscribeCallback
	}
)

//NewSubscriber
func NewSubscriber(ctx context.Context, address string, password string) (s *Subscriber, err error) {
	var (
		conn redis.Conn
	)
	if conn, err = redis.Dial("tcp", address); err != nil {
		return
	}

	//验证redis密码
	if password != "" {
		if _, authErr := conn.Do("AUTH", password); authErr != nil {
			conn.Close()
			return nil, fmt.Errorf("redis auth password error: %s", authErr)
		}
	}

	s = &Subscriber{
		client: &redis.PubSubConn{Conn: conn},
		cbMap:  make(map[string]SubscribeCallback),
	}

	ctx, s.cancel = context.WithCancel(ctx)

	go s.Receive(ctx)

	return

}

func (this *Subscriber) Close() (err error) {
	if this.client == nil {
		return
	}
	if err = this.client.Close(); err != nil {
		return
	}

	//退出协程
	this.cancel()
	return
}

func (this *Subscriber) Subscribe(channel interface{}, cb SubscribeCallback) (err error) {
	if err = this.client.Subscribe(channel); err != nil {
		return
	}

	this.cbMap[channel.(string)] = cb

	return
}

func (this *Subscriber) Receive(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			switch res := this.client.Receive().(type) {
			case redis.Message:
				channel := (*string)(unsafe.Pointer(&res.Channel))
				message := (*string)(unsafe.Pointer(&res.Data))
				this.cbMap[*channel](*channel, *message)
			case error:
				logrus.Infof("sentinel subscriber receive err:%v", res)
				time.Sleep(10 * time.Second)
			}

		}

	}
}
