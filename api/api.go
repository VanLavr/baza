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



/* 
	connection to database...
*/
var (
	DB *sql.DB
	connectionErr error
)
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


/* 
	creating functions for
	endpoints...
*/ 
// struct -> JSON...
type BAZA struct {
	ID   int    `json:"ID"`
	Baza string `json:"baza"`
}

// select certain baza...
func GetBazaByID(c *gin.Context) {
	var BAZAS []BAZA
	id := c.Param("id")

	rows, rowsErr := DB.Query("SELECT * FROM BAZAS WHERE ID = ?", id)
	if rowsErr != nil {
		log.Fatalf("GetBazaByID: %d, %v", id, rowsErr)
	}
	defer rows.Close()

	for rows.Next() {
		var bz BAZA
		if scanErr := rows.Scan(&bz.ID, &bz.Baza); scanErr != nil {
			log.Fatalf("GetBazaByID: %d, %v", id, scanErr)
		}
		BAZAS = append(BAZAS, bz)
	}

	if rowErr := rows.Err(); rowErr != nil {
		log.Fatalf("GetBazaByID: %d, %v", id, rowErr)
	}

	c.IndentedJSON(http.StatusOK, BAZAS)
}

// select all bazas...
func GetAllBazas(c *gin.Context) {
	var BAZAS []BAZA

	rows, rowsErr := DB.Query("SELECT * FROM BAZAS")
	if rowsErr != nil {
		log.Fatalf("GetAllBazas: %v", rowsErr)
	}
	defer rows.Close()

	for rows.Next() {
		var bz BAZA
		if scanErr := rows.Scan(&bz.ID, &bz.Baza); scanErr != nil {
			log.Fatalf("GetAllBazas: %v", scanErr)
		}
		BAZAS = append(BAZAS, bz)
	}

	if rowErr := rows.Err(); rowErr != nil {
		log.Fatalf("GetAllBazas: %v", rowErr)
	}

	c.IndentedJSON(http.StatusOK, BAZAS)
}

// create your own baza...
func CreateYourBaza(c *gin.Context) {
	var MyBaza BAZA

	if bindErr := c.BindJSON(&MyBaza); bindErr != nil {
		log.Fatalf("CreateYourBaza: %v", bindErr)
	}
	_, InsertionErr := DB.Query("INSERT INTO BAZAS(ID, BAZA) VALUES(?, ?)", MyBaza.ID, MyBaza.Baza)
	if InsertionErr != nil {
		log.Fatalf("CreateYourBaza: %v", InsertionErr)
	}

	c.IndentedJSON(http.StatusCreated, MyBaza)
}