package v2

import (
	"drawCenter/internal/drawCenter/model"
	"drawCenter/pkg/file"
	"drawCenter/pkg/logger"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
)

type App struct {
	DrawName string `yaml:"drawName"`
}

var appConfig = App{}
var defaultConfigFile = "config.yaml"
var defaultConfigContent = `drawName: 美丽中国`

func GetDrawName(context *gin.Context) {
	drawName := getConfigInfo(defaultConfigFile)
	logger.Info("drawName ", drawName)
	var respModel model.RespModel
	respModel.Code = http.StatusOK
	respModel.Msg = "success"
	respModel.Data = drawName
	context.JSON(http.StatusOK, respModel)
}
func getConfigInfo(configPath string) string {
	path := defaultConfigFile
	if len(configPath) > 0 {
		path = configPath
	}
	configExist, _ := file.PathExists(path)
	if !configExist {
		file.WriteFileWithIoUtil(defaultConfigFile, defaultConfigContent)
	}
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		logger.Error("setting.Setup, fail to parse "+defaultConfigFile+": %v", err)
	}
	err = yaml.Unmarshal(yamlFile, &appConfig)
	if err != nil {
		logger.Error(err.Error())
	}
	return appConfig.DrawName
}
