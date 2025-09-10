package mysqlclient

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var MDB *sql.DB

func InitMysql(user, password, host, dbname, port string) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, dbname)

	MDB, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(fmt.Sprintf("Mysql 연결 실패: %v, %v", err, dsn))
	}

	if err := MDB.Ping(); err != nil {
		panic(fmt.Sprintf("Mysql 연결 실패: %v", err))
	}
	log.Println("Mysql 연결 성공")
}
