package main

import (
	conn "goPrj/src/conn"
	crud "goPrj/src/crud"
	"net/http"

	"github.com/gorilla/mux"
	  //"github.com/gin-gonic/gin"
)


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


	/*
	//use gin libs
	r := gin.Default() //We’re using the Default router because Gin provides some useful middlewares we can use to debug our server.

	r.GET("/", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"User": "hello world"})    
  })

  r.Run()
  */


}

/*
var(
	users = []crud.User{
		{
			Username: "Flynn",
			Password: "123",
		},
		{
			Username: "Dung",
			Password: "123",
		},
	}
	books = []crud.Book{
		{
			Title: "math",
			Author: "Newton",
			CallNumber: 1234,
		},
		{
			Title: "phisical",
			Author: "Newton",
			CallNumber: 5678,
		},
	}

	
)
*/


