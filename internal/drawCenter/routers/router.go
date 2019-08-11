package routers

import (
	"drawCenter/internal/drawCenter/routers/api"
	"drawCenter/internal/drawCenter/routers/api/v2"
	"github.com/gin-gonic/gin"
	"net/http"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(Cors())
	router.StaticFS("/dist", http.Dir("web/dist"))
	router.StaticFS("/public", http.Dir("web/public"))
	router.StaticFS("/fonts", http.Dir("web/fonts"))
	router.GET("/ping", api.Ping)
	router.StaticFile("/", "web/dist/index.html")

	apiV2 := router.Group("/api/v2")
	// 为 multipart forms 设置较低的内存限制 (默认是 32 MiB)
	// router.MaxMultipartMemory = 8 << 20  // 8 MiB
	apiV2.POST("/upload", v2.UploadExcel)
	apiV2.POST("/toExcel", v2.ToExcel)
	apiV2.GET("/drawName", v2.GetDrawName)
	return router
}
