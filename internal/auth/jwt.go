package auth

import (
	"github.com/golang-jwt/jwt"
)

/*
change-able claims:

	type StandardClaims struct {
	    Audience  string `json:"aud,omitempty"`
	    ExpiresAt int64  `json:"exp,omitempty"`
	    Id        string `json:"jti,omitempty"`
	    IssuedAt  int64  `json:"iat,omitempty"`
	    Issuer    string `json:"iss,omitempty"`
	    NotBefore int64  `json:"nbf,omitempty"`
	    Subject   string `json:"sub,omitempty"`
	}
*/
func GenerateJWT() (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{})

	token_str, err := claims.SignedString([]byte("dummy-secret-key"))
	if err != nil {
		return "", err
	}
	return token_str, nil
}
