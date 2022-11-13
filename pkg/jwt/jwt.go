package jwt

import (
	"errors"
	"txp/restapistarter/pkg/config"
	"txp/restapistarter/pkg/datetime"

	"github.com/golang-jwt/jwt/v4"
)

// Create the JWT key used to create the signature
var JwtKey = []byte(config.GetEnvValue("secret"))

// jwt.StandardClaims is an embedded type
type Claims struct {
	Payload any `json:"payload"`
	jwt.RegisteredClaims
}

// GenerateToken generates a new token
func GenerateToken(payload any) string {
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
	expirationTime := datetime.AddDate(0, 0, 3)
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
