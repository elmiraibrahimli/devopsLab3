package db

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/lib/pq"
)

const (
    host     = "34.201.171.30 "
    port     = 5432 // Default port
    user     = "postgres"
    password = "root"
    dbname   = "devopslab3"
)

var Db *sql.DB

func Connect() {
    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
        "password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)

    var err error
    Db, err = sql.Open("postgres", psqlInfo)
    if err != nil {
        log.Fatal(err)
    }

    err = Db.Ping()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Successfully connected!")
}
















// package db

// import (
//     "database/sql"
//     "log"

//     _ "github.com/lib/pq"

// )

// var db *sql.DB


// func InitDB() {
//     connStr := "postgres://your_username:your_password@localhost/your_database?sslmode=disable"
//     var err error
//     db, err = sql.Open("postgres", connStr)
//     if err != nil {
//         log.Fatal(err)
//     }

//     err = db.Ping()
//     if err != nil {
//         log.Fatal(err)
//     }

//     log.Println("Connected to PostgreSQL database!")
// }
