package initialize

import (
	"github.com/garyburd/redigo/redis"
	"log"
	"recount/global"
)

func RedisInit() {
	if global.RedisPool == nil {
		global.RedisPool = &redis.Pool{
			MaxIdle:     5, //最大空闲连接数
			MaxActive:   0, //最大链接数，0无限制
			IdleTimeout: 0, //链接不关闭
			Dial: func() (redis.Conn, error) {
				dial, err := redis.Dial("tcp", "101.43.217.23:6379", redis.DialPassword("200312"))
				if err != nil {
					log.Println("redis链接失败", err.Error())
				}
				return dial, err
			},
		}
		global.RedisPool.Get()
	}
}
