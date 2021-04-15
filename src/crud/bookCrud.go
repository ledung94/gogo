package crud

import (
	"encoding/json"
	conn "goPrj/src/conn"
	"net/http"

	"gorm.io/gorm"
	//"github.com/jinzhu/gorm"
)

type Book struct{
	gorm.Model
	Title string `json:"title"`
	Author string `json:"author"`
	CallNumber int `json:"callnumber"`
	UserID int `json:"userid"`
}


func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	var book Book
	json.NewDecoder(r.Body).Decode(&book)
	conn.Db.Create(&book)
	json.NewEncoder(w).Encode(&book)
}

func GetBooks(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	var books []Book
	conn.Db.Find(&books)
	json.NewEncoder(w).Encode(&books)
}