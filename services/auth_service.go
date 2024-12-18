package services

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword generate a hash from a password
func HashPassword(password string) (string, error) {
	// Genera el hash con un costo (cost) de trabajo. Cuanto mayor sea el costo, más seguro, pero más lento.
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// CheckPasswordHash verify if the password is correct
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
