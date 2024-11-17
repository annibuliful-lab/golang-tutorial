package main

import (
	"backend/src/db"
	"fmt"
	"net/http"
)

func main(){
	mux := http.NewServeMux()
	db,dbError := db.GetSQLiteClient("./dev.db")
	if dbError != nil{
		panic(dbError)
	}

	fmt.Println(db)

	port := "8080"
	fmt.Printf("Server running at http://localhost:%s/\n", port)
	err := http.ListenAndServe(":"+port, mux)

	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}