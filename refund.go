package epayments

import (
	"net/http"
)

type Refund struct {
	MerchantID    string  `json:"merchant_id" validate:"required"`
	IncrementID   string  `json:"increment_id" validate:"required"`
	Currency      string  `json:"currency" validate:"required"`
	RefundFee     float64 `json:"refund_fee" validate:"required"`
	RefundReason  string  `json:"refund_reason"`
	RefundTradeNo string  `json:"refund_trade_no"`
	RefundOrder   string  `json:"refund_order"`
	SignType      string  `json:"sign_type" validate:"required"`
	NonceStr      string  `json:"nonce_str" validate:"required"`
	Signature     string  `json:"signature" validate:"required"`
	Service       string  `json:"service" validate:"required"`
}

func (e *Refund) Do(cfg Config) (RefundResponse, int, error) {
	var response RefundResponse
	if statusCode, err := cfg.Sign(e); err != nil {
		return response, statusCode, err
	}

	parameters, err := toURLParams(e)
	if err != nil {
		return response, http.StatusInternalServerError, err
	}

	url := cfg.Endpoint + "?" + parameters
	statusCode, err := DoRequest("GET", url, nil, &response)
	if err != nil {
		return response, statusCode, err
	}

	if statusCode, err := cfg.Verify(&response); err != nil {
		return response, statusCode, err
	}

	statusCode, err = response.Validate()
	if err != nil {
		return response, statusCode, err
	}

	return response, statusCode, nil
}

func (e *Refund) SetSignature(signature string) {
	(*e).Signature = signature
}

func (e *Refund) SetSignType(signType string) {
	(*e).SignType = signType
}
