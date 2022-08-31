package authutil

import "golang.org/x/crypto/bcrypt"

func Hash(hash string, len int) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(hash), len)

	result := string(bytes)

	return result, err
}

func Compare(hashStr string, origin string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(origin), []byte(hashStr))
	return err == nil
}
