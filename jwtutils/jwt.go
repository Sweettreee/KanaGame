package jwtutils

import (
	"fmt"
	"time"
	"os"
	"github.com/golang-jwt/jwt/v5"
) // v5 문서의 API와 예제 기준

func SecretFromEnv() ([]byte, error) {
	// 1) 평문 시크릿 사용
	if v := os.Getenv("JWTSCERETKEY"); v != "" {
		return []byte(v), nil // HMAC은 []byte가 필요 [4]
	}

	return nil, fmt.Errorf("missing JWT_SECRET")
}

func CreateToken(uid string) (string, error) { // 패키지 외부에서 사용하려면 식별자를 대문자로 시작해 export해야 함
	secretKey, errenv := SecretFromEnv()
	if(errenv!=nil){
		fmt.Println("jwt Error : ", errenv)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		// create JWT token by jew.NewWithClaims()
		// method : HS256
		jwt.MapClaims{ // relevant information
			"UID": uid,
			"ISS": "kanagame",
			"exp": time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secretKey) // secretekey
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyToken(tokenString string) error {
	secretKey, errenv := SecretFromEnv()
	if(errenv!=nil){
		fmt.Println("jwt Error : ", errenv)
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// For parsing and verifying
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}
