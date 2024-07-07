package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "go_user:gouser1234@(127.0.0.1:3306)/go_api?parseTime=true");

	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	} 

	defer db.Close()
    
	errPing := db.Ping()
	if errPing != nil {
		log.Fatalf("Error pinging database: %v", errPing)
	} else {
		fmt.Println("Database connected and pinged successfully")
	}
	name := "riyad";
	email := "riyadkarim@gmail.com";
	phoneNumber := "01732461622" 

	createUser, err := db.Exec(`INSERT INTO users(name, email, phoneNumber) VAlUES (?,?,?)`, name, email, phoneNumber);

	if err!= nil {
		panic("user not created")
	}

	userID, err := createUser.LastInsertId() 

	if err!= nil {
		panic("User _id not define")
	}
	println("Last user ", userID)

	fmt.Println("Get all the user"); 

	type user struct{
		Name string `json:"name"`;
		Email string `json:"email"`;
		PhoneNumber string `json:"phoneNumber"`;
	}

	// find the query 
	rows, err := db.Query(`SELECT name, email, phoneNumber FROM users`);
	if err!= nil {
		panic("Error on query data")
	}
	defer rows.Close() 

	var users []user 

	for rows.Next() {
		var singleUser user
		err := rows.Scan(&singleUser.Name, &singleUser.Email, &singleUser.PhoneNumber);

		if err != nil {
			panic("user not valid")
		}
		users = append(users, singleUser)

	}

	fmt.Print("All users \n", users)

	newErr := rows.Err()

	if newErr != nil {
		panic("Create a last variable ")
	}


}