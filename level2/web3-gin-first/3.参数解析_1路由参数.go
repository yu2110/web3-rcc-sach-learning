package main

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	router := gin.Default()

// 	//curl --location --request GET '127.0.0.1:8080/findUser/jack/001'
// 	router.GET("/findUser/:username/:userid", FindUser)
// 	//curl --location --request GET '127.0.0.1:8080/downloadFile/img/fruit.png'
// 	router.GET("/downloadFile/*filepath", UserPage)
// 	router.Run(":8080")
// }

// func FindUser(c *gin.Context) {
// 	username := c.Param("username")
// 	userid := c.Param("userid")
// 	c.String(http.StatusOK, "user name is %s, userid is %s", username, userid)
// }

// func UserPage(c *gin.Context) {
// 	filepath := c.Param("filepath")
// 	c.String(http.StatusOK, "filepath is %s", filepath)
// }
