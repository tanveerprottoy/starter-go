package jwtpkg

import (
	"errors"

	"github.com/tanveerprottoy/starter-go/stdlib/pkg/config"

	"github.com/golang-jwt/jwt/v4"
)

// Create the JWT key used to create the signature
var JwtKey = []byte(config.GetEnvValue("secret"))

// jwt.RegisteredClaims is an embedded type
type Payload struct {
	Id string `json:"id"`
}

// jwt.RegisteredClaims is an embedded type
type Claims struct {
	Payload Payload `json:"payload"`
	jwt.RegisteredClaims
}

// GenerateToken generates a new token
func GenerateToken(payload Payload) string {
	/* RegisteredClaims: jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(expirationTime),
		IssuedAt:  jwt.NewNumericDate(datetime.Now()),
		NotBefore: jwt.NewNumericDate(datetime.Now()),
		Issuer:    "test",
		Subject:   "somebody",
		ID:        "1",
		Audience:  []string{"somebody_else"},
	}, */
	// Declare the expiration time of the token
	expirationTime := time.AddDate(0, 0, 3)
	claims := &Claims{
		Payload: payload,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Issuer:    "test",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(JwtKey)
	return tokenString
}

func VerifyToken(tokenBody string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(
		tokenBody,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			return JwtKey, nil
		},
	)
	if err != nil {
		return nil, errors.New("malformed token")
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}
