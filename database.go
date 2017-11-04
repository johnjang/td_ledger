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


/*
func main() {
    rows, _ := db.Query("select name from items")
    var name string
    for rows.Next() {
        rows.Scan(&name)
        fmt.Println(name)
    }
    insertItem(true , 30,1,1514,"test2", 5.12, 3.0)
}
*/


func insertItem(credit bool, day int, month int, year int, name string, withdraw float64, deposit float64) {
    if db == nil {
        fmt.Println("database error...")
    }
    var cre = 0
    if credit {
        cre = 1
    }

    stmt, err := db.Prepare ("INSERT INTO items (id, day, month, year, name, withdraw, deposit) values (?, ?, ?, ?, ?, ?, ?)")
    if err != nil {
        fmt.Println("shit")
        return
    }
    //res, err := stmt.Exec(cre, day, month, year, name, withdraw, deposit)
    _, err = stmt.Exec(cre, day, month, year, name, withdraw, deposit)
    if err != nil {
        fmt.Println("sshit")
        return
    }
    /*
    id, err := res.LastInsertId()
    if err != nil {
        fmt.Println("sshit")
        return
    }
    */
}

func queryItem(sd, sm, sy, ed, em, ey int) {
    fmt.Println(sd, sm, sy, ed, em, ey)
    //select * from items where day >= start and day <= end and month >= m
    condition1 := "day >= ? and month >= ? and year >=?"
    condition2 := "day <= ? and month <= ? and year <=?"
    query := "SELECT * FROM items WHERE " + condition1 + " and " + condition2
    rows, err := db.Query(query, sd, sm, sy, ed, em, ey)
    if err != nil {
        fmt.Println(err)
    }
    var name string
    var withdraw float64
    var deposit float64

    for rows.Next() {
        err = rows.Scan(&name, &withdraw, &deposit)
        println(name, withdraw, deposit)
    }


    rows.Close()

}




















