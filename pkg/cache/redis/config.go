package redis

const (
	REDIS_MODE_NORMAL   = "normal"   //普通模式
	REDIS_MODE_SENTINEL = "sentinel" //哨兵模式
	REDIS_MODE_CLUSTER  = "cluster"  //集群模式

)

type Config struct {
	Address     []string //地址
	Mode        string   `default:"normal"` //部署模式
	Password    string   `default:""`       //密码
	MaxIdle     int      `default:"30"`     //最大空闲连接数 单位:秒
	IdleTimeout int      `default:"240"`    //连接存活时间 单位:秒
	DbNumber    int      `default:"6"`      //默认连接库编号

	MonitorName string `default:""` //集群名称，sentinel必须
}

func (c Config) Endpoint() []string {
	if len(c.Address) == 0 {
		return []string{"localhost:6379"}
	}
	return c.Address
}

func (c Config) EndpointFirst() string {
	if len(c.Address) == 0 {
		return "localhost:6379"
	}
	return c.Address[0]
}
