package mysqlclient

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

type _StmtList struct {
	InsertUser *sql.Stmt
	LoginUser  *sql.Stmt
}

var StmtList _StmtList
var MDB *sql.DB
var once sync.Once

func InitMysql() *sql.DB {
	GetMysqlConnection()

	var err error

	StmtList.InsertUser, err = MDB.Prepare("INSERT INTO users (login_id, login_pw, username) VALUES (?, sha2(?, 256), ?)")
	if err != nil {
		log.Fatalf("InsertUser 문 준비 실패: %v", err)
	}

	StmtList.LoginUser, err = MDB.Prepare("SELECT uid, login_id, username, password FROM users WHERE username = ?")
	if err != nil {
		log.Fatalf("LoginUser 문 준비 실패: %v", err)
	}

	return MDB
}

func GetMysqlConnection() *sql.DB {
	once.Do(func() {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
			os.Getenv("MYSQL_USER"),
			os.Getenv("MYSQL_PASSWORD"),
			os.Getenv("MYSQL_HOST"),
			os.Getenv("MYSQL_PORT"),
			os.Getenv("MYSQL_DB"),
		)

		mdb, err := sql.Open("mysql", dsn)
		MDB = mdb
		if err != nil {
			panic(fmt.Sprintf("Mysql 연결 실패: %v", err))
		}

		if err := MDB.Ping(); err != nil {
			panic(fmt.Sprintf("Mysql 연결 실패: %v", err))
		}

		log.Println("Mysql 연결 성공")
	})

	return MDB
}
