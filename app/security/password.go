package security

import (
	"golang.org/x/crypto/bcrypt"
)

// Password to hash
func PasswordToHash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// Compare hash with password
func CompareHashWithPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
