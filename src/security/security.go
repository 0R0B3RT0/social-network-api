package security

import "golang.org/x/crypto/bcrypt"

// Hash encrypt the password
func Hash(pass string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
}

// VerifyPassword compare password and hash
func VerifyPassword(hash, pass string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
}
