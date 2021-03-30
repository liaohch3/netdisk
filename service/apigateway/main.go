package main

import router2 "netdisk/service/apigateway/router"

func main() {
	router := router2.Router()
	err := router.Run(":13848")
	if err != nil {
		panic(err.Error())
	}
}
