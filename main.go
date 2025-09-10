package main

import (
	"KanaGame/mysqlclient"
	"KanaGame/redisclient"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(fmt.Sprintf(".env파일 불러오기 실패: %v", err))
	}

	redisclient.InitRedis(fmt.Sprintf("%v:%v", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")), os.Getenv("REDIS_PASSWORD"), 0)
	mysqlclient.InitMysql(os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_DB"), os.Getenv("MYSQL_PORT"))

	err = redisclient.RDB.Set(redisclient.Ctx, "greeting", "Hello Go Redis", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := redisclient.RDB.Get(redisclient.Ctx, "greeting").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Redis에서 가져온 값:", val)
}
