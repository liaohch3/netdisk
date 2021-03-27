package middleware

import (
	"net/http"
	"netdisk/cache"
)

func AuthHandler(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(
		func(writer http.ResponseWriter, request *http.Request) {
			request.ParseForm()
			name := request.Query("username")
			session := request.Query("token")
			realSession, err := cache.GetSession(name)
			if err != nil || session != realSession {
				writer.WriteHeader(http.StatusForbidden)
				return
			}
			handlerFunc(writer, request)
		},
	)
}
