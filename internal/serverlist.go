/*
 Copyright (C) 2023 furu04

 This program is free software: you can redistribute it and/or modify
 it under the terms of the GNU Affero General Public License as published
 by the Free Software Foundation, either version 3 of the License, or
 (at your option) any later version.

 This program is distributed in the hope that it will be useful,
 but WITHOUT ANY WARRANTY; without even the implied warranty of
 MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 GNU Affero General Public License for more details.

 You should have received a copy of the GNU Affero General Public License
 along with this program.  If not, see <https://www.gnu.org/licenses/>.

*/

package internal

import (
	"database/sql"
	"time"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/ini.v1"
)

func serverlist(username string, pages int) ([]byte, error) {
	// DBに格納されているサーバの情報一覧を取得する．サーバの詳細は別関数で取得できるようにする．
}