package sqldb

import "database/sql"

// ConnectDB opens a connection to the database
func ConnectDB() *sql.DB {
	db, err := sql.Open("mysql", "chintan:chintan@123@tcp(192.168.64.2)/crmMaster")
	if err != nil {
		panic(err.Error())
	}

	return db
}
