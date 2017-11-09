package main

import (
    "bytes"
    "strings"
    "fmt"
    "database/sql"
     _ "github.com/mattn/go-sqlite3"
)


//for database connection
var db *sql.DB
func init() {
    var err error
    db, err = sql.Open("sqlite3", "/opt/playground/td_ledger/data/sql/ledger.db")
    handleError(err)
    err = db.Ping()
    handleError(err)
}

/*
func main() {
    fmt.Println(queryUncategorizedItems())
}
*/

func checkDatabase() bool {
    if db == nil {
        fmt.Println("database error...")
        return false
    }
    return true

}

//QEURY [DELETE]
func deleteCategory(item, category string) {

}

//QUERY [POST]
func insertItem(credit bool, date string, name string, withdraw float64, deposit float64) {
    if checkDatabase() {
        var cre = 0
        if credit {
            cre = 1
        }

        stmt, err := db.Prepare ("INSERT INTO items (date, name, withdraw, deposit, id) values (?, ?, ?, ?, ?)")
        handleError(err)
        _, err = stmt.Exec(date, name, withdraw, deposit, cre)
        handleError(err)
    }
}

/*
    NOTE: Assumes item already exists in the items table
*/
func insertCategory(item, category string) {
    if checkDatabase() {
        stmt, err := db.Prepare("INSERT OR REPLACE INTO categories (name, category) values (?, ?)")
        handleError(err)
        _, err = stmt.Exec(item, strings.ToUpper(category))
        handleError(err)
    }
}

//QUERY  [GET]
func queryCategoriesForItem(item string) string {
    query := "SELECT category FROM categories where name = ?"
    rows, err := db.Query(query, item)

    handleError(err)

    res := rowResultCategory(rows, false, true)

    var categories bytes.Buffer
    for _, element := range res {
        categories.WriteString(element[1])
        categories.WriteString(", ")
    }
    categoriesString := categories.String()
    if len(categoriesString) > 2 {
        return categoriesString[:len(categoriesString)-2]
    }
    return "Uncategorized"
}


func queryItemsForCategory(category string) string {
    query := "SELECT name FROM categories where category= ?"
    rows, err := db.Query(query, category)

    handleError(err)

    res := rowResultCategory(rows, true, false)

    var items bytes.Buffer
    for _, element := range res {
        items.WriteString(element[0])
        items.WriteString(", ")
    }
    itemsString := items.String()
    if len(itemsString) > 2 {
        return itemsString[:len(itemsString)-2]
    }
    return "None..."
}

func queryCategory(category string) bool {
    query := "SELECT * FROM categories where category = ? COLLATE NOCASE"
    rows, err := db.Query(query, category)

    if handleError(err) {
        return false
    }

    res := rowResultCategory(rows, true, true)
    if len(res) == 0 {
        return false
    }

    return true
}

func queryUncategorizedItems() []string{
    var err error
    var res = []string{}

    query := "select distinct t1.name from items t1 left join categories t2 on t2.name = t1.name where t2.name is null"
    rows, err := db.Query(query)

    for rows.Next() {
        var name string
        err = rows.Scan(&name)
        if !handleError(err) {
            res = append(res, name)
        }
    }

    rows.Close()
    return res

}

func queryItem(item string) bool {
    query := "SELECT * FROM items where name = ? COLLATE NOCASE"
    rows, err := db.Query(query, item)

    if handleError(err) {
        return false
    }

    res := rowResult(rows)
    if len(res) == 0 {
        return false
    }

    return true
}

func queryItemDateStart(start string) [][5]string {
    query := "SELECT * FROM items where date >= ? ORDER BY date DESC"
    rows, err := db.Query(query, start)
    handleError(err)

    return rowResult(rows)
}


func queryItemDateEnd(end string) [][5]string {
    query := "SELECT * FROM items where date <= ? ORDER BY date DESC"
    rows, err := db.Query(query, end)
    handleError(err)

    return rowResult(rows)
}

func queryItemDateRange(start, end string) [][5]string {
    query := "SELECT * FROM items where date >= ? and date <= ? ORDER BY date DESC"
    rows, err := db.Query(query, start, end)
    handleError(err)

    return rowResult(rows)
}

func queryItemName(name string) [][5]string {
    query := "SELECT * FROM items where name = ? ORDER BY date DESC"
    rows, err := db.Query(query, name)
    handleError(err)

    return rowResult(rows)
}

func queryItemNameDateRange(name, start, end string) [][5]string {
    query := "SELECT * FROM items where name = ? and date >= ? and date <= ? ORDER BY date DESC"
    rows, err := db.Query(query, name, start, end)
    handleError(err)
    return rowResult(rows)

}

func handleError(err error) bool{
    if err != nil {
        fmt.Println(err)
        return true
    }
    return false
}

func rowResultCategory(rows *sql.Rows, name, category bool) [][2]string {
    var err error
    var res = [][2]string{}

    if name && category {
        for rows.Next() {
            var name string
            var category string
            err = rows.Scan(&name, &category)
            if !handleError(err) {
                res = append(res, [2]string{name, category})
            }
        }
    }

    if category {
        for rows.Next() {
            var category string
            err = rows.Scan(&category)
            if !handleError(err) {
                res = append(res, [2]string{"", category})
            }
        }
    }

    if name {
        for rows.Next() {
            var name string
            err = rows.Scan(&name)
            if !handleError(err) {
                res = append(res, [2]string{name, ""})
            }
        }
    }


    rows.Close()
    return res
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
        if !handleError(err) {
            res = append(res, [5]string{date, name, withdraw, deposit, id})
        }

    }
    rows.Close()
    return res
}

