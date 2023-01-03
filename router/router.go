package router

import "github.com/gin-gonic/gin"

func Getting(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "pang",
	})
}

func GetUserName(ctx *gin.Context) {
	userName := ctx.Param("userName")
	ctx.JSON(200, gin.H{
		"userName": userName,
	})
}
