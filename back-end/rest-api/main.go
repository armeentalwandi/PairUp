package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	//"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
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
	query := `CREATE TABLE IF NOT EXISTS buddies(id int primary key auto_increment, name text,
		email text, safety int, friendly int, startlat float, startlong float, endlat float, endlong float)`
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
	Id        string  `json:"Id"`
	Name      string  `json:"Name"`
	Email     string  `json:"Email"`
	Safety    int     `json:"Safety"`
	Friendly  int     `json:"Friendly"`
	StartLat  float64 `json:"StartLat"`
	StartLong float64 `json:"StartLong"`
	EndLat    float64 `json:"EndLat"`
	EndLong   float64 `json:"EndLong"`
}

var Users []User

func insert(db *sql.DB, w User) error {
	query := "INSERT INTO buddies(name, email, safety, friendly, startlat, startlong, endlat, endlong) VALUES (?, ?, ?, ?, ?, ?, ?, ?,?)"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return err
	}
	defer stmt.Close()
	res, err := stmt.ExecContext(ctx, w.Name, w.Email, w.Safety, w.Friendly, w.StartLat, w.StartLong, w.EndLat, w.EndLong)
	if err != nil {
		log.Printf("Error %s when inserting row into user table", err)
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return err
	}
	log.Printf("%d user created ", rows)
	return nil
}

func createUser(w http.ResponseWriter, r *http.Request) {
	db, err := dbConnection()
	fmt.Print(err)
	res, _ := ioutil.ReadAll(r.Body)
	var user User
	json.Unmarshal(res, &user)

	Users = append(Users, user)

	insForm, err := db.Prepare("INSERT INTO buddies(Id, Name, Email, StartLat , StartLong, EndLat, EndLong, Safety, Friendly) VALUES (?,?,?,?,?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	insForm.Exec(user.Id, user.Name, user.Email, user.StartLat, user.StartLong, user.EndLat, user.EndLong, user.Safety, user.Friendly)
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
			user.StartLat = updatedUser.StartLat
			user.StartLong = updatedUser.StartLong
			user.EndLat = updatedUser.EndLat
			user.EndLong = updatedUser.EndLong
			user.Email = updatedUser.Email
			user.Friendly = updatedUser.Friendly
			user.Name = updatedUser.Name
			user.Safety = updatedUser.Safety
			Users[index] = user
			json.NewEncoder(w).Encode(user)
		}
	}
}

func updateLocation(w http.ResponseWriter, r *http.Request) {
	// id variable problem here, tryna find id in query parameter
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Println("Id:")
	fmt.Println(id)
	var updatedUser User
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &updatedUser)
	for index, user := range Users {
		fmt.Println(user)
		if user.Id == updatedUser.Id {
			user.StartLat = updatedUser.StartLat
			user.StartLong = updatedUser.StartLong
			user.EndLat = updatedUser.EndLat
			user.EndLong = updatedUser.EndLong
			Users[index] = user
			json.NewEncoder(w).Encode(user)
			fmt.Println("User:")
			fmt.Println(user)
		}
	}
	fmt.Println("UpdatedUser:")
	fmt.Println(updatedUser)

}

func getUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Users)
}

// func returnResult(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	key := vars["id"]
// 	dist := vars["Dist"]
// 	//loop over all the users
// 	//if the user.Id equals the key we pass in
// 	//return the user encoded as JSON

// 	for _, user := range Users {
// 		i, err := strconv.ParseFloat(dist, 64)
// 		fmt.Print(err)
// 		if user.Id != key {
// 			if user.Dist-i <= 2 {
// 				json.NewEncoder(w).Encode(user)
// 			}

// 		}
// 	}
// }

func homePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "homepage vibes")
	fmt.Println("ENDPOINT: Homepage")
}

func Cors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=ascii")
	w.Header().Set("Secure", "false")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
}

func handleRequests() {
	router := mux.NewRouter()

	router.HandleFunc("/", homePage)
	router.HandleFunc("/users", getUsers)

	router.HandleFunc("/create", createUser).Methods("POST")
	router.HandleFunc("/location", updateLocation).Methods("PUT")
	router.HandleFunc("/users/{id}", updateUser).Methods("PUT")
	//router.HandleFunc("/results", returnResult)

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":10000", handler))
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

	Users = []User{}

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
