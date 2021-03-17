package middleware

import (
	"net/http"
	"netdisk/cache"
)

func AuthHandler(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(
		func(writer http.ResponseWriter, request *http.Request) {
			request.ParseForm()
			name := request.Form.Get("username")
			session := request.Form.Get("token")
			realSession, ok := cache.GetSession(name)
			if !ok || session != realSession {
				writer.WriteHeader(http.StatusForbidden)
				return
			}
			handlerFunc(writer, request)
		},
	)
}
