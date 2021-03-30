package main

import (
	"netdisk/cache"
	"netdisk/entity"
	"netdisk/service/user/handler"
	"netdisk/service/user/proto"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/transport/grpc"
)

func main() {
	initial()
	etcd := etcd.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{"127.0.0.1:12379"} //地址
	})
	//创建一个新的服务
	service := micro.NewService(
		micro.Name("go.micro.service.user"),
		micro.Registry(etcd),
		micro.Transport(grpc.NewTransport()), //修改传输协议
	)

	service.Init()

	err := proto.RegisterUserServiceHandler(service.Server(), new(handler.User))
	if err != nil {
		panic(err.Error())
	}

	err = service.Run()
	if err != nil {
		panic(err.Error())
	}
}

func initial() {
	entity.InitOrm()
	cache.InitCache()
}
