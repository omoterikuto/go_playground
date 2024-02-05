package main

import (
	"fmt"
	"time"

	"github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println(dataSource("user", "password", "host", "dbName"))
}
func dataSource(userName string, password string, host string, dbName string) string {
	c := mysql.Config{
		DBName:    dbName,
		User:      userName,
		Passwd:    password,
		Addr:      host,
		Net:       "tcp",
		ParseTime: true,
		Collation: "utf8mb4",
		Loc:       time.Local,
	}
	return c.FormatDSN()
}
