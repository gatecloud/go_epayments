package epayments

type Verifier interface {
	GetSignature() string
}
