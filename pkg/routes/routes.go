package routes

import (
	"github.com/gorilla/mux"
	"github/Lakshya429/postgress/pkg/controler"
)


var RouterRegistry = func( r *mux.Router){
	r.HandleFunc("/api/books" , controler.GetBooks).Methods("GET")
	r.HandleFunc("/api/books" , controler.CreateBook).Methods("POST")
	r.HandleFunc("/api/books/{id}" , controler.GetBook).Methods("GET")
	r.HandleFunc("/api/books/{id}" , controler.UpdateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}" , controler.DeleteBook).Methods("DELETE")
}