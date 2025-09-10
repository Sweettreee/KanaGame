package mysqlclient

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var MDB *sql.DB

func initMysql() {
	MDB, err := sql.Open("mysql", "root@tcp(localhost:3306)/KanaGame")
	if err != nil {
		panic(fmt.Sprintf("Mysql 연결 실패: %v", err))
	}

	if err := MDB.Ping(); err != nil {
		panic(fmt.Sprintf("Mysql 연결 실패: %v", err))
	}
	log.Println("Mysql 연결 성공")
}
