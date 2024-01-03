package persistence

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const DRIVER = "mysql"
const USER = `u528869443_oseiasc2kedis`
const PASS = "oseiasc2Costa1M$"
const DBNAME = "u528869443_kedis_db"

//const HOSTNAME = "br-asc-web815.main-hosting.eu"

func Connect() *sql.DB {
	URL := fmt.Sprintf(`%s:%s@tcp(82.180.153.52)/%s`, USER, PASS, DBNAME)
	con, err := sql.Open(DRIVER, URL)
	if err != nil {
		panic(err.Error())
	}

	return con
}
