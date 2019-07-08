package epayments

import "net/http"

// Refer to 2.8
// TradeClosed is used to close the trade
type TradeClosed struct {
	MerchantID  string `json:"merchant_id" validate:"required"`
	IncrementID string `json:"increment_id" validate:"required"`
	SignType    string `json:"sign_type" validate:"required"`
	NonceStr    string `json:"nonce_str" validate:"required"`
	Service     string `json:"service" validate:"required"`
	Signature   string `json:"signature" validate:"required"`
}

func (e *TradeClosed) Do(cfg Config) (RefundQueryResponse, int, error) {
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

func (e *TradeClosed) SetSignature(signature string) {
	(*e).Signature = signature
}

func (e *TradeClosed) SetSignType(signType string) {
	(*e).SignType = signType
}
