package validator

import (
	"log"
	"errors"
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"fmt"
)

var (
	DB *sql.DB
	connectionErr error
)

func Connecting() {
	var cfg = mysql.Config {
		User:                 "root",
		Passwd:               "12345",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "baza",
		AllowNativePasswords: true,
	}
	
	DB, connectionErr = sql.Open("mysql", cfg.FormatDSN())
	if connectionErr != nil {
		log.Fatal(connectionErr)
	}
	
	pingErr := DB.Ping()
	if pingErr != nil {
		fmt.Println("did not connected...")
	}
}

func IsUniqueId(id int) error {
	Connecting()

	var idsArray []int

	ids, idsErr := DB.Query("SELECT id FROM BAZAS")
	if idsErr != nil {
		log.Fatalf("GetAllIds: %v", idsErr)
	}

	for ids.Next() {
		var idFromDataBase int
		if scanErr := ids.Scan(&idFromDataBase); scanErr != nil {
			log.Fatalf("GetAllIds: %v", idsErr)
		}

		idsArray = append(idsArray, idFromDataBase)
	}

	for i := 0; i < len(idsArray); i++ {
		if id == idsArray[i] {
			return errors.New("Id you checked out is not unique!")
		}
	}

	return nil
}
