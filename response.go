package epayments

type Response struct {
	ErrorCode
	MerchantID  string `json:"merchant_id"`
	IncrementID string `json:"increment_id"`
	PayJson     string `json:"payJson"`
	NonceStr    string `json:"nonce_str"`
	Signature   string `json:"signature"`
	SignType    string `json:"sign_type"`
}

func (e *Response) GetSignature() string {
	return e.Signature
}
