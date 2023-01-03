package router

import (
	"net/http"
	"rest-api-demo/model"

	"github.com/gin-gonic/gin"
)

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

func GetUserNameAndAction(ctx *gin.Context) {
	userName := ctx.Param("userName")
	action := ctx.Param("action")
	ctx.JSON(200, gin.H{
		"userName": userName,
		"action":   action,
	})
}

func GetFirstName(ctx *gin.Context) {
	firstName := ctx.Query("firstName")
	ctx.JSON(200, gin.H{
		"firstName": firstName,
	})
}

func GetPerson(ctx *gin.Context) {
	var person model.Person

	err := ctx.ShouldBindQuery(&person)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"name":    person.Name,
		"address": person.Address,
	})
}
