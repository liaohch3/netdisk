package handler

import (
	"context"
	"netdisk/model"
	"netdisk/service/user/proto"
)

type User struct{}

func (u *User) SignUp(ctx context.Context, req *proto.SignUpReq, resp *proto.SignUpResp) error {
	// todo 校验name, passwd
	err := model.CreateUser(req.UserName, req.Password)
	if err != nil {
		resp.Resp = &proto.BaseResp{
			StatusCode:    -1,
			StatusMessage: "fail to create user",
		}
		return err
	}

	resp.Resp = &proto.BaseResp{
		StatusCode:    0,
		StatusMessage: "success",
	}
	return nil
}
