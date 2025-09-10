package main

import (
	"KanaGame/redisclient"
	"fmt"
)

func main() {
	redisclient.InitRedis("localhost:6379", "", 0)

	fmt.Println("starting...")

	err := redisclient.RDB.Set(redisclient.Ctx, "greeting", "Hello Go Redis", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := redisclient.RDB.Get(redisclient.Ctx, "greeting").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Redis에서 가져온 값:", val)
}
