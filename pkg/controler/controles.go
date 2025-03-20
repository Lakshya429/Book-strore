package controler

import (
	"github/Lakshya429/postgress/pkg/models"
	"net/http"
	"strconv"
	"encoding/json"
	"github.com/gorilla/mux"
)

func GetBooks(w http.ResponseWriter, r *http.Request){
	books := models.GetAllBooks()
	res , _ := json.Marshal(books)
	w.Header().Add("Content-Type" , "application/json")
    w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBook(w http.ResponseWriter , r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]
	ID , _ := strconv.Atoi(id)

	book := models.GetBook(ID)
	res , _ := json.Marshal(book)
	w.Header().Add("Content-Type" , "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w  http.ResponseWriter , r *http.Request){
	book := &models.Book{}
	json.NewDecoder(r.Body).Decode(book)
	book = models.CreateBook(book)
	res , _ := json.Marshal(book)
	w.Header().Add("Content-Type" , "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter , r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]
	Id  , _ := strconv.Atoi(id)

	book := &models.Book{}
	json.NewDecoder(r.Body).Decode(book)
	book = models.UpdateBook(Id , book)
	res , _ := json.Marshal(book)
	w.Header().Add("Content-Type" , "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func DeleteBook(w http.ResponseWriter , r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]
	ID , _ := strconv.Atoi(id)
	models.DeleteBook(ID)
}
