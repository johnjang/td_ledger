package main

import (
    "fmt"
    "database/sql"
     _ "github.com/mattn/go-sqlite3"
)




//for database connection
var db *sql.DB
func init() {
    var err error
    db, err = sql.Open("sqlite3", "/opt/playground/ledger/data/sql/ledger.db")
    if err != nil {
        //log error and...?
        fmt.Println(err)
    }
    if err = db.Ping(); err != nil {
        fmt.Println(err)
    }
}


func main() {

    rows, _ := db.Query("select name from items")

    var name string
    for rows.Next() {
        rows.Scan(&name)
        fmt.Println(name)

    }

}








