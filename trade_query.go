package epayments

import (
	"net/http"
)

type TradeQuery struct {
	MerchantID  string `json:"merchant_id" validate:"required"`
	IncrementID string `json:"increment_id" validate:"required"`
	SignType    string `json:"sign_type" validate:"required"`
	NonceStr    string `json:"nonce_str" validate:"required"`
	Signature   string `json:"signature" validate:"required"`
	Service     string `json:"service" validate:"required"`
}

func (e *TradeQuery) Do(cfg Config) (TradeQueryResponse, int, error) {
	var response TradeQueryResponse
	if statusCode, err := cfg.Sign(e); err != nil {
		return response, statusCode, err
	}

	parameters, err := ToURLParams(e)
	if err != nil {
		return response, http.StatusInternalServerError, err
	}

	url := cfg.Endpoint + "?" + parameters
	statusCode, err := DoRequest("GET", url, nil, &response)
	if err != nil {
		return response, statusCode, err
	}

	// if statusCode, err := cfg.Verify(&response); err != nil {
	// 	return response, statusCode, err
	// }

	statusCode, err = response.Validate()
	if err != nil {
		return response, statusCode, err
	}

	return response, statusCode, nil
}

func (e *TradeQuery) SetSignature(signature string) {
	(*e).Signature = signature
}

func (e *TradeQuery) SetSignType(signType string) {
	(*e).SignType = signType
}
