package main

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	router := gin.Default()
// 	/**
// 	使用form-data
// 	curl --location --request POST '127.0.0.1:8080/register' \
// 	--form 'username="jack"' \
// 	--form 'password="123456"'
// 	*/
// 	router.POST("/register", RegisterUser)
// 	/**
// 	使用json
// 	curl --location --request POST '127.0.0.1:8080/update' \
// 	--header 'Content-Type: application/json' \
// 	--data-raw '{
// 		"username":"username",
// 		"password":"123456"
// 	}'
// 	*/
// 	router.POST("/update", UpdateUser)
// 	router.Run(":8080")
// }

// func RegisterUser(c *gin.Context) {
// 	//PostForm方法默认解析application/x-www-form-urlencoded和
// 	// multipart/form-data类型的表单
// 	username := c.PostForm("username")
// 	password := c.PostForm("password")
// 	c.String(http.StatusOK, " successfully registerd, user name is %s, password is %s", username, password)
// }

// func UpdateUser(c *gin.Context) {
// 	var form map[string]string
// 	c.ShouldBind(&form)
// 	c.String(http.StatusOK, " successfully update, user name is [%s], password is [%s]", form["username"], form["password"])
// }
