package jwtutils

import (
	"fmt"
	"time"
	"github.com/golang-jwt/jwt/v5"
	"os"
) // v5 문서의 API와 예제 기준

// var secretKey = []byte("secret-key") 로컬에서 돌릴 때 사용
var secretkey = os.Getenv("JWTSCERETKEY");

func CreateToken(uid string) (string, error) { // 패키지 외부에서 사용하려면 식별자를 대문자로 시작해 export해야 함
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, 
		// create JWT token by jew.NewWithClaims()
		// method : HS256
		jwt.MapClaims{ // relevant information
		"UID" : uid,
		"ISS" : "kanagame",
		"exp" : time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(secretkey) // secretekey
	if err != nil {
		return "",err
	}
	return tokenString, nil
}

func VerifyToken(tokenString string) error {
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
