package demo_mysql

import (
	"database/sql"
	"fmt"
	"log"
	"testing"
)

const (
	username = "hzj"
	password = "Mysql@123"
	hostname = "127.0.0.1:3306"
	dbname   = "ttbb"
)

func dsn(dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
}

/*
https://golangbot.com/connect-create-db-mysql/
*/
func TestDbName(t *testing.T) {
	db, err := sql.Open("mysql", dsn(""))
	if err != nil {
		log.Printf("Error %s when opening DB\n", err)
		return
	}
	defer db.Close()
}
