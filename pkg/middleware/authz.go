package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
)

type Auther interface {
	Authorize(sub, obj, act string) (bool, error)
}

func Authz(a Auther) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sub := ctx.GetString("user_id")
		obj := ctx.Request.URL.Path
		act := ctx.Request.Method
		log.Printf("sub: %s, obj: %s, act: %s", sub, obj, act)
		allow, err := a.Authorize(sub, obj, act)
		if err != nil {
			log.Printf("authorize failed: %s", err)
			ctx.AbortWithStatus(403)
			return
		}
		if !allow {
			log.Printf("authorize failed: %s", err)
			ctx.AbortWithStatus(403)
			return
		}
		ctx.Next()
	}
}
