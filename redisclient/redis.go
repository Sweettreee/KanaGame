package redisclient

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

var RDB *redis.Client
var Ctx = context.Background()

func InitRedis(addr, password string, db int) {
	log.Println("before Client")
	RDB = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	log.Println("after Client/before Ping")
	// 연결 테스트
	pong, err := RDB.Ping(Ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Redis 연결 실패: %v", err))
	}
	log.Println("Redis 연결 성공:", pong)
}
