package jwt

import (
	"math/rand"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

var jwtTimeout int
var jwtSecret []byte

type Claims struct {
	Name string `json:"name"`
	jwt.RegisteredClaims
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, nil
}

func GenerateToken(name string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Duration(jwtTimeout) * time.Minute)

	claims := Claims{
		name,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			Issuer:    "iam",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func InitJwtWrapper(timeout int) error {
	jwtTimeout = timeout

	jwtSecret = make([]byte, 32)
	_, err := rand.Read(jwtSecret)
	return err
}
