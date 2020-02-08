package crypto

import (
	"golang.org/x/crypto/bcrypt"
)

// PasswordEncrypt パスワードをhash化
func PasswordEncrypt(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

// CompareHashAndPassword hashと非hashパスワード比較
func CompareHashAndPassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
