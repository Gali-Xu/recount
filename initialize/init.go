package initialize

import "runtime"

func Init() {
	MongoInit()
	//SocketClientInit()
	RedisInit()
	if runtime.GOOS != "linux" {
		return
	}
}
