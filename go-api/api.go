package main

import (
	"fmt"
	"go-api/models"
	"go-api/routes"
	"log"
	"net/http"
)

func main() {
	port := "9898"
	models.TestConnection()
	fmt.Printf("Api corriendo en puerto %s\n", port)
	r := routes.NewRouter()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
}
