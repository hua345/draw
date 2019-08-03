package main

import (
	"ginExample/pkg/logger"
	"ginExample/routers"
	"github.com/pkg/browser"
)

const LogPath = "./zap.log"

func main() {
	logger.InitLog("info", LogPath)
	routersInit := routers.InitRouter()

	// 打开浏览器
	const url = "http://127.0.0.1:9000"
	err := browser.OpenURL(url)
	if err != nil {
		panic(err)
	}
	// default listen and serve on 0.0.0.0:8080
	err = routersInit.Run(":9000")
	if err != nil {
		panic(err)
	}

}
