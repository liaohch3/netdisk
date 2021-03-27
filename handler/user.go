package handler

import (
	"fmt"
	"net/http"
	"netdisk/cache"
	"netdisk/entity"
	"netdisk/model"
	"netdisk/utils"

	"github.com/gin-gonic/gin"
)

func SignUpHandler(c *gin.Context) {
	// GET 请求时返回index页面
	c.HTML(http.StatusOK, "signup.html", nil)
}

func DoSignUpHandler(c *gin.Context) {
	// todo 这些操作应该放在网关里
	name := c.PostForm("username")
	passwd := c.PostForm("password")

	fmt.Printf("name: %v, passwd: %d\n", name, passwd)

	// todo 校验name, passwd
	err := model.CreateUser(name, passwd)
	if err != nil {
		c.String(-1, "CreateUser fail, err: %v", err)
		return
	}

	c.String(0, "SUCCESS")
}

func SignInHandler(c *gin.Context) {
	// GET 请求时返回index页面
	c.HTML(http.StatusOK, "signin.html", nil)
}
func DoSignInHandler(c *gin.Context) {
	// 1. 校验用户名密码
	// todo 这些操作应该放在网关里
	name := c.PostForm("username")
	passwd := c.PostForm("password")
	user, err := entity.GetUserByName(name)
	if err != nil {
		c.String(-1, "GetUserByName fail, err: %v", err)
		return
	}
	if user.Pwd != passwd {
		c.JSONP(http.StatusForbidden, nil)
		return
	}

	// 2. 生成session
	session := utils.GenSession(name)
	cache.UpdateSessionMap(name, session)

	// 3. 重定向到首页
	c.JSON(http.StatusOK, map[string]interface{}{
		"location": "/static/view/home.html",
		"username": name,
		"token":    session,
	})
}

func UserInfoHandler(c *gin.Context) {
	// todo 这些操作应该放在网关里
	name := c.Query("username")
	session := c.Query("token")
	realSession, err := cache.GetSession(name)
	if err != nil || session != realSession {
		c.String(http.StatusForbidden, "")
		return
	}

	user, err := entity.GetUserByName(name)
	if err != nil {
		c.String(-1, "GetUserByName fail, err: %v", err)
		return
	}

	resp := utils.NewSuccessMsg(user)
	c.JSONP(http.StatusOK, resp)
}
