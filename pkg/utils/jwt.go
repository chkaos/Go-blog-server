package utils

import (
	"Go-blog-server/pkg/setting"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret = []byte(setting.JwtSecret)

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     int    `json:"role"`
	jwt.StandardClaims
}

func GenerateToken(username, password string, role int) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(48 * time.Hour)

	claims := Claims{
		username,
		password,
		role,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "astella",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	fmt.Println(err)
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
