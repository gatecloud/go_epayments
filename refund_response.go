package epayments

type RefundResponse struct {
	ErrorCode
	RefundFee     string `json:"refund_fee"`
	RefundTradeNo string `json:"refund_trade_no"`
	RefundPayTime string `json:"refund_pay_time"`
	RefundState   string `json:"refund_state"`
	NonceStr      string `json:"nonce_str"`
	Signature     string `json:"signature"`
	SignType      string `json:"sign_type"`
}
