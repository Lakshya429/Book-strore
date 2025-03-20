package main

import (
	"fmt"
	"log"
	"github.com/gorilla/mux"
	"github/Lakshya429/postgress/pkg/routes"
	"net/http"
	"github/Lakshya429/postgress/pkg/models"
)

func main() {
	models.Setup()
	router := mux.NewRouter()
	routes.RouterRegistry(router)
	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}