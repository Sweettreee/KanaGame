package main

import (
	mysql "KanaGame/mysqlclient"
	redis "KanaGame/redisclient"
	"KanaGame/router"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func Init() {
	// .env 파일 데이터 가져오기
	err := godotenv.Load()
	if err != nil {
		panic(fmt.Sprintf(".env파일 불러오기 실패: %v", err))
	}

	// Mysql 세팅하기
	mysql.InitMysql()
	// mysql.GetMysqlConnection()

	// Redis 세팅하기
	redis.InitRedis()
}

func main() {
	Init()

	port := os.Getenv("SERVER_PORT")

	r := router.SetupRouter()
	r.Run("0.0.0.0:" + port)
}
