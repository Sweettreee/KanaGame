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
	jwt.RegisterdClaims
}

func CreateAccessToken(uid int) (string, error) { // 패키지 외부에서 사용하려면 식별자를 대문자로 시작해 export해야 함
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		Claims{
			UID: uid,
			RegisterdClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(20 * time.Minute)), // 20분
				Issuer:    "KanaGame",
			},
		})

	tokenString, err := token.SignedString(secretkey) // secretekey
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func CreateRefreshToken(uid int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		Claims{
			UID: uid,
			RegisterdClaims: jwt.RegisteredClaims{
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

func StoreRefreshToken(uid int, refreshToken string, ttl time.Duration) error {
	RDB := redisclient.InitRedis()
	return RDB.Set(redisclient.Ctx, fmt.Sprintf("refresh:%v", uid), refreshToken, ttl).Err()
}

func DeleteRefreshToken(uid int) error {
	RDB := redisclient.InitRedis()
	return RDB.Del(redisclient.Ctx, fmt.Sprintf("refresh:%v", uid)).Err()
}

func VerifyToken(tokenString string) error {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return secretkey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}
	return nil
}

func RefreshAccessToken(refreshToken string) (string, error) {

	if err := VerifyToken(refreshToken); err != nil {
		return "", err
	}

	// 2️⃣ Claims 파싱
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
