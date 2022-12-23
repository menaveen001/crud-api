package database

import (
	"curd-api/entity"
	"database/sql"
	"fmt"
	"log"
	"os"
)

//make connection to database

var db *sql.DB

func CreatConnectionTodb() {
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	databaseName := os.Getenv("DB_NAME")
	var err error
	db, err = sql.Open("mysql", username+":"+password+"@tcp(localhost:3306)/"+databaseName)
	if err != nil {

		fmt.Println("error validating sql.Open argument")

	}

	log.Println("successful connection to Database")

}

//create data into database

func InsertUser(u entity.User) int64 {

	insForm, err := db.Prepare("INSERT INTO userdata(id,name,dob,address,description,createdat) VALUES(?,?,?,?,?,?)")
	if err != nil {
		fmt.Println(err)
		return 0
	}

	//execute query

	res, err := insForm.Exec(u.Id, u.Name, u.Dob, u.Address, u.Description, u.CreatedAt)
	if err != nil {
		fmt.Println("query not executed", err)
		return 0
	}

	rows, err := res.RowsAffected()
	if err != nil {
		fmt.Println("rows not inserted", err)
		return 0
	}
	log.Printf("rows inserted,%v\n", rows)
	return rows

}

// get all data from database

func SelectData() []entity.User {

	var q []entity.User
	queryRes, err := db.Query("SELECT * FROM userdata")
	if err != nil {
		fmt.Println(err, queryRes)
	}
	for queryRes.Next() {

		var s entity.User
		err := queryRes.Scan(&s.Id, &s.Name, &s.Dob, &s.Address, &s.Description, &s.CreatedAt)

		if err != nil {
			fmt.Println(err)
		}
		q = append(q, s)
	}
	log.Println(q)
	return q
}

// get user data by id from database
func SelectDataByID(id int) []entity.User {

	var user1 []entity.User
	queryRes, err := db.Query("SELECT *FROM userdata WHERE id =?", id)
	if err != nil {
		fmt.Println(err, queryRes)
	}

	if queryRes.Next() {
		var user entity.User
		err := queryRes.Scan(&user.Id, &user.Name, &user.Dob, &user.Address, &user.Description, &user.CreatedAt)
		if err != nil {
			log.Println(err)

		}
		user1 = append(user1, user)

	}

	log.Println(user1)
	return user1

}

// update user data by id in database

func UpdateUserById(p entity.User) int64 {
	updateRes, err := db.Prepare("update userdata set name = ?, dob =?, address =?, description=?, createdt=? where id = ?")
	if err != nil {
		fmt.Println(err)
	}

	// execute query
	res, err := updateRes.Exec(p.Id, p.Name, p.Dob, p.Address, p.Description, p.CreatedAt)
	if err != nil {
		fmt.Println(err, res)
	}
	a, err := res.RowsAffected()
	if err != nil {
		fmt.Println("data is not updated", err, a)

	}

	fmt.Printf("data is updated ,%v\n", a)
	log.Println(a)
	return 0

}

// delete data by id from data base
func DeleteDataByID(id int) []entity.User {

	var user1 []entity.User
	queryRes, err := db.Query("delete from userdata where id = ?", id)
	if err != nil {
		fmt.Println(err, queryRes)
	}

	if queryRes.Next() {
		var user entity.User
		err := queryRes.Scan(&user.Id, &user.Name, &user.Dob, &user.Address, &user.Description, &user.CreatedAt)
		if err != nil {
			log.Println(err)

		}
		user1 = append(user1, user)

	}

	log.Println(user1)
	return user1
}
