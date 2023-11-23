package main

import (
	"fmt"
	"github.com/daemgo/gopkg/pkg/engine"
	"github.com/gin-gonic/gin"
)

var (
	userResourceType = engine.ResourceType{
		Scope:    engine.ResourceScope("management"),
		Resource: "user",
	}
)

func main() {

	ginEngine := engine.New()
	group := ginEngine.Group("api/v1")

	group.GET("/user", userResourceType, func(ctx *gin.Context) {
		ctx.JSONP(200, "SUCCESS")
	})

	group.PUT("/user/:user_id", userResourceType, func(ctx *gin.Context) {
		userID := ctx.Param("user_id")
		ctx.JSONP(200, fmt.Sprintf("%s + %s", "SUCCESS", userID))
	})

	err := ginEngine.Run(":3000")
	if err != nil {
		panic("Server start failed, please check code.....")
	}

}
