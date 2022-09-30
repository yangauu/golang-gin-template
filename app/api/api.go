package api

import (
	"github.com/gin-gonic/gin"
)

// 路由
func Default(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/query", ApiExamQuery)
		api.GET("/wechatpayresult", ApiWechatPayResult)

		api.POST("/delete", ApiExamDel)
		api.POST("/update", ApiExamEdit)
		api.POST("/insert", ApiExamInsert)
		api.POST("/uploadfile", ApiUploadFile)
		api.POST("/wechatpay", ApiWechatPay)
		api.POST("/alipaypay", ApiAliPay)
	}
}
