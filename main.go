package main

import (
	"recount/initialize"
	"recount/router"
)

func init() {
	initialize.Init()
}
func main() {
	//service.Socket()
	engine := router.GetEngine()
	if err := engine.Run(":8080"); err != nil {
		panic(err)
	}
}
