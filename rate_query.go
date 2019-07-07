package epayments

import (
	"net/http"
)

// Refer to 2.10
// RateQuery requests the exchange rate
type RateQuery struct {
	MerchantID      string `json:"merchant_id" validate:"required"`
	Currency        string `json:"currency" validate:"required"`
	PaymentChannels string `json:"payment_channels" validate:"required"`
	BusinessType    string `json:"business_type"`
	SignType        string `json:"sign_type" validate:"required"`
	NonceStr        string `json:"nonce_str" validate:"required"`
	Signature       string `json:"signature" validate:"required"`
	Service         string `json:"service" validate:"required"`
}

func (e *RateQuery) Do(cfg Config) (RateQueryResponse, int, error) {
	var response RateQueryResponse
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

func (e *RateQuery) SetSignature(signature string) {
	(*e).Signature = signature
}

func (e *RateQuery) SetSignType(signType string) {
	(*e).SignType = signType
}
