package handler

import (
	"context"
	"fmt"
	"net/http"
	"netdisk/service/user/proto"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/transport/grpc"

	"github.com/gin-gonic/gin"
)

var (
	userCli proto.UserService
)

func init() {
	etcd := etcd.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{"127.0.0.1:12379"} //地址
	})
	//创建一个新的服务
	service := micro.NewService(
		micro.Name("go.micro.service.user"),
		micro.Registry(etcd),
		micro.Transport(grpc.NewTransport()), //修改传输协议
	)

	userCli = proto.NewUserService("go.micro.service.user", service.Client())
}

func SignUpHandler(c *gin.Context) {
	// GET 请求时返回index页面
	c.HTML(http.StatusOK, "signup.html", nil)
}

func DoSignUpHandler(c *gin.Context) {
	name := c.PostForm("username")
	passwd := c.PostForm("password")
	ctx := context.Background()

	resp, err := userCli.SignUp(ctx, &proto.SignUpReq{
		UserName: name,
		Password: passwd,
	})
	if err != nil {
		fmt.Printf("err: %v\n", err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    resp.Resp.StatusCode,
		"message": resp.Resp.StatusMessage,
	})
}
