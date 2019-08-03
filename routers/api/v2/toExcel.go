package v2

import (
	"ginExample/model"
	"ginExample/pkg/excel"
	"ginExample/pkg/logger"
	"github.com/gin-gonic/gin"
)

func ToExcel(context *gin.Context) {
	var drawResult model.DrawResult
	err := context.BindJSON(&drawResult)
	if err != nil {
		logger.Error(err)
		context.JSON(200, gin.H{"errcode": 400, "description": "Post Data Err"})
		return
	} else {
		excel.CreateExcelV2(&drawResult)

		context.Header("content-disposition", `attachment; filename=`+"drawResult.xlsx")
		context.Writer.Header().Add("Content-Type", "application/octet-stream")
		context.Header("Content-Transfer-Encoding", "binary")
		context.File(excel.DrawResultExcelV2)
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
