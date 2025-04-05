package main

// import (
// 	"net/http"

// 	"github.com/gin-contrib/sessions"
// 	"github.com/gin-contrib/sessions/cookie"
// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	router := gin.Default()

// 	store := cookie.NewStore([]byte("secret"))

// 	router.Use(sessions.Sessions("mysession", store))
// 	router.GET("/incr", func(c *gin.Context) {
// 		session := sessions.Default(c)
// 		var count int
// 		v := session.Get("count")

// 		if v == nil {
// 			count = 0
// 		} else {
// 			count = v.(int)
// 			count++
// 		}

// 		session.Set("count", count)
// 		session.Save()
// 		c.JSON(200, gin.H{"count": count})
// 	})

// 	router.GET("/hello", Hello)

// 	router.Run(":8080")
// }

// func Hello(c *gin.Context) {
// 	c.String(http.StatusOK, "hello")

// }
