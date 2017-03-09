package middleware

import (
	"context"
	"net/http"
	"time"

	ae "github.com/mantishK/galore/apperror"
	al "github.com/mantishK/galore/log"
	"github.com/mantishK/galore/model"
	"github.com/mantishK/galore/reply"
)

type Adapter func(http.Handler) http.Handler

func Adapt(hf http.HandlerFunc, adapters ...Adapter) http.Handler {
	var h http.Handler
	for _, adapter := range adapters {
		if h == nil {
			h = adapter(hf)
		} else {
			h = adapter(h)
		}
	}
	return h
}

func AccessLog() Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			al.Access("IP: " + r.RemoteAddr)
			h.ServeHTTP(w, r)
		})
	}
}

func Authorize() Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("X-TOKEN")
			if token == "" {
				reply.Err(w, ae.Forbidden(""))
				return
			}
			userToken := model.UserToken{}
			userToken.Token = token
			err := userToken.GetUserIdFromToken()
			if err != nil || userToken.UserId == 0 {
				reply.Err(w, ae.Forbidden(""))
				return
			}
			now := time.Now()
			timeDifference, err := time.ParseDuration("1680000h")
			if err != nil {
				reply.Err(w, ae.Internal("", err))
				return
			}
			timeDifferencePlusModified := userToken.Modified.Add(timeDifference)
			if now.After(timeDifferencePlusModified) {
				reply.Err(w, ae.TokenExpired(""))
				return
			}
			ctx := context.WithValue(r.Context(), "user_id", userToken.UserId)
			ctx = context.WithValue(ctx, "user_token", userToken.Token)
			h.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
