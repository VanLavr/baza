package main

import (
	"fmt"
	"BAZA/api"
	"github.com/gin-gonic/gin"
	//"net/http"
	//"os"
	//"github.com/go-sql-driver/mysql"
	//"log"
	//"database/sql"
)

func main() {
	fmt.Printf("\n\n\n")
	api.Greeting()
	api.ConnectingToDataBase()
	fmt.Printf("\n\n\n")

	
	router := gin.Default()
	// get all baza endpoint...
	router.GET("/baza", api.GetAllBazas)
	// get a certain baza endpoint...
	router.GET("/baza/:id", api.GetBazaByID)
	// create your own baza endpoint...
	router.POST("/baza/create", api.CreateYourBaza)
	// delete baza endpoint... 
	router.DELETE("baza/delete/:id", api.DeleteBaza)

	router.Run("localhost:8080")
}