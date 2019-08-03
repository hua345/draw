package v1

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
		excel.CreateExcel(&drawResult)

		context.Header("content-disposition", `attachment; filename=`+"drawResult.xlsx")
		context.Writer.Header().Add("Content-Type", "application/octet-stream")
		context.File(excel.DrawResultExcel)
	}
}
