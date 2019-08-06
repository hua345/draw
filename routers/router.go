package routers

import (
	"ginExample/routers/api"
	"ginExample/routers/api/v1"
	"ginExample/routers/api/v2"
	"github.com/gin-gonic/gin"
	"net/http"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(Cors())
	router.StaticFS("/dist", http.Dir("dist"))
	router.StaticFS("/public", http.Dir("public"))
	router.StaticFS("/fonts", http.Dir("fonts"))
	router.GET("/ping", api.Ping)
	router.StaticFile("/", "dist/index.html")
	apiV1 := router.Group("/api/v1")
	// 为 multipart forms 设置较低的内存限制 (默认是 32 MiB)
	// router.MaxMultipartMemory = 8 << 20  // 8 MiB
	apiV1.POST("/upload", v1.UploadExcel)
	apiV1.POST("/toExcel", v1.ToExcel)

	apiV2 := router.Group("/api/v2")
	// 为 multipart forms 设置较低的内存限制 (默认是 32 MiB)
	// router.MaxMultipartMemory = 8 << 20  // 8 MiB
	apiV2.POST("/upload", v2.UploadExcel)
	apiV2.POST("/toExcel", v2.ToExcel)
	apiV2.GET("/drawName", v2.GetDrawName)
	return router
}
