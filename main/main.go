package main

import (
	"fmt"
	"BAZA/api"
	"github.com/gin-gonic/gin"
	"net/http"
	//"os"
	//"github.com/go-sql-driver/mysql"
	"log"
	//"database/sql"
)

func main() {
	fmt.Printf("\n\n\n")
	api.Greeting()
	api.ConnectingToDataBase()
	fmt.Printf("\n\n\n")

	
	router := gin.Default()
	
	// get all endpoint...
	MyBaza, bazaErr := api.GetAllBazas()
	if bazaErr != nil {
		log.Fatal(bazaErr)
	}
	router.GET("/baza", func (c *gin.Context) {
		c.IndentedJSON(http.StatusOK, MyBaza)
	})

	//router.POST()
	//router.DELETE()

	router.Run("localhost:8080")
}