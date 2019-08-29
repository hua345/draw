package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/browser"
	"net/http"
)

func main() {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.StaticFS("/dist", http.Dir("web/dist"))
	router.StaticFS("/public", http.Dir("web/public"))
	router.StaticFS("/lang-data", http.Dir("web/lang-data"))
	router.StaticFS("/tesseract", http.Dir("web/tesseract"))

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.StaticFile("/", "web/dist/index.html")
	// 打开浏览器
	const url = "http://127.0.0.1:9000"
	err := browser.OpenURL(url)
	if err != nil {
		panic(err)
	}
	// default listen and serve on 0.0.0.0:8080
	err = router.Run(":9000")
	if err != nil {
		panic(err)
	}
}
