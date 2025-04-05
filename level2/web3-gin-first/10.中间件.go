package main

// import (
// 	"fmt"
// 	"net/http"
// 	"time"

// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	router := gin.Default()
// 	//注册全局中间件
// 	router.Use(GlobalMiddleware(), TimeMiddleware())
// 	//curl --location --request GET 'http://localhost:8080/v1/hello'
// 	//局部中间件
// 	router.GET("/login", LocalMiddleware(), Login)
// 	router.GET("/hello", Hello)

// 	router.Run(":8080")
// }

// func GlobalMiddleware() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		fmt.Println("全局中间件被执行...")
// 	}
// }

// func LocalMiddleware() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		fmt.Println("局部中间件被执行...")
// 	}
// }

// func TimeMiddleware() gin.HandlerFunc {
// 	return func(context *gin.Context) {
// 		// 记录开始时间
// 		start := time.Now()
// 		// 执行后续调用链
// 		context.Next()
// 		// 计算时间间隔
// 		duration := time.Since(start)
// 		// 输出纳秒，以便观测结果
// 		fmt.Println("请求用时: ", duration.Nanoseconds())
// 	}
// }

// func Login(c *gin.Context) {
// 	c.String(http.StatusOK, "login")

// }

// func Hello(c *gin.Context) {
// 	c.String(http.StatusOK, "hello")

// }
