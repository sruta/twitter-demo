package helpers

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
	"twitter-uala/internal/domain"
)

type Claims struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}

var JwtSecret string
var ExpirationTime time.Duration

func InitializeJWT(secret string, expiration time.Duration) {
	JwtSecret = secret
	ExpirationTime = expiration
}

func GenerateJwtToken(user domain.User) (string, error) {
	claims := Claims{
		ID:    user.ID,
		Email: user.Email,
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
