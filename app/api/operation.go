package api

import (
	"go-template/app/conf"
	"go-template/app/db"
	"go-template/app/pay/alipay"
	"go-template/app/pay/wechat"
	"go-template/app/qiniu"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 支付宝支付
func ApiAliPay(c *gin.Context) {
	url, err := alipay.ErweimeAlipay()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusBadRequest,
			"msg":    "ApiAliPay 支付下单出错",
			"data":   "",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"msg":    "支付下单成功",
		"data":   url,
	})
}

// 微信支付
func ApiWechatPay(c *gin.Context) {
	jsparam, err := wechat.JspPay()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusBadRequest,
			"msg":    "获取支付参数出错",
			"data":   "",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"msg":    "获取支付参数成功",
		"data":   jsparam,
	})
}

// 上传
func ApiUploadFile(c *gin.Context) {
	f, err := c.FormFile("file")

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusBadRequest,
			"msg":    err.Error(),
		})
		return
	}

	url, err := qiniu.UploadToQiNiu("image/", f)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusBadRequest,
			"msg":    "上传失败",
			"url":    "",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"msg":    "上传成功",
		"url":    url,
	})
}

// 查询
func ApiExamQuery(c *gin.Context) {
	result, _ := db.ExamQuery()

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"msg":    "请求成功",
		"data":   result,
	})
}

// 查询支付结果
func ApiWechatPayResult(c *gin.Context) {
	wechat.JspPayResult()

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"msg":    "请求成功",
		"data":   "result",
	})
}

// 删除
func ApiExamDel(c *gin.Context) {
	json := conf.UserCollection{}
	c.Bind(&json)

	err := db.ExamDel(json.Id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"msg":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"msg":    "删除成功",
	})
}

// 添加
func ApiExamInsert(c *gin.Context) {
	json := conf.UserCollection{}
	c.Bind(&json)

	err := db.ExamInsert(json)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"msg":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"msg":    "添加成功",
	})
}

// 更新
func ApiExamEdit(c *gin.Context) {
	json := conf.UserCollection{}
	c.Bind(&json)

	err := db.ExamEdit(json)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"msg":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"msg":    "更新成功",
	})
}
