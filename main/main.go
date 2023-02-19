package main

import (
	"fmt"
	"BAZA/api"
	"github.com/gin-gonic/gin"
	//"net/http"
	//"os"
	//"github.com/go-sql-driver/mysql"
	"log"
	//"database/sql"
)

func main() {
	fmt.Printf("\n\n\n")
	api.Greeting()
	api.ConnectingToDataBase()
	
	MyBaza, bazaErr := api.GetBazaByID(3)
	if bazaErr != nil {
		log.Fatal(bazaErr)
	}
	fmt.Printf("baza by id 3: %v\n", MyBaza)
	
	fmt.Printf("\n\n\n")

	router := gin.Default()
	router.GET("/baza", api.GetTests)

	//router.POST()
	//router.DELETE()

	router.Run("localhost:8080")
}