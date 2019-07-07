package epayments

type RateQueryResponse struct {
	ErrorCode
	Currency  string  `json:"currency"`
	RateTime  string  `json:"rate_time"`
	Rate      float64 `json:"rate"`
	NonceStr  string  `json:"nonce_str"`
	Signature string  `json:"signature"`
	SignType  string  `json:"sign_type"`
}

func (e *RateQueryResponse) GetSignature() string {
	return e.Signature
}
