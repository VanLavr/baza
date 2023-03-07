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
	working with database...
*/
// connection to database...
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
		fmt.Println("did not connected...")
	}	
}

// get all ids from database...
func GetAllIds() (allIds []int) {
	ids, idsErr := DB.Query("SELECT id FROM BAZAS")
	if idsErr != nil {
		log.Fatalf("GetAllIds: %v", idsErr)
	}

	for ids.Next() {
		var idFromDataBase int
		if scanErr := ids.Scan(&idFromDataBase); scanErr != nil {
			log.Fatalf("GetAllIds: %v", idsErr)
		}

		allIds = append(allIds, idFromDataBase)
	}

	return
}


/* 
	creating CRUD controller...
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
		log.Fatalf("CreateYourBaza (can't bind...): %v", bindErr)
		fmt.Println(MyBaza)
	}
	_, InsertionErr := DB.Query("INSERT INTO BAZAS(ID, BAZA) VALUES(?, ?)", MyBaza.ID, MyBaza.Baza)
	if InsertionErr != nil {
		log.Fatalf("CreateYourBaza (can't execute query...): %v", InsertionErr)
	}

	c.IndentedJSON(http.StatusCreated, MyBaza)
}

// delete baza...
func DeleteBaza(c *gin.Context) {
	id := c.Param("id")

	_, DeleteError := DB.Query("DELETE FROM BAZAS WHERE ID = ?", id)
	if DeleteError != nil {
		log.Fatalf("DeleteBaza (can't execute query...): %v", DeleteError)
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"status": "deleted",
	})
}