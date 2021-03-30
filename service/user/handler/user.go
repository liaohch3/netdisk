package handler

import (
	"context"
	"netdisk/cache"
	"netdisk/entity"
	"netdisk/model"
	"netdisk/service/user/proto"
	"netdisk/utils"
	"time"
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
		return nil
	}

	resp.Resp = &proto.BaseResp{
		StatusCode:    0,
		StatusMessage: "success",
	}
	return nil
}

// todo 确认是返回error还是nil
func (u *User) SignIn(ctx context.Context, req *proto.SignInReq, resp *proto.SignInResp) error {
	// 1. 校验用户名密码
	user, err := entity.GetUserByName(req.UserName)
	if err != nil {
		resp.Resp = &proto.BaseResp{
			StatusCode:    -1,
			StatusMessage: "GetUserByName fail",
		}
		return nil
	}
	if user.Pwd != req.Password {
		resp.Resp = &proto.BaseResp{
			StatusCode:    -1,
			StatusMessage: "password check fail, invalid user",
		}
		return nil
	}

	// 2. 生成session
	session := utils.GenSession(req.UserName)
	err = cache.UpdateSessionMap(session, user.Id, 30*24*60*60*time.Second)
	if err != nil {
		resp.Resp = &proto.BaseResp{
			StatusCode:    -1,
			StatusMessage: "UpdateSessionMap fail",
		}
		return nil
	}

	resp.Location = "/static/view/home.html"
	resp.UserName = req.UserName
	resp.Session = session
	resp.Resp = &proto.BaseResp{
		StatusCode:    0,
		StatusMessage: "success",
	}
	return nil
}

func (u *User) GetUserInfo(ctx context.Context, req *proto.GetUserInfoReq, resp *proto.GetUserInfoResp) error {
	user, err := entity.GetUserByUserID(req.UserID)
	if err != nil {
		resp.Resp = &proto.BaseResp{
			StatusCode:    -1,
			StatusMessage: "GetUserByUserID fail",
		}
		return nil
	}

	resp.UserInfo = &proto.User{
		ID:     user.Id,
		Name:   user.Name,
		Status: int32(user.UserStatus),
	}
	resp.Resp = &proto.BaseResp{
		StatusCode:    0,
		StatusMessage: "success",
	}
	return nil
}
