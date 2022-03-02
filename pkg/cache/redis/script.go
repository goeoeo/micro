package redis

import "github.com/gomodule/redigo/redis"

var (
	//设置集合，原子性设置过期时间
	hsetScript = redis.NewScript(1, `
	local exists=redis.call("EXISTS",KEYS[1])
	
	local num=redis.call("HSET",KEYS[1],ARGV[1],ARGV[2])
	if exists 
	then 
       redis.call("EXPIRE",KEYS[1],ARGV[3])
	end
	return num
`)

	//设置有序集合，原子性设置过期时间
	zaddScript = redis.NewScript(1, `
	local exists=redis.call("EXISTS",KEYS[1])
	
	local num=redis.call("ZADD",KEYS[1],ARGV[1],ARGV[2])
	if exists 
	then 
       redis.call("EXPIRE",KEYS[1],ARGV[3])
	end
	return num
`)
)
