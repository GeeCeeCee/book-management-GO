package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/GeeCeeCee/bookings/pkg/models"
	"github.com/GeeCeeCee/bookings/pkg/utils"
	"github.com/gorilla/mux"
)

var newBook models.Book

type Error struct {
	Error string `json:"error"`
}


func GetBook(w http.ResponseWriter, r *http.Request) {
	newBook := models.GetAllBooks()
	res, _ := json.Marshal(newBook)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookID := vars["bookID"]
	ID, err := strconv.ParseInt(bookID, 0, 0)

	if err != nil {
		fmt.Println("Error parsing")
	}

	bookDetails, _ := models.GetBookByID(ID)

	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)


	if res, _ := json.Marshal(bookDetails); (*bookDetails == models.Book{}) {
		bookDetailsError := &Error{"No such book exists."}
		res, _ := json.Marshal(bookDetailsError)
		w.Write(res)
	} else {
		w.Write(res)
	}
}


func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b:= CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {

	vars:= mux.Vars(r)

	bookID := vars["bookID"]
	ID, err := strconv.ParseInt(bookID, 0, 0)

	if err != nil {
		fmt.Println("There was an error in parsing the book details")
	}


	b:= models.DeleteBook(ID)
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)

	ID:= updateBook.ID

	bookDetails, _ := models.GetBookByID(ID)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if updateBook.ID != bookDetails.ID {
		err := Error{"The book is not present in the database"}
		res, _ := json.Marshal(err)
		w.Write(res)
		return
	}
	
	models.DeleteBook(updateBook.ID)

	updateBook.CreateBook()

	allBooks := models.GetAllBooks()

	res, _ := json.Marshal(allBooks)
	
	w.Write(res)

}