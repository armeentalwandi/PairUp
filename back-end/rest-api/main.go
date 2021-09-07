package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

const (
	username = "root"
	password = "Gimshaz19!"
	hostname = "127.0.0.1:3306"
	dbname   = "walkers"
)

func dsn(dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
}

func dbConnection() (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn(""))
	if err != nil {
		log.Printf("Error %s when opening DB\n", err)
		return nil, err
	}
	//defer db.Close()

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	res, err := db.ExecContext(ctx, "CREATE DATABASE IF NOT EXISTS "+dbname)
	if err != nil {
		log.Printf("Error %s when creating DB\n", err)
		return nil, err
	}
	no, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when fetching rows", err)
		return nil, err
	}
	log.Printf("rows affected %d\n", no)

	db.Close()
	db, err = sql.Open("mysql", dsn(dbname))
	if err != nil {
		log.Printf("Error %s when opening DB", err)
		return nil, err
	}
	//defer db.Close()

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	ctx, cancelfunc = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	err = db.PingContext(ctx)
	if err != nil {
		log.Printf("Errors %s pinging DB", err)
		return nil, err
	}
	log.Printf("Connected to DB %s successfully\n", dbname)
	return db, nil

}

func createUsersTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS students(id int primary key auto_increment, name text,
		email text, safety int, friendly int, dist float)`
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	res, err := db.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when creating product table", err)
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when getting rows affected", err)
		return err
	}
	log.Printf("Rows affected when creating table: %d", rows)
	return nil
}

type User struct {
	Id       string  `json:"Id"`
	Name     string  `json:"Name"`
	Email    string  `json:"Email"`
	Safety   int     `json:"Safety"`
	Friendly int     `json:"Friendly"`
	Dist     float64 `json:"Dist"`
}

var Users []User

// func insert(db *sql.DB, w User) error {
// 	query := "INSERT INTO students(name, email, safety, friendly, dist) VALUES (?, ?, ?, ?, ?)"
// 	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancelfunc()
// 	stmt, err := db.PrepareContext(ctx, query)
// 	if err != nil {
// 		log.Printf("Error %s when preparing SQL statement", err)
// 		return err
// 	}
// 	defer stmt.Close()
// 	res, err := stmt.ExecContext(ctx, w.name, w.email, w.safety, w.friendly, w.dist)
// 	if err != nil {
// 		log.Printf("Error %s when inserting row into user table", err)
// 		return err
// 	}
// 	rows, err := res.RowsAffected()
// 	if err != nil {
// 		log.Printf("Error %s when finding rows affected", err)
// 		return err
// 	}
// 	log.Printf("%d user created ", rows)
// 	return nil
// }

func createUser(w http.ResponseWriter, r *http.Request) {
	res, _ := ioutil.ReadAll(r.Body)
	var user User
	json.Unmarshal(res, &user)
	Users = append(Users, user)
	json.NewEncoder(w).Encode(user)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var updatedUser User
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &updatedUser)
	for index, user := range Users {
		if user.Id == id {
			user.Dist = updatedUser.Dist
			user.Email = updatedUser.Email
			user.Friendly = updatedUser.Friendly
			user.Name = updatedUser.Name
			user.Safety = updatedUser.Safety
			Users[index] = user
			json.NewEncoder(w).Encode(user)
		}
	}

}
func getUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Users)
}

func returnUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	dist := vars["Dist"]
	//loop over all the users
	//if the user.Id equals the key we pass in
	//return the user encoded as JSON

	for _, user := range Users {
		i, err := strconv.ParseFloat(dist, 64)
		fmt.Print(err)
		if user.Id != key {
			if user.Dist-i <= 2 {
				json.NewEncoder(w).Encode(user)
			}

		}
	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "homepage vibes")
	fmt.Println("ENDPOINT: Homepage")
}

func handleRequests() {
	router := mux.NewRouter()
	router.HandleFunc("/", homePage)
	router.HandleFunc("/users", getUsers)

	router.HandleFunc("/create", createUser).Methods("POST")
	router.HandleFunc("/users/{id}", updateUser).Methods("PUT")
	router.HandleFunc("/results", returnUser)
	log.Fatal(http.ListenAndServe(":10000", router))
}

func main() {
	db, err := dbConnection()
	if err != nil {
		log.Printf("Error %s when getting db connection", err)
		return
	}
	defer db.Close()

	log.Printf("Successfully connected to database")
	err = createUsersTable(db)
	if err != nil {
		log.Printf("table failed with %s", err)
		return
	}

	Users = []User{
		User{Name: "Neha", Email: "n123@gmail.com", Safety: 5, Friendly: 5, Dist: 1.5}}

	handleRequests()

}

// w := user{
// 	name:     "neha",
// 	email:    "n123@gmail.com",
// 	safety:   5,
// 	friendly: 5,
// 	dist:     1.5,
// }
// err = insert(db, w)
// if err != nil {
// 	log.Printf("Insert product failed with error %s", err)
// 	return
// }
