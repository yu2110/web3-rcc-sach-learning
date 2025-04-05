package main

// import (
// 	"fmt"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"github.com/gin-gonic/gin/binding"
// )

// type LoginUser struct {
// 	Username string `bind:"required" json:"username" form:"username" uri:"username"`
// 	Password string `bind:"required" json:"password" form:"password" uri:"password"`
// }

// func main() {
// 	router := gin.Default()

// 	/**
// 	Json 数据绑定
// 	curl --location --request POST '127.0.0.1:8080/loginWithJSON' \
// 	--header 'Content-Type: application/json' \
// 	--data-raw '{
// 		"username":"root",
// 		"password":"root"
// 	}'
// 	*/
// 	router.POST("/loginWithJSON", Login)
// 	/**
// 	表单数据绑定
// 	curl --location --request POST '127.0.0.1:8080/loginWithForm' \
// 	--form 'username="root"' \
// 	--form 'password="root"'
// 	*/
// 	router.POST("/loginWithForm", Login)
// 	/**
// 	URL 数据绑定
// 	curl --location --request GET '127.0.0.1:8080/loginWithQuery/root/root'
// 	*/
// 	router.GET("/loginWithQuery/:username/:password", Login)
// 	router.Run(":8080")
// }

// func Login(c *gin.Context) {

// 	var login LoginUser

// 	//当使用 URL 参数时，我们应该手动指定解析方式
// 	//err := c.ShouldBindUri(&login); err == nil
// 	if c.ShouldBind(&login) == nil && login.Password != "" && login.Username != "" {
// 		c.String(http.StatusOK, "login successfully!")

// 	} else {
// 		c.String(http.StatusBadRequest, "login failed!")
// 	}
// 	fmt.Println(login)

// }

// //多次绑定

// type formA struct {
// 	Username string `bind:"required" json:"username" form:"username" uri:"username"`
// 	Password string `bind:"required" json:"password" form:"password" uri:"password"`
// }

// type formB struct {
// 	Username string `bind:"required" json:"username" form:"username" uri:"username"`
// 	Password string `bind:"required" json:"password" form:"password" uri:"password"`
// }

// func SomeHandler(c *gin.Context) {
// 	objA := formA{}
// 	objB := formB{}
// 	// 读取 c.Request.Body 并将结果存入上下文。
// 	if errA := c.ShouldBindBodyWith(&objA, binding.JSON); errA == nil {
// 		c.String(http.StatusOK, `the body should be formA`)
// 		// 这时, 复用存储在上下文中的 body。
// 	}
// 	if errB := c.ShouldBindBodyWith(&objB, binding.JSON); errB == nil {
// 		c.String(http.StatusOK, `the body should be formB JSON`)
// 		// 可以接受其他格式
// 	}
// 	if errB2 := c.ShouldBindBodyWith(&objB, binding.XML); errB2 == nil {
// 		c.String(http.StatusOK, `the body should be formB XML`)
// 	}
// }
