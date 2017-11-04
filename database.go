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
    handleError(err)
    err = db.Ping()
    handleError(err)
}


func main() {
    date1 := "10-25-2017"
    date2 := "10-26-2017"
    name := "FALAFEL KING"
    fmt.Println(queryItemDateRange(date1, date2))
    fmt.Println(queryItemName(name))
    fmt.Println(queryItemNameDateRange(name, date1, date2))
}


func insertItem(credit bool, date string, name string, withdraw float64, deposit float64) {
    if db == nil {
        fmt.Println("database error...")
    }
    var cre = 0
    if credit {
        cre = 1
    }

    stmt, err := db.Prepare ("INSERT INTO items (date, name, withdraw, deposit, id) values (?, ?, ?, ?, ?)")
    handleError(err)
    _, err = stmt.Exec(date, name, withdraw, deposit, cre)
    handleError(err)
}


//QUERY 
func queryItemDateStart(start string) [][5]string {
    query := "SELECT * FROM items where date >= ?"
    rows, err := db.Query(query, start)
    handleError(err)

    return rowResult(rows)
}


func queryItemDateEnd(end string) [][5]string {
    query := "SELECT * FROM items where date <= ?"
    rows, err := db.Query(query, end)
    handleError(err)

    return rowResult(rows)
}

func queryItemDateRange(start, end string) [][5]string {
    query := "SELECT * FROM items where date >= ? and date <= ?"
    rows, err := db.Query(query, start, end)
    handleError(err)

    return rowResult(rows)
}

func queryItemName(name string) [][5]string {
    query := "SELECT * FROM items where name = ?"
    rows, err := db.Query(query, name)
    handleError(err)

    return rowResult(rows)
}

func queryItemNameDateRange(name, start, end string) [][5]string {
    query := "SELECT * FROM items where name = ? and date >= ? and date <= ?"
    rows, err := db.Query(query, name, start, end)
    handleError(err)
    return rowResult(rows)

}

func handleError(err error) {
    if err != nil {
        fmt.Println(err)
    }
}

func rowResult(rows *sql.Rows) [][5]string {
    var err error
    var res = [][5]string{}

    for rows.Next() {
        var date string
        var name string
        var withdraw string
        var deposit string
        var id string
        err = rows.Scan(&date, &name, &withdraw, &deposit, &id)
        handleError(err)

        res = append(res, [5]string{date, name, withdraw, deposit, id})
    }
    rows.Close()
    return res
}

