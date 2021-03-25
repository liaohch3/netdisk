package handler

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"netdisk/cache"
	"netdisk/entity"
	"netdisk/model"
	"netdisk/utils"
	"time"
)

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// GET 请求时返回index页面
		data, err := ioutil.ReadFile("./static/view/signup.html")
		if err != nil {
			io.WriteString(w, fmt.Sprintf("open file fail: %v", err))
			return
		}
		io.WriteString(w, string(data))
		return
	} else if r.Method == http.MethodPost {
		r.ParseForm()
		// todo 这些操作应该放在网关里
		name := r.Form.Get("username")
		passwd := r.Form.Get("password")
		phone := r.Form.Get("phone")
		email := r.Form.Get("email")

		phone = fmt.Sprintf("135-%v", time.Now().Format("15:04:05"))
		email = fmt.Sprintf("abc@cba-%v", time.Now().Format("15:04:05"))

		// todo 校验name, passwd
		err := model.CreateUser(name, passwd, phone, email)
		if err != nil {
			io.WriteString(w, fmt.Sprintf("CreateUser fail, err: %v", err))
			return
		}
		w.Write([]byte("SUCCESS"))
	}
}

func SignInHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// GET 请求时返回index页面
		data, err := ioutil.ReadFile("./static/view/signin.html")
		if err != nil {
			io.WriteString(w, fmt.Sprintf("open file fail: %v", err))
			return
		}
		io.WriteString(w, string(data))
		return
	} else if r.Method == http.MethodPost {
		// 1. 校验用户名密码
		r.ParseForm()
		// todo 这些操作应该放在网关里
		name := r.Form.Get("username")
		passwd := r.Form.Get("password")
		user, err := entity.GetUserByName(name)
		if err != nil {
			io.WriteString(w, fmt.Sprintf("GetUserByName fail, err: %v", err))
			return
		}
		if user.Pwd != passwd {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		// 2. 生成session
		session := utils.GenSession(name)
		cache.UpdateSessionMap(name, session)

		// 3. 重定向到首页
		//w.Write([]byte("http://" + r.Host + "/static/view/home.html"))
		resp := utils.NewSuccessMsg(map[string]interface{}{
			"location": "http://" + r.Host + "/static/view/home.html",
			"username": name,
			"token":    session,
		})
		w.Write(resp.JsonByte())
	}
}

func UserInfoHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	// todo 这些操作应该放在网关里
	name := r.Form.Get("username")
	session := r.Form.Get("token")
	realSession, err := cache.GetSession(name)
	if err != nil || session != realSession {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	user, err := entity.GetUserByName(name)
	if err != nil {
		io.WriteString(w, fmt.Sprintf("GetUserByName fail, err: %v", err))
		return
	}

	resp := utils.NewSuccessMsg(user)
	w.Write(resp.JsonByte())
}
