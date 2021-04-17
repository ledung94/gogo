package main

import (
	"fmt"
	conn "goPrj/src/conn"
	crud "goPrj/src/crud"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	//"github.com/gin-gonic/gin"
)

/*
func main(){

	//connect database
	conn.Conndb()
	//close database when this function finish
	//defer conn.Db.Close() // in gorm v1 have close func but in v2 not have
	sqlDB, err := conn.Db.DB()
	if err != nil {
		panic(err)
	}
	defer sqlDB.Close()
	//make migration to the db if they have not already to create
	conn.Db.AutoMigrate(&crud.User{})
	conn.Db.AutoMigrate(&crud.Book{})

	//API router
	//crud.CreateUser(users)

	//conn.Db.Create(&users) //dung khi bien users khong phai la slice

	//use mux to handle router
	router := mux.NewRouter()
	router.HandleFunc("/users",crud.CreateUser).Methods("POST")
	router.HandleFunc("/users",crud.GetAllUser).Methods("GET")
	router.HandleFunc("/books",crud.CreateBook).Methods("POST")
	router.HandleFunc("/books",crud.GetBooks).Methods("GET")
	router.HandleFunc("/users/{id}",crud.GetSingleUser).Methods("GET")
	router.HandleFunc("/users/{id}",crud.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}",crud.DeleteUser).Methods("DELETE")
	http.ListenAndServe(":8080",router)



	//use gin libs
	r := gin.Default() //We’re using the Default router because Gin provides some useful middlewares we can use to debug our server.

	r.GET("/", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"User": "hello world"})
  })

  r.Run()



}
*/
func main(){
	//connect database
	conn.Conndb()
	//close database when this function finish
	//defer conn.Db.Close() // in gorm v1 have close func but in v2 not have
	sqlDB, err := conn.Db.DB()
	if err != nil {
		panic(err)
	}
	defer sqlDB.Close()
	//make migration to the db if they have not already to create
	conn.Db.AutoMigrate(&crud.User{})
	conn.Db.AutoMigrate(&crud.Book{})

	//API
	handleRequest()

}

func handleRequest(){
	router := mux.NewRouter()
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./src/static/"))))
	router.HandleFunc("/", homePage)
	router.HandleFunc("/users",crud.CreateUser).Methods("POST")
	router.HandleFunc("/login",crud.Login).Methods("POST")
	router.HandleFunc("/users",crud.GetAllUser).Methods("GET")
	router.HandleFunc("/books",crud.CreateBook).Methods("POST")
	router.HandleFunc("/books",crud.GetBooks).Methods("GET")
	router.HandleFunc("/users/{id}",crud.GetSingleUser).Methods("GET")
	router.HandleFunc("/users/{id}",crud.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}",crud.DeleteUser).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080",router))
}

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w,"Welcome to the Home page")
	fmt.Println("Endpoint Hit: Home page")
}