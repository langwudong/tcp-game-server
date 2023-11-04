package mysql

import (
	"golang.org/x/crypto/bcrypt"
)

func EncryptPwd(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

func ComparePwd(password string, encryptionPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encryptionPwd), []byte(password))
	if err != nil {
		return false
	}
	return true
}
