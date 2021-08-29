package main

import (  
    "context"
    "database/sql"
    "fmt"
    "log"
    "time"
	_ "github.com/go-sql-driver/mysql"
)

const (  
    username = "root"
    password = "TECHnova123"
    hostname = "127.0.0.1:3306"
    dbname   = "walkers"
)

func dsn(dbName string) string {  
    return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
}

func dbConnection() (*sql.DB, error){  
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


type user struct {
	name string
	email string
	safety int
	friendly int
    dist float64
}
func insert(db *sql.DB, w user) error {  
	query := "INSERT INTO students(name, email, safety, friendly, dist) VALUES (?, ?, ?, ?, ?)"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return err
	}
	defer stmt.Close()
	res, err := stmt.ExecContext(ctx, w.name, w.email, w.safety, w.friendly, w.dist)
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

func main() {
	db, err := dbConnection()
	if err != nil {
		log.Printf("Error %s when getting db connection", err)
		return
	}
	defer db.Close()

	log.Printf("Successfully connected to database")
	w := user {  
        name:  "neha",
		email: "n123@gmail.com",
        safety: 5,
		friendly: 5,
		dist: 1.5,
    }
	err = insert(db, w)  
	if err != nil {  
		log.Printf("Insert product failed with error %s", err)
		return
    }
}


