package auth

import "golang.org/x/crypto/bcrypt"

// takes in a plaintext string, returns a hashed and salted password from bcrypt, plus error code.
func ProcessPassword(password string) (string, error) {
	//Use standard cost here. use bcrypt.MaxCost if needed for extra security.
	hashedPW, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPW), nil
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
