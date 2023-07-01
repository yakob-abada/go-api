package service

import "golang.org/x/crypto/bcrypt"

func NewAppCrypto() *AppCrypto {
	return &AppCrypto{}
}

type AppCrypto struct {
}

func (ap *AppCrypto) GenerateFromPassword(password []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
}

func (ap *AppCrypto) CompareHashAndPassword(hashedPassword, password []byte) error {
	return bcrypt.CompareHashAndPassword(hashedPassword, password)
}
