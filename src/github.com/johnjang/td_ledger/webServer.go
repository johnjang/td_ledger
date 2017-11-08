package main

import (
    "bytes"
    "log"
    "net/http"
    "fmt"
    "github.com/gorilla/mux"
    "strings"
)

/*
change YYYY-MM-DD or
       DD-MM-YYYY or
       MM-DD it will assume to be a day and month of current year
       DD,  it will assume to be a day with current monty and year or
        into 
       MM-DD-YYYY format
*/
func parseDate(date string) string {
    /*
        it is DD-MM-YYYY or YYYY-MM-DD
    */
    if len(date)==10 {
        date := strings.Split(date, "-")
        if len(date[0]) == 4 {  //YYYY-MM-DD format
            return date[1] + "-" + date[2] + "-" + date[0]
        } else { //DD-MM-YYYY format
            return date[1] + "-" + date[0] + "-" + date[2]
        }
    }
    return "none"
}

func dateStart(w http.ResponseWriter, req *http.Request) {
    start := mux.Vars(req)["start"]
    start = parseDate(start)
    result := inputDateStart(start)

    w.Write([]byte(result))
}

func dateStartEnd(w http.ResponseWriter, req *http.Request) {
    start := mux.Vars(req)["start"]
    start = parseDate(start)

    end := mux.Vars(req)["end"]
    end = parseDate(end)

    result := inputDateRange(start, end)

    w.Write([]byte(result))
}

func addCategory(w http.ResponseWriter, req *http.Request) {
    //add caregory
    item := mux.Vars(req)["item"]
    if !queryItem(item) {
        w.Write([]byte("Given item not found..."))
        return
    }
    insertCategory(item, mux.Vars(req)["category"])

    w.Write([]byte("done"))
}

func listCategory(w http.ResponseWriter, req *http.Request) {
    name := mux.Vars(req)["name"]

    var res bytes.Buffer
    if !queryItem(name) {
        res.WriteString("no such item...")
    } else {
        res.WriteString(queryCategoriesForItem(name))
    }
    res.WriteString("\n\n")
    if !queryCategory(name) {
        res.WriteString("no such category...")
    } else {
        res.WriteString(queryItemsForCategory(name))
    }
    w.Write([]byte(res.String()))

}


func main() {
    fmt.Println("MUX started...")
    router := mux.NewRouter()
    router.HandleFunc("/ledger/date/{start}", dateStart).Methods("GET")
    router.HandleFunc("/ledger/date/{start}/{end}", dateStartEnd).Methods("GET")
    router.HandleFunc("/ledger/category/add/{item}/{category}", addCategory).Methods("PUT")
    router.HandleFunc("/ledger/category/list/{name}", listCategory).Methods("GET")

    log.Fatal(http.ListenAndServe(":8080", router))
}


/*

//curl -X GET "http://127.0.0.1:8080/hellotestt?arg1=this"
    end := req.FormValue("end")
    if end != "" {
        end = parseDate(end)
    }
    fmt.Println(end)
    w.Write([]byte("hello\nworld"))

*/
