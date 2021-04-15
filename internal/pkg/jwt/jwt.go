package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret = []byte("123@456")

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

// 生成jwt签名字符串
func GenerateToken(username, password string) (string, error) {
	// 组装payload信息
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)
	claims := Claims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "blog",
		},
	}

	// 组装header，设置签名算法
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 生成签名（base64url -> hmacsha256 = token）
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

// 验证jwt字符串有效性
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
