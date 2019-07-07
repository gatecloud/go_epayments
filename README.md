# epayments
Open source Golang SDK for ePayments payment interface (ePayments聚合支付接口)

This SDK is developed based on [ePayment API document](https://www.kiwifast.com/doc/index.html)  



### Checklist

:black_square_button: 2.1 聚合支付接口  
:black_square_button: 2.2 页面跳转同步通知接口  
:white_check_mark: 2.3 支付结果异步通知  
:white_check_mark: 2.3+ 支付结果异步通知验证接口  
:white_check_mark: 2.4 支付订单交易查询接口  
:white_check_mark: 2.5 申请退款接口  
:white_check_mark: 2.6 错误异常响应返回参数  
:white_check_mark: 2.7 退款查询接口  
:black_square_button: 2.8 交易关闭接口  
:black_square_button: 2.9 刷卡支付接口  
:white_check_mark: 2.10 汇率查询接口  
:black_square_button: 2.11 聚合二维码支付接口  
:white_check_mark: 2.12 小程序支付接口  
:black_square_button: 2.13 聚合APP支付接口  
:black_square_button: 2.14 自定义二维码支付接口  


### Install

```
$ go get github.com/OscarZhou/epayments
```


### Notice

1. Your merchant ID binds with a specific currency  
2. Some fields are not displayed in the API document. For example, `rate` in `AsyncResult`, `TradeQueryResponse` and `RefundQueryResponse`  



### Sample

#### 2.4  

```
config := epayments.Config{
    SignKey:  "YOUR_SIGN_KEY",
    Endpoint: "https://www.kiwifast.com/api/v1/info/smartpay",
}

tradeQuery := &epayments.TradeQuery{
    MerchantID:  "YOUR_MERCHANT_ID",
    IncrementID: "1101",
    NonceStr:    "YptpkflFlO",
    Service:     "create_trade_query",
}

response, statusCode, err := tradeQuery.Do(config)
if err != nil {
    fmt.Printf("result:%v\nhttp code:%d\nmessage:%s\n",response, statusCode, err)
}


``` 


#### 2.5 


```
config := epayments.Config{
    SignKey:  "YOUR_SIGN_KEY",
    Endpoint: "https://www.kiwifast.com/api/v1/info/smartpay",
}

refund := &epayments.Refund{
    MerchantID:   "YOUR_MERCHANT_ID",
    IncrementID:  "1101",
    RefundFee:    100.00,
    RefundReason: "",
    Currency:     "CNY",
    NonceStr:     ""YptpkflFlO",",
    Service:      "create_trade_refund",
}

response, statusCode, err := refund.Do(config)
if err != nil {
    fmt.Printf("result:%v\nhttp code:%d\nmessage:%s\n",response, statusCode, err)
}

```


#### 2.7

```
config := epayments.Config{
    SignKey:  "YOUR_SIGN_KEY",
    Endpoint: "https://www.kiwifast.com/api/v1/info/smartpay",
}

refundQuery := &epayments.RefundQuery{
    MerchantID:    "YOUR_MERCHANT_ID",
    IncrementID:   "1101",
    RefundTradeNo: "R201655846",
    NonceStr:      "YptpkflFlO",
    Service:       "create_trade_refund_query",
}

response, statusCode, err := tradeQuery.Do(config)
if err != nil {
    fmt.Printf("result:%v\nhttp code:%d\nmessage:%s\n",response, statusCode, err)
}

```
