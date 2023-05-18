package helpers

type Service struct {
	method *Helper
}

type Helper interface {
	Encryption(plainText string)
}
