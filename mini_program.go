package epayments

import (
	"net/http"
)

// Refer to 2.12
// MiniProgram requests the payment by wechat mini program
type MiniProgram struct {
	MerchantID      string  `json:"merchant_id" validate:"required"`
	IncrementID     string  `json:"increment_id" validate:"required"`
	SubAppID        string  `json:"sub_appid" validate:"required"`
	SubOpenID       string  `json:"sub_openid" validate:"required"`
	GrandTotal      float64 `json:"grandtotal" validate:"required"`
	Currency        string  `json:"currency" validate:"required"`
	ValidMins       int64   `json:"valid_mins"`
	PaymentChannels string  `json:"payment_channels" validate:"required"`
	NotifyURL       string  `json:"notify_url" validate:"required"`
	Subject         string  `json:"subject"`
	Describe        string  `json:"describe" validate:"required"`
	OrderData       string  `json:"order_data"`
	SignType        string  `json:"sign_type" validate:"required"`
	NonceStr        string  `json:"nonce_str" validate:"required"`
	Signature       string  `json:"signature" validate:"required"`
	Service         string  `json:"service" validate:"required"`
}

func (e *MiniProgram) Do(cfg Config) (Response, int, error) {
	var response Response
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

func (e *MiniProgram) SetSignature(signature string) {
	(*e).Signature = signature
}

func (e *MiniProgram) SetSignType(signType string) {
	(*e).SignType = signType
}
