package epayments

import "net/http"

type AsyncResultVerification struct {
	MerchantID  string `json:"merchant_id"`
	IncrementID string `json:"increment_id"`
	NotifyID    string `json:"notify_id"`
	SignType    string `json:"sign_type"`
	NonceStr    string `json:"nonce_str"`
	Signature   string `json:"signature"`
	Service     string `json:"service"`
}

func (e *AsyncResultVerification) Do(cfg Config) (int, error) {
	var errCode ErrorCode

	if err := cfg.Sign(e); err != nil {
		return http.StatusInternalServerError, err
	}

	parameters, err := ToURLParams(e)
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
