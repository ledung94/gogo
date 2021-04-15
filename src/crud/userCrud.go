package crud

import (
	"encoding/json"
	//"fmt"
	conn "goPrj/src/conn"
	"net/http"

	"github.com/gorilla/mux"
	//"github.com/jinzhu/gorm"
	"gorm.io/gorm"
	//"gorm.io/driver/postgres"
)

type User struct{
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	Books []Book `json:"books"`
}

func GetAllUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	var users []User
	conn.Db.Find(&users)
	json.NewEncoder(w).Encode(&users)
}

func GetSingleUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	var user User
	var books []Book
	params := mux.Vars(r)
	conn.Db.First(&user,params["id"])
	conn.Db.Model(&user).Association("Books").Find(&books)
	user.Books = books
	json.NewEncoder(w).Encode(&user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	var user User
	param := mux.Vars(r)
	conn.Db.First(&user,param["id"])
	json.NewDecoder(r.Body).Decode(&user)
	conn.Db.Save(&user)
	json.NewEncoder(w).Encode(&user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user User
	conn.Db.Delete(&user, params["id"])
	json.NewEncoder(w).Encode("The USer is Deleted Successfully!")
}

func CreateUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	conn.Db.Create(&user)
	conn.Db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&user)
	//conn.Db.Save(&user)
	json.NewEncoder(w).Encode(&user)
}
/*
func CreateUser(users []User) {
	
	for index := range users {
		conn.Db.Create(&users[index])
	}
}
*/
/*
func CreateUser(users []User) {
	//make migration to the db if they have not already to create
	conn.Db.Create(&users)
}
*/