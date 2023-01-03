package main

import (
	"rest-api-demo/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/ping", router.Getting)
	r.GET("/users/:userName", router.GetUserName)
	r.GET("/users/:userName/*action", router.GetUserNameAndAction)
	r.GET("/welcome", router.GetFirstName)
	r.GET("/test", router.GetPerson)

	r.Run()
}
