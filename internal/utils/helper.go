package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
    // Generate a salted hash for the password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return "", err
    }
    return string(hashedPassword), nil
}

// Compares a hashed password with a plain text password
func ComparePasswords(hashedPassword, password string) error {
    // Compare the hashed password with the plain text password
    return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
