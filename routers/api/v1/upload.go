package v1

import (
	"ginExample/model"
	"ginExample/pkg/excel"
	"ginExample/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
)

const uploadPath = "upload"

func UploadExcel(context *gin.Context) {
	// 单文件
	file, _ := context.FormFile("file")
	logger.Info(file.Filename)

	// 上传文件至指定目录
	excelPath := path.Join(uploadPath, file.Filename)
	err := context.SaveUploadedFile(file, excelPath)
	if err != nil {
		panic(err)
	}
	excelData := excel.ParseExcel(excelPath)
	var respModel model.RespModel
	respModel.Code = http.StatusOK
	respModel.Msg = "success"
	respModel.Data = excelData
	context.JSON(http.StatusOK, respModel)
}
