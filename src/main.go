package main

import (
	"backend/src/db"
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main(){
	err := godotenv.Load("./.env")
	if err != nil {
	  log.Fatal("Error loading .env file")
	}
	
	mux := http.NewServeMux()
	db,dbError := db.GetSQLiteClient("./dev.db")
	if dbError != nil{
		panic(dbError)
	}

	fmt.Println(db)

	port := "8080"
	fmt.Printf("Server running at http://localhost:%s/\n", port)
	err = http.ListenAndServe(":"+port, mux)

	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}