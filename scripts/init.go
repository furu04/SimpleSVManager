package scripts

import (
	"database/sql"
	"time"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/ini.v1"
)

func init() {
	cfg, err := ini.Load("../configs/config.ini")
	if err != nil {
		panic(err)
	}

	DB_TYPE := cfg.Section("database").Key("DB_TYPE").String()
	DB_HOST := cfg.Section("database").Key("DB_HOST").String()
	DB_PORT := cfg.Section("database").Key("DB_PORT").String()
	DB_NAME := cfg.Section("database").Key("DB_NAME").String()
	DB_USER := cfg.Section("database").Key("DB_USER").String()
	DB_PASSWORD := cfg.Section("database").Key("DB_PASSWORD").String()

	if DB_TYPE != "mysql" {
		log.Fatalf("サポートされていないDBを利用しています．")
	}

	authInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)

	db, err := sql.Open("mysql", authInfo)
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (username varchar(25), mailaddress varchar(50), hash text, regip text, roll text)")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	// その他のテーブルも機能を実装する毎に作成する
}