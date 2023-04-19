package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"database/sql"
	"github.com/go-sql-driver/mysql"
)

func dbConn(maxConnections int) *sql.DB {
	db, err := sql.Open("mysql", "root:@/userdata")
	dbError(err)
	db.SetMaxOpenConns(maxConnections)
	db.SetMaxIdleConns(maxConnections)
	
	return db
}

func dbClose(db *sql.DB) {
	db.Close()
}

func dbError(err error) {
	if(err != nil) {
		panic(err)
	}
}


//testdata 

type user struct {
	rowid int `json:"rowid"`
	name  string `json:"name"`
	pw    string `json:"pw"`
	email string `json:"email"`
}

// var users = []user{
// 	{rowid: 1, name: "testuser", pw:"efjpokewfnewpofewf", email:"test@test.com"},
// }



func getUsers(c *gin.Context) {
	db := dbConn(10)
	defer dbClose(db)
	smt, err := db.Prepare("SELECT * FROM users")
	dbError(err)
	defer smt.Close()
	users, err := smt.Exec()
	c.IndentedJSON(http.StatusOK, users)
}


//main function
func main() {
	router := gin.Default()
	//define routes
	router.GET("/user", getUsers)
	fmt.Println("Starting Rest Api on port 8080")
	//start server
	router.Run("127.0.0.1:8080")
}