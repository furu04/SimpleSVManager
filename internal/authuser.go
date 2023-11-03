/*
 Copyright (C) 2023 furu04

 This program is free software: you can redistribute it and/or modify it under the terms of the GNU Affero General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

 This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Affero General Public License for more details.

 You should have received a copy of the GNU Affero General Public License along with this program. If not, see https://www.gnu.org/licenses/.
*/

package internal

import (
	"gopkg.in/ini.v1"
	"database/sql"
	"fmt"
	"time"
	"log"
	_"github.com/go-sql-driver/mysql"
)

func Authuser(username, password, mailaddress, ip string) (bool, string) {
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
	defer db.Close()

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	result, err := db.Query("SELECT * from users where username = ?", username)
	if err != nil {
		panic(err)
	}
	defer result.Close()

	if result.Next() {
		var username, mailaddress, hash, regip, roll string
		err := result.Scan(&username, &mailaddress, &hash, &regip, &roll)
		if err != nil {
			panic(err)
		}
		
		auth := CompareHashAndPassword(password, hash)
		if auth {
			return true, ""
		} else {
			return false, "パスワードが間違っています．"
		}

	} else {
		return false, "ユーザが存在しません．"
	}
}