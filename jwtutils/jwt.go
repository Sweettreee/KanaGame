package jwtutils

import (
	"KanaGame/redisclient"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretkey = os.Getenv("JWTSCERETKEY")

type Claims struct {
	UID int
	jwt.RegisteredClaims
}

// Access토큰을 생성하는 함수
func CreateAccessToken(uid int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		Claims{
			UID: uid,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(20 * time.Minute)), // 20분
				Issuer:    "KanaGame",
			},
		})

	tokenString, err := token.SignedString(secretkey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// Refresh토큰을 생성하는 함수
func CreateRefreshToken(uid int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		Claims{
			UID: uid,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Minute)), // 7일
				Issuer:    "KanaGame",
				Subject:   "Refresh",
			},
		})

	tokenString, err := token.SignedString(secretkey)
	if err != nil {
		return "", err
	}
	StoreRefreshToken(uid, tokenString, 7*24*time.Minute)

	return tokenString, nil
}

// Refresh토큰을 Redis에 캐싱하는 함수
func StoreRefreshToken(uid int, refreshToken string, ttl time.Duration) error {
	RDB := redisclient.InitRedis()
	return RDB.Set(redisclient.Ctx, fmt.Sprintf("refresh:%v", uid), refreshToken, ttl).Err()
}

// Redis에 저장된 Refresh토큰을 지우는 함수
func DeleteRefreshToken(uid int) error {
	RDB := redisclient.InitRedis()
	return RDB.Del(redisclient.Ctx, fmt.Sprintf("refresh:%v", uid)).Err()
}

// 토큰을 인증하는 함수
func VerifyToken(tokenString string) (int, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return secretkey, nil
	})

	if err != nil {
		return 0, err
	}

	if !token.Valid {
		return 0, fmt.Errorf("invalid token")
	}
	return claims.UID, nil
}

// Refresh토큰을 받아서 Access토큰을 재생성하는 함수
func RefreshAccessToken(refreshToken string) (string, error) {

	if _, err := VerifyToken(refreshToken); err != nil {
		return "", err
	}

	claims := &Claims{}
	token, _ := jwt.ParseWithClaims(refreshToken, claims, func(t *jwt.Token) (interface{}, error) {
		return secretkey, nil
	})

	if token == nil || !token.Valid {
		return "", fmt.Errorf("invalid refresh token")
	}

	RDB := redisclient.InitRedis()
	storedToken, err := RDB.Get(redisclient.Ctx, fmt.Sprintf("refresh:%v", claims.UID)).Result()
	if err != nil || storedToken != refreshToken {
		return "", fmt.Errorf("refresh token mismatch")
	}

	return CreateAccessToken(claims.UID)
}
