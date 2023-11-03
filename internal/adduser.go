package internal

import (
	"database/sql"
	"time"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/ini.v1"
)

func Add(username, password, mailaddress, ip, roll string) (bool, string) {
	cfg, err := ini.Load("../configs/config.ini")
	//iniファイルの場所を指定する
	if err != nil {
		panic(err)
	}

	if len(password) >= 72 || len(password) <= 6 {
		log.Fatalf("72文字以上，または6文字以下のパスワードは設定できません．")
	}
	if len(username) > 25 {
		log.Fatalf("25文字以上のユーザー名は設定できません")
	}
	if len(mailaddress) > 50 {
		log.Fatalf("50文字以上のメールアドレスは入力できません")
	}
	
	DB_TYPE := cfg.Section("database").Key("DB_TYPE").String()
	DB_HOST := cfg.Section("database").Key("DB_HOST").String()
	DB_PORT := cfg.Section("database").Key("DB_PORT").String()
	DB_NAME := cfg.Section("database").Key("DB_NAME").String()
	DB_USER := cfg.Section("database").Key("DB_USER").String()
	DB_PASSWORD := cfg.Section("database").Key("DB_PASSWORD").String()
	
	hash, err := Hashing(password)

	if err != nil {
		panic(err)
	}

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

	_, err = db.Exec("INSERT INTO users (username, mailaddress, hash, regip, roll) VALUES (?, ?, ?, ?, ?) ", username, mailaddress, hash, ip, roll)
	
	if err != nil {
		panic(err)
	}

	defer db.Close()

	// ../scripts/init.goを完成させ，DBの仕様を完全に決めたらユーザー名，メールアドレス重複時のエラーメッセージを返すようにする

	return true, ""
}