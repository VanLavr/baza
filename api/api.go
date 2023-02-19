package api

import (
	"fmt"
	//"os"
	"github.com/go-sql-driver/mysql"
	"log"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Greeting() {
	fmt.Println("Working!")
}

type test struct {
	ID   int    `json:"ID"`
	Baza string `json:"baza"`	
}

var Tests = []test {
	{ID: 5, Baza: "test"},
}

func GetTests(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Tests)
}



type BAZA struct {
	ID   int    `json:"ID"`
	Baza string `json:"baza"`
}

var (
	DB *sql.DB
	connectionErr error
)

// connection to database...
func ConnectingToDataBase() {
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
		fmt.Println("connected...")
	}	
}

func GetBazaByID(id int) ([]BAZA, error) {
	var BAZAS []BAZA

	rows, rowsErr := DB.Query("SELECT * FROM BAZAS WHERE ID = ?", id)
	if rowsErr != nil {
		return nil, fmt.Errorf("GetBazaByID: %d, %v", id, rowsErr)
	}
	defer rows.Close()

	for rows.Next() {
		var bz BAZA
		if scanErr := rows.Scan(&bz.ID, &bz.Baza); scanErr != nil {
			return nil, fmt.Errorf("GetBazaByID: %d, %v", id, scanErr)
		}
		BAZAS = append(BAZAS, bz)
	}

	if rowErr := rows.Err(); rowErr != nil {
		return nil, fmt.Errorf("GetBazaByID: %d, %v", id, rowErr)
	}

	return BAZAS, nil
}