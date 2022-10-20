package api

import (
	"go-template/app/conf"
	"go-template/app/db"
	"go-template/app/pay/alipay"
	"go-template/app/pay/wechat"
	"go-template/app/qiniu"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary 删除七牛文件
// @Tags 文件
// @Accept mpfd
// @Param filename query string true "文件名称"
// @Produce json
// @Success 200 {object} conf.SuccApiDelQiNiuFile
// @Failure 400 {object} conf.ErrApiDelQiNiuFile
// @Router /api/delqiniufile [get]
func ApiDelQiNiuFile(c *gin.Context) {
	f := c.Query("filename")
	err := qiniu.DeleteToQiNiu("image/" + f)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "删除失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "删除成功",
		"data": f,
	})
}

// 支付宝支付
func ApiAliPay(c *gin.Context) {
	url, err := alipay.ErweimeAlipay()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "ApiAliPay 支付下单出错",
			"data": "",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "支付下单成功",
		"data": url,
	})
}

// 微信支付
func ApiWechatPay(c *gin.Context) {
	jsparam, err := wechat.JspPay()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "获取支付参数出错",
			"data": "",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "获取支付参数成功",
		"data": jsparam,
	})
}

// 上传
func ApiUploadFile(c *gin.Context) {
	f, err := c.FormFile("file")

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  err.Error(),
		})
		return
	}

	url, err := qiniu.UploadToQiNiu("image/", f)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "上传失败",
			"url":  "",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "上传成功",
		"url":  url,
	})
}

// 查询支付结果
func ApiWechatPayResult(c *gin.Context) {
	wechat.JspPayResult()

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "请求成功",
		"data": "result",
	})
}

// @Summary 删除用户
// @Tags 用户
// @Accept mpfd
// @Produce json
// @Param id body string true "用户id"
// @Success 200 {object} conf.SuccApiExamDel
// @Failure 400 {object} conf.ErrApiExamDel
// @Router /api/delete [post]
func ApiExamDel(c *gin.Context) {
	json := conf.User{}
	c.Bind(&json)

	err := db.ExamDel(json.Id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "删除成功",
	})
}

// 添加
func ApiExamInsert(c *gin.Context) {
	json := conf.User{}
	c.Bind(&json)

	err := db.ExamInsert(json)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "添加成功",
	})
}

// @Summary 获取用户列表
// @Tags 用户
// @Accept mpfd
// @Produce json
// @Param page query int true "页码"
// @Success 200 {object} conf.SuccApiExamQuery
// @Failure 400 {object} conf.ErrApiExamQuery
// @Router /api/query [get]
func ApiExamQuery(c *gin.Context) {
	p := c.Query("page")
	page, _ := strconv.Atoi(p)

	result, err := db.ExamQuery(page)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "获取列表出错",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "获取成功",
		"data": result,
	})
}

// 更新
func ApiExamEdit(c *gin.Context) {
	json := conf.User{}
	c.Bind(&json)

	err := db.ExamEdit(json)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "更新成功",
	})
}
