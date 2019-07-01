package epayments

type RefundQueryResponse struct {
	ErrorCode
	MerchantID      string  `json:"merchant_id"`
	IncrementID     string  `json:"increment_id"`
	Currency        string  `json:"currency"`
	CreatedAt       string  `json:"created_at"`
	RefundFee       float64 `json:"refund_fee"`
	RefundTradeNo   string  `json:"refund_trade_no"`
	PaymentChannels string  `json:"payment_channels"`
	RefundReason    string  `json:"refund_reason"`
	RefundState     string  `json:"refund_state"`
	NonceStr        string  `json:"nonce_str"`
	Signature       string  `json:"signature"`
}
