package epayments

type Signaturer interface {
	SetSignature(signature string)
	SetSignType(signType string)
}
