package conndb

import (
	"fmt"
	"log"

	//"database/sql"
	//_ "github.com/lib/pq"
	"gorm.io/gorm"
  	"gorm.io/driver/postgres"

	//"github.com/jinzhu/gorm"
	"net/http"

	//"github.com/jinzhu/gorm/dialects/postgres"
	//_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	host = "localhost"
	port = 2345
	user = "postgres"
	password = "11235813"
	dbname = "mytestdb"
)
var Db *gorm.DB

func Conndb(){

	/* connect to db by lib/pq
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s database=%s sslmode=disable",host,port,user,password,dbname)
	db, err := sql.Open("postgres",connStr)
	if err != nil{
		panic(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil{
		panic(err)
	}

	fmt.Println("Successfully connected")
	*/

	//connect to db by gorm libs
	var err error
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s database=%s sslmode=disable",host,port,user,password,dbname)
	Db, err = gorm.Open(postgres.Open(connStr), &gorm.Config{}) // open db
	if err != nil{ // if can not connect to db, response err
		log.Fatal(err)
	} else {
		fmt.Println("Successfully connected")
	}
	
	/*
	err = Db.DB().Ping() // check connection to db still alive or not
	if err != nil{ // if can not connect to db, response err
		//panic(err) // panic ~ exception in C#
		log.Fatal(err)
	}
	*/
	
}

type HandleData interface{
	CreateData(w http.ResponseWriter, r *http.Request)
	UpdateData(w http.ResponseWriter, r *http.Request)
	DeleteData(w http.ResponseWriter, r *http.Request)
	GetAllData(w http.ResponseWriter, r *http.Request)
	GetSingleData(w http.ResponseWriter, r *http.Request)
}
