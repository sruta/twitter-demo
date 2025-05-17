package pkg

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	ID int64 `json:"id"`
	jwt.StandardClaims
}

var JwtSecret string
var ExpirationTime time.Duration

func InitializeJWT(secret string, expiration time.Duration) {
	JwtSecret = secret
	ExpirationTime = expiration
}

func GenerateJwtToken(id int64) (string, error) {
	claims := Claims{
		ID: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(ExpirationTime).Unix(),
		},
	}

	unsignedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	signedToken, err := unsignedToken.SignedString([]byte(JwtSecret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func VerifyJwtToken(str string) (Claims, error) {
	var claims Claims
	token, err := jwt.ParseWithClaims(str, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(JwtSecret), nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrSignatureInvalid) {
			return claims, errors.New("invalid token signature")
		}
	}

	if !token.Valid {
		return claims, errors.New("invalid token")
	}

	return claims, nil
}
