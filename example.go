package main

import (
	"fmt"
	"net/http"
	"log"
	"github.com/gin-gonic/gin"
	"os"
	"io"
)

type Login struct {
	User	 string `form:"user" json:"user" xml:"user" 	binding:"required"`
	Password string `form:"password" json:"password" xml:"password"		binding:"required"`
}

func main() {
	// tạo log file
	gin.DisableConsoleColor()
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	router := gin.Default()

	router.POST("/loginJSON", func(c *gin.Context)	{
		var json Login
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if json.User != "manu" || json.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	router.GET("/user/:name/*action", func(d *gin.Context) {
		name := d.Param("name")
		action := d.Param("action")
		message := name + " is " + action
		d.String(http.StatusOK, message)
	})

	router.POST("/user/:name/*action", func(c *gin.Context) {
		// c.FullPath() == "/user/:name/*action" // true
		fmt.Println(c.FullPath())
	})

	router.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.DefaultQuery("lastname", "Thanh")
		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})

	router.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "Yến") // Báo các biến đầu vào từ cái Form gửi lên nếu không có gì
		// giá trị sẽ là yến

		c.JSON(200, gin.H{
			"status": "posted",
			"message": message,
			"nick": nick,
		})
	})

	router.POST("/post", func(c *gin.Context) {
		id := c.Query("id") // cái này là ở trên địa chỉ uri
		page := c.DefaultQuery("page", "Querry1") // cái này cũng thế
		name := c.PostForm("name")
		message := c.PostForm("message")

		fmt.Printf("id: %s; page: %s; name: %s; message: %s \n", id, page, name, message)
	})

	router.POST("/post2", func(c *gin.Context) {
		ids := c.QueryMap("ids")
		names := c.PostFormMap("names")

		fmt.Printf("ids: %v, names: %s\n", ids, names)
		// lệnh http://localhost:8080/post2?ids[a]=1234&ids[b]=hello cùng với đó là names post kiểu form url-encoded
		// lên
	})

	router.POST("/upload", func(c *gin.Context){
		file, _ := c.FormFile("file")

		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})

	router.POST("/multi_upload", func(c *gin.Context){
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]

		for _,file := range files {
			log.Println(file.Filename)

		}
		c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
	})

	
	router.Run()
}
