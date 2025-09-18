package redisclient

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/redis/go-redis/v9"
)

var RDB *redis.Client
var Ctx = context.Background()
var once sync.Once

func InitRedis() *redis.Client {
	once.Do(func() {
		RDB = redis.NewClient(&redis.Options{
			Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
			Password: os.Getenv("REDIS_PASSWORD"),
			DB:       0,
		})

		pong, err := RDB.Ping(Ctx).Result()
		if err != nil {
			panic(fmt.Sprintf("Redis 연결 실패: %v", err))
		}
		log.Println("Redis 연결 성공:", pong)
	})

	return RDB
}

func CloseRedis() {
	if RDB != nil {
		if err := RDB.Close(); err == nil {
			panic(fmt.Sprintf("Redis 연결 닫기 실패: %v", err))
		} else {
			log.Printf("Redis 연결 닫기 성공")
		}
		RDB = nil
	}
}
