package utils

import (
	"Go-blog-server/pkg/setting"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret = []byte(setting.JwtSecret)

type Claims struct {
	ID int          `json:"id"`
	Username string `json:"username"`
	Role     int    `json:"role"`
	jwt.StandardClaims
}

func GenerateToken(id int, username string, role int) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(72 * time.Hour)

	claims := Claims{
		id,
		username,
		role,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "chakos",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
