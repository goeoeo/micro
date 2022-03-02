package redis

import (
	"time"

	"git.internal.yunify.com/benchmark/benchmark/pkg/cache/redis/normal"
	"git.internal.yunify.com/benchmark/benchmark/pkg/cache/redis/sentinel"

	"git.internal.yunify.com/benchmark/benchmark/pkg/proc"

	"github.com/go-redsync/redsync"
	"github.com/gomodule/redigo/redis"

	"git.internal.yunify.com/benchmark/benchmark/pkg/cache"
)

type Pool interface {
	Init() error
	Get() redis.Conn
	Close()
}

type RedisCli struct {
	name              string
	config            *Config
	pool              Pool             //redigo连接池
	redisMutexFactory *redsync.Redsync //redis锁
}

func NewRedisCli(name string) *RedisCli {
	return &RedisCli{
		name:   name,
		config: new(Config),
	}
}

func (r *RedisCli) Config() (string, interface{}) {
	return r.name, r.config
}

func (r *RedisCli) Init() (err error) {

	switch r.config.Mode {
	case REDIS_MODE_CLUSTER:
		//TODO

	case REDIS_MODE_SENTINEL:
		//哨兵模式
		r.pool = sentinel.New(r.config.MonitorName, r.config.Endpoint()...).WithMasterConfig(&normal.Node{
			Password:    r.config.Password,
			MaxIdle:     r.config.MaxIdle,
			IdleTimeout: r.config.IdleTimeout,
			DbNumber:    r.config.DbNumber,
		})

	default:
		//普通模式
		r.pool = &normal.Node{
			Address:     r.config.EndpointFirst(),
			Password:    r.config.Password,
			MaxIdle:     r.config.MaxIdle,
			IdleTimeout: r.config.IdleTimeout,
			DbNumber:    r.config.DbNumber,
		}

	}

	//初始化连接池
	if err = r.pool.Init(); err != nil {
		return
	}

	//redis锁
	r.redisMutexFactory = redsync.New([]redsync.Pool{r.pool})

	proc.AddShutdownListener(func() {
		r.pool.Close()
	})

	return
}

//redis锁
func (r *RedisCli) NewMutex(key string, options ...redsync.Option) *redsync.Mutex {
	//默认配置
	defaultOptions := []redsync.Option{
		redsync.SetExpiry(time.Second * 30),
		redsync.SetTries(1),
	}
	//传入配置覆盖
	defaultOptions = append(defaultOptions, options...)

	return r.redisMutexFactory.NewMutex(key, defaultOptions...)
}

//cache实现
func (r *RedisCli) Cache() cache.Cache {
	return NewConn(r.pool.Get())
}
