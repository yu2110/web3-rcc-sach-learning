package main

// import (
// 	"net/http"
// 	"net/url"

// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	router := gin.Default()
// 	/**
// 		curl --location --request POST 'http://localhost:8080/upload' \
// 	    --form 'file=@"/C:/Users/user/Pictures/Camera Roll/a.jpg"'
// 	*/
// 	router.POST("/upload", uploadFile)
// 	/**
// 		curl --location --request POST 'http://localhost:8080/uploadFiles' \
// 	    --form 'files=@"/C:/Users/Stranger/Pictures/Camera Roll/a.jpg"' \
// 	    --form 'files=@"/C:/Users/Stranger/Pictures/Camera Roll/123.jpg"' \
// 	    --form 'files=@"/C:/Users/Stranger/Pictures/Camera Roll/girl.jpg"'
// 	*/
// 	router.POST("/uploadFiles", uploadFiles)
// 	/**
// 	curl --location --request GET 'http://localhost:8080/download/a.jpg'
// 	*/
// 	router.GET("/download/:filename", downLoad)
// 	router.Run(":8080")
// }

// func uploadFile(c *gin.Context) {

// 	file, err := c.FormFile("file")
// 	if err != nil {
// 		c.String(http.StatusBadRequest, "%+v", err)
// 		return
// 	}

// 	err = c.SaveUploadedFile(file, "./"+file.Filename)
// 	if err != nil {
// 		c.String(http.StatusBadRequest, "%+v", err)
// 		return
// 	}

// 	c.String(http.StatusOK, "upload %s size: %d byte successfully !", file.Filename, file.Size)

// }

// func uploadFiles(c *gin.Context) {
// 	form, _ := c.MultipartForm()
// 	files := form.File["files"]
// 	for _, file := range files {
// 		err := c.SaveUploadedFile(file, "./"+file.Filename)
// 		if err != nil {
// 			c.String(http.StatusBadRequest, "upload failed")
// 			return
// 		}
// 	}

// 	c.String(http.StatusOK, "upload %d files successfully!", len(files))
// }

// func downLoad(c *gin.Context) {
// 	filename := c.Param("filename")
// 	// 方式一
// 	// c.FileAttachment("/", filename)

// 	//方式二
// 	response, request := c.Writer, c.Request

// 	// 写入响应头
// 	// response.Header().Set("Content-Type", "application/octet-stream") 以二进制流传输文件
// 	response.Header().Set("Content-Disposition", `attachment; filename*=UTF-8''`+url.QueryEscape(filename)) // 对文件名进行安全转义
// 	response.Header().Set("Content-Transfer-Encoding", "binary")
// 	http.ServeFile(response, request, filename)
// }
