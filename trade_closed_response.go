package epayments

type TradeClosedResponse struct {
	ErrorCode
	MerchantID   string `json:"merchant_id"`
	IncrementID  string `json:"increment_id"`
	RefundStatus string `json:"refund_status"`
	NonceStr     string `json:"nonce_str"`
	Signature    string `json:"signature"`
	SignType     string `json:"sign_type"`
}

func (e *TradeClosedResponse) GetSignature() string {
	return e.Signature
}
