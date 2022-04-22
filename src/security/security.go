package security

import "golang.org/x/crypto/bcrypt"

// Hash string
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// Compare a password and a passwored hashed and return if equals
func VerifyPassword(passwordHashed string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordHashed), []byte(password))
}
