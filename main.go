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

	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(maxConnections)
	db.SetMaxIdleConns(maxConnections)
	
	return db
}


//testdata 

type user struct {
	rowid int `json:"rowid"`
	name  string `json:"name"`
	pw    string `json:"pw"`
	email string `json:"email"`
}

var users = []user{
	{rowid: 1, name: "testuser", pw:"efjpokewfnewpofewf", email:"test@test.com"},
}

func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

func main() {
	router := gin.Default()
	router.GET("/user", getUsers)
	fmt.Println("Starting Rest Api on port 8080")
	router.Run("127.0.0.1:8080")
}