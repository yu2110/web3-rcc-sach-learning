package main

// import "github.com/gin-gonic/gin"

// func main() {
// 	router := gin.Default()

// 	//参数化路由 http://localhost:8080/users/123
// 	router.GET("/users/:id", func(c *gin.Context) {

// 		id := c.Param("id")

// 		c.String(200, "user id: %s", id)

// 	})

// 	//路由组 http://localhost:8080/api/v1/user/123
// 	v1 := router.Group("/api/v1")
// 	{
// 		v1.GET("/users", func(c *gin.Context) {
// 			c.String(200, "list of users")
// 		})

// 		v1.POST("/users", func(c *gin.Context) {
// 			c.String(200, "create user")
// 		})

// 		v1.GET("/user/:id", func(c *gin.Context) {
// 			id := c.Param("id")
// 			c.String(200, "get user by id : %s", id)
// 		})

// 		v1.PUT("/user/:id", func(c *gin.Context) {
// 			id := c.Param("id")
// 			c.String(200, "update user by id : %s", id)
// 		})

// 		v1.DELETE("/user/:id", func(c *gin.Context) {
// 			id := c.Param("id")
// 			c.String(200, "delete user by id : %s", id)
// 		})

// 	}

// 	router.Run(":8080")
// }
