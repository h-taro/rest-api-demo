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
	r.POST("/post", router.PostData)

	v1 := r.Group("/v1")
	v1.GET("/ping", router.GetV1Ping)

	v2 := r.Group("/v2")
	v2.GET("/ping", router.GetV2Ping)

	r.GET("/person", router.GetPersonResponse)

	r.Run()
}
