package v2

import (
	"drawCenter/internal/drawCenter/model"
	"drawCenter/internal/pkg/excel"
	"drawCenter/pkg/file"
	"drawCenter/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
)

const uploadPath = "upload"

func UploadExcel(context *gin.Context) {
	// 单文件
	uploadFile, _ := context.FormFile("file")
	logger.Info(uploadFile.Filename)

	//检查目录
	uploadPathExist, err := file.PathExists(uploadPath)
	if err != nil {
		panic(err)
	}
	if !uploadPathExist {
		file.PathMkDir(uploadPath)
	}
	// 上传文件至指定目录
	excelPath := path.Join(uploadPath, uploadFile.Filename)
	err = context.SaveUploadedFile(uploadFile, excelPath)
	if err != nil {
		panic(err)
	}
	excelData := excel.ParseExcelV2(excelPath)
	var respModel model.RespModel
	respModel.Code = http.StatusOK
	respModel.Msg = "success"
	respModel.Data = excelData
	context.JSON(http.StatusOK, respModel)
}
