package main

import (
	"netdisk/cache"
	"netdisk/entity"
	router2 "netdisk/service/apigateway/router"
)

func main() {
	initial()
	router := router2.Router()
	err := router.Run(":13848")
	if err != nil {
		panic(err.Error())
	}
}

func initial() {
	entity.InitOrm()
	cache.InitCache()
}
