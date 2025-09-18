package main

import (
	mysql "KanaGame/mysqlclient"
	redis "KanaGame/redisclient"
	"KanaGame/router"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
)

func Init() {
	// .env 파일 데이터 가져오기
	err := godotenv.Load()
	if err != nil {
		panic(fmt.Sprintf(".env파일 불러오기 실패: %v", err))
	}
	// Redis 세팅하기
	redis.InitRedis()

	// Mysql 세팅하기
	mysql.InitMysql()

}

func CloseResources() {
	redis.CloseRedis()
	mysql.CloseMysql()
}

func main() {
	Init()

	port := os.Getenv("SERVER_PORT")

	r := router.SetupRouter()

	server := &http.Server{
		Addr:    "0.0.0.0:" + port,
		Handler: r,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Printf("Shutdown Server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: ", err)
	}
	CloseResources()
	log.Printf("Server exiting")
}
