package main

// import (
// 	"log"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	router := gin.Default()
// 	router.GET("/hello", Hello)
// 	router.Run(":8080")
// }

// func (c *gin.Context) Copy() *gin.Context

// func Hello(c *gin.Context) {
// 	ctx := c.Copy()

// 	go func() {
// 		// 子协程应该使用Context的副本，不应该使用原始Context
// 		log.Println("异步处理函数: ", ctx.HandlerNames())
// 	}()
// 	log.Println("接口处理函数: ", c.HandlerNames())
// 	c.String(http.StatusOK, "hello")

// }
