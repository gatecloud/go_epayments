package epayments

type TradeQueryResponse struct {
	ErrorCode
	GrandTotal      float64 `json:"grandtotal"`
	ReceiptAmount   float64 `json:"receipt_amount"`
	Currency        string  `json:"currency"`
	Subject         string  `json:"subject"`
	Describe        string  `json:"describe"`
	TradeNo         string  `json:"trade_no"`
	CreatedAt       string  `json:"created_at"`
	GmtPayment      string  `json:"gmt_payment"`
	TradeStatus     string  `json:"trade_status"`
	PaymentChannels string  `json:"payment_channels"`
	NonceStr        string  `json:"nonce_str"`
	Signature       string  `json:"signature"`
	SignType        string  `json:"sign_type"`
}

func (e *TradeQueryResponse) GetSignature() string {
	return e.Signature
}
