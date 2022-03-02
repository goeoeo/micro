package cache

type Cache interface {
	Remember(key string, value interface{}, getValue func() (interface{}, error), ignoreRdsErr bool, expireSeconds ...int) (err error)
	Get(key string, value interface{}) (err error)
	Set(key string, value interface{}, expireSeconds ...int) (err error)
	Del(keys ...string) (err error)
	Exists(key string) (bool, error)
	INCR(key string, expires ...int) (num int64, err error)
	HSet(key, filed, value string, expireSeconds ...int) (err error)
	HGet(key, filed string) (value string, err error)
	HDel(key, filed string) (err error)
	//设置有序集合
	ZSet(key, value string, score float64, expireSeconds ...int) (err error)
	//获取有序集合
	ZGet(key string, limit int, asc ...bool) (res []string, err error)
	Close() error
}
