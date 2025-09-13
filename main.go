package main

import (
	"KanaGame/mysqlclient"
	"KanaGame/redisclient"
	"fmt"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(fmt.Sprintf(".env파일 불러오기 실패: %v", err))
	}

	RDB := redisclient.InitRedis()
	//mysqlclient.InitMysql()
	mysqlclient.GetMysqlConnection()

	err = RDB.Set(redisclient.Ctx, "greeting", "Hello Go Redis", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := RDB.Get(redisclient.Ctx, "greeting").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Redis에서 가져온 값:", val)
}
