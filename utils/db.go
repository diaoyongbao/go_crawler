package utils

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"time"
)

var (
	userName  string = "root"
	password  string = "Jwt@1234"
	ipAddrees string = "172.18.63.145"
	port      int    = 3306
	dbName    string = "crawler"
	charset   string = "utf8"
)
func ConnectMysql() *sqlx.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", userName, password, ipAddrees, port, dbName, charset)
	Db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("mysql connect failed, detail is [%v]", err.Error())
	}
	return Db
}

func ping(Db *sqlx.DB) {
	err := Db.Ping()
	if err != nil {
		fmt.Println("ping failed")
	} else {
		fmt.Println("ping success")
	}
}

//将成功的数据写入数据库
func AddProxy(Db *sqlx.DB,p string) {
	result, err := Db.Exec("insert into proxy(proxy,checkDate)  values(?,?)",p, time.Now())
	if err != nil {
		fmt.Printf("data insert faied, error:[%v]", err.Error())
		return
	}
	id, _ := result.LastInsertId()
	fmt.Printf("insert success, last id:[%d]\n", id)
}