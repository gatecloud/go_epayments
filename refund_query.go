package epayments

import (
	"net/http"
)

// Refer to 2.7
// RefundQuery is used to execute refunding query
type RefundQuery struct {
	MerchantID    string `json:"merchant_id" validate:"required"`
	IncrementID   string `json:"increment_id" validate:"required"`
	RefundTradeNo string `json:"refund_trade_no" validate:"required"`
	SignType      string `json:"sign_type" validate:"required"`
	NonceStr      string `json:"nonce_str" validate:"required"`
	Signature     string `json:"signature" validate:"required"`
	Service       string `json:"service" validate:"required"`
}

func (e *RefundQuery) Do(cfg Config) (RefundQueryResponse, int, error) {
	var response RefundQueryResponse
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

func (e *RefundQuery) SetSignature(signature string) {
	(*e).Signature = signature
}

func (e *RefundQuery) SetSignType(signType string) {
	(*e).SignType = signType
}
