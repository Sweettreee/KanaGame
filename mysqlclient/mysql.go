package mysqlclient

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type _StmtList struct {
	InsertUser *sql.Stmt
	LoginUser  *sql.Stmt
}

var StmtList _StmtList
var MDB *sql.DB

func InitMysql() *sql.DB {
	GetMysqlConnection()

	var err error

	StmtList.InsertUser, err = MDB.Prepare("INSERT INTO users (username, password) VALUES (?, ?)")
	if err != nil {
		log.Fatalf("InsertUser 문 준비 실패: %v", err)
	}

	StmtList.LoginUser, err = MDB.Prepare("SELECT id, username, password FROM users WHERE username = ?")
	if err != nil {
		log.Fatalf("LoginUser 문 준비 실패: %v", err)
	}

	return MDB
}

func GetMysqlConnection() *sql.DB {

	if MDB != nil {
		return MDB
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_DB"))

	MDB, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(fmt.Sprintf("Mysql 연결 실패: %v", err))
	}

	if err := MDB.Ping(); err != nil {
		panic(fmt.Sprintf("Mysql 연결 실패: %v", err))
	}
	log.Println("Mysql 연결 성공")
	return MDB
}
