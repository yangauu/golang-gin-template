package wechat

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/wechat/v3"
)

// JSP支付订单创建
func JspPayCreate() (string, error) {
	bm := make(gopay.BodyMap)
	bm.Set("appid", "wx418bda3b36df2747")
	bm.Set("description", "测试Jsapi支付商品")
	bm.Set("out_trade_no", strconv.FormatInt(time.Now().Unix(), 10))
	bm.Set("time_expire", time.Now().Add(10*time.Minute).Format(time.RFC3339))
	bm.Set("notify_url", "https://www.fmm.ink")
	bm.SetBodyMap("amount", func(bm gopay.BodyMap) {
		bm.Set("total", 1).Set("currency", "CNY")
	})
	bm.SetBodyMap("payer", func(bm gopay.BodyMap) {
		bm.Set("openid", "o9LCN5NYKS_VM8Y90_e_Wqs7OTl8")
	})

	wxRsp, err := client.V3TransactionJsapi(context.TODO(), bm)
	if err != nil {
		log.Println("JsapiPayCreate 下单出错", err)
		return "", err
	}
	return wxRsp.Response.PrepayId, nil
}

// JSP支付
func JspPay() (*wechat.JSAPIPayParams, error) {
	e := &wechat.JSAPIPayParams{}

	pid, err := JspPayCreate()
	if err != nil {
		log.Println("JspPay 下单出错", err)
		return e, err
	}

	jsapi, err := client.PaySignOfJSAPI("wx418bda3b36df2747", pid)
	if err != nil {
		log.Println(err)
		return e, err
	}

	return jsapi, nil
}

// JSP查询订单  OrderNoType 1-微信订单号，2-商户订单号，3-微信侧回跳到商户前端时用于查单的单据查询id（查询支付分订单中会使用）
func JspPayResult() {
	wxRsp, err := client.V3TransactionQueryOrder(context.TODO(), 2, strconv.FormatInt(time.Now().Unix(), 10))
	if err != nil {
		log.Println("JspPayResult 查询结果出错", err)
		return
	}
	log.Println("交易结果", wxRsp.Response.TradeState)
}
