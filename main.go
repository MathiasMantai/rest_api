package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"database/sql"
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
	dbError(err)
	c.IndentedJSON(http.StatusOK, users)
}

func getUserById(c *gin.Context) {
	//get url param
	id := c.Param("id")

	db := dbConn(10)
	defer dbClose(db)
	smt, err := db.Prepare("SELECT * FROM users WHERE rowid = ?")
	dbError(err)
	defer smt.Close()
	smt.Exec(id)
}

func setUser(c *gin.Context) {
	var newUser user

	if err := c.BindJSON(&newUser); err != nil {
		return 
	}

	db := dbConn(10)
	defer dbClose(db)
	smt, err := db.Prepare("INSERT INTO users (name, pw, email) VALUES (?, ?, ?)")
	dbError(err)
	defer smt.Close()
	smt.Exec(newUser)
}


//main function
func main() {
	router := gin.Default()
	//define routes
	router.GET("/user", getUsers)
	router.GET("/user/:id", getUserById)
	router.POST("/user", setUser)
	fmt.Println("Starting Rest Api on port 8080")
	//start server
	router.Run("127.0.0.1:8080")
}
