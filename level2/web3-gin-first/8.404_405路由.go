package main

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	router := gin.Default()

// 	v1 := router.Group("v1")
// 	{
// 		v1.GET("/hello", Hello)
// 		v1.GET("/login", Login)
// 	}

// 	//注册处理器
// 	//<h1>404 Page Not Found</h1>
// 	router.NoRoute(func(ctx *gin.Context) {
// 		ctx.String(http.StatusNotFound, "<h1>404 Page Not Found</h1>")
// 	})

// 	//curl --location --request OPTIONS 'http://localhost:8080/v2/delete'
// 	router.NoMethod(func(ctx *gin.Context) {
// 		ctx.String(http.StatusMethodNotAllowed, "method not allow")
// 	})

// 	router.Run(":8080")
// }

// func Hello(c *gin.Context) {

// }

// func Login(c *gin.Context) {

// }
