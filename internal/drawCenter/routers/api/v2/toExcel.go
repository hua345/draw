package v2

import (
	"drawCenter/internal/drawCenter/model"
	"drawCenter/internal/pkg/excel"
	"drawCenter/pkg/file"
	"drawCenter/pkg/logger"
	"github.com/gin-gonic/gin"
	"path"
)

func ToExcel(context *gin.Context) {
	var drawResult model.DrawResult
	err := context.BindJSON(&drawResult)
	if err != nil {
		logger.Error(err)
		context.JSON(200, gin.H{"errcode": 400, "description": "Post Data Err"})
		return
	} else {
		//检查目录
		uploadPathExist, err := file.PathExists(uploadPath)
		if err != nil {
			panic(err)
		}
		if !uploadPathExist {
			file.PathMkDir(uploadPath)
		}
		excel.CreateExcelV2(&drawResult, uploadPath)

		context.Header("content-disposition", `attachment; filename=`+"drawResult.xlsx")
		context.Writer.Header().Add("Content-Type", "application/octet-stream")
		context.Header("Content-Transfer-Encoding", "binary")
		context.File(path.Join(uploadPath, excel.DrawResultExcelV2))
		//context.Header("Content-Type", "application/octet-stream")
		//context.Header("Content-Disposition", "attachment; filename="+"drawResult.xlsx")
		//context.Header("Content-Transfer-Encoding", "binary")
		//excelFile, err := excelize.OpenFile(excel.DrawResultExcelV2)
		//if err != nil {
		//	logger.Error(err)
		//}
		//_ = excelFile.Write(context.Writer)
	}
}
