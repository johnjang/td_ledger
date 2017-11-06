package main

import (
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



func main() {
    fmt.Println("MUX started...")
    router := mux.NewRouter()
    router.HandleFunc("/ledger/date/{start}", dateStart).Methods("GET")
    router.HandleFunc("/ledger/date/{start}/{end}", dateStartEnd).Methods("GET")
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
