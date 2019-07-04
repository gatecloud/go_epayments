package epayments

import "net/http"

type AsyncResultVerification struct {
	MerchantID  string `json:"merchant_id" validate:"required"`
	IncrementID string `json:"increment_id" validate:"required"`
	NotifyID    string `json:"notify_id" validate:"required"`
	SignType    string `json:"sign_type" validate:"required"`
	NonceStr    string `json:"nonce_str" validate:"required"`
	Signature   string `json:"signature" validate:"required"`
	Service     string `json:"service" validate:"required"`
}

func (e *AsyncResultVerification) Do(cfg Config) (int, error) {
	var errCode ErrorCode

	if statusCode, err := cfg.Sign(e); err != nil {
		return statusCode, err
	}

	parameters, err := toURLParams(e)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	url := cfg.Endpoint + "?" + parameters
	statusCode, err := DoRequest("GET", url, nil, &errCode)
	if err != nil {
		return statusCode, err
	}

	return errCode.Validate()
}

func (e *AsyncResultVerification) SetSignature(signature string) {
	(*e).Signature = signature
}

func (e *AsyncResultVerification) SetSignType(signType string) {
	(*e).SignType = signType
}
