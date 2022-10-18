package api

import (
	_ "go-template/docs"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// 路由
func Default(r *gin.Engine) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, func(conf *ginSwagger.Config) {
		conf.Title = "模版文档"
	}))

	api := r.Group("/api")
	{
		api.GET("/query", ApiExamQuery)
		api.GET("/delqiniufile", ApiDelQiNiuFile)
		api.GET("/wechatpayresult", ApiWechatPayResult)

		api.POST("/delete", ApiExamDel)
		api.POST("/update", ApiExamEdit)
		api.POST("/alipaypay", ApiAliPay)
		api.POST("/insert", ApiExamInsert)
		api.POST("/wechatpay", ApiWechatPay)
		api.POST("/uploadfile", ApiUploadFile)
	}
}
