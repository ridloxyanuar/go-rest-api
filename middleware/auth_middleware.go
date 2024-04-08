package middleware

import (
	"net/http"
	"ridloxyanuar/go-rest-api/helper"
	"ridloxyanuar/go-rest-api/model/web"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (m *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("X-API-Key") != "TES" {
		w.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
			Data:   nil,
		}

		helper.WriteToResponseBody(w, webResponse)
	} else {
		m.Handler.ServeHTTP(w, r)
	}
}
