package main

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	router := gin.Default()
// 	//curl --location --request GET '127.0.0.1:8080/findUser?username=jack&userid=001'
// 	//curl --location --request GET '127.0.0.1:8080/findUser'
// 	router.GET("/findUser", FindUser)
// 	router.Run(":8080")
// }

// func FindUser(c *gin.Context) {
// 	username := c.DefaultQuery("username", "defaultUser")
// 	userid := c.Query("userid")
// 	c.String(http.StatusOK, "user name is %s, userid is %s", username, userid)
// }
