package epayments

type AsyncResult struct {
	NotifyID        string  `form:"notify_id" json:"notify_id"`
	MerchantID      string  `form:"merchant_id" json:"merchant_id"`
	IncrementID     string  `form:"increment_id" json:"increment_id"`
	GrandTotal      float64 `form:"grandtotal" json:"grandtotal"`
	ReceiptAmount   float64 `form:"receipt_amount" json:"receipt_amount"`
	Currency        string  `form:"currency" json:"currency"`
	Subject         string  `form:"subject" json:"subject"`
	Describe        string  `form:"describe" json:"describe"`
	SignType        string  `form:"sign_type" json:"sign_type"`
	Signature       string  `form:"signature" json:"signature"`
	Service         string  `form:"service" json:"service"`
	TradeNo         string  `form:"trade_no" json:"trade_no"`
	NotifyTime      string  `form:"notify_time" json:"notify_time"`
	CreatedAt       string  `form:"created_at" json:"created_at"`
	GmtPayment      string  `form:"gmt_payment" json:"gmt_payment"`
	TradeStatus     string  `form:"trade_status" json:"trade_status"`
	PaymentChannels string  `form:"payment_channels" json:"payment_channels"`
}
