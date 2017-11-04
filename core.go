package main

import (
    "fmt"
    "time"
)

type DateParms struct {
    day, month, year int
}
type RangeParms struct {
    start, end DateParms
}

//select * from tables where day > x and day < y and month.... year....


//get range from start tne end inclusively
/*
    if end is is not given, then it will be start to current date
    

*/

func getDataDateRange(dateRange RangeParms) {
    empty := DateParms{}

    if dateRange.end == empty {
        current_date := parseDate(time.Now().Local().Format("01/02/2006"))
        dateRange.end = DateParms{current_date[0], current_date[1], current_date[2]}
    }

    fmt.Println(dateRange)
    queryItem(dateRange.start.day, dateRange.start.month, dateRange.start.year, 
                dateRange.end.day, dateRange.end.month, dateRange.end.year)
//func queryItem(sd, sm, sy, ed, em, ey int) {

}


func main() {

    //fmt.Println("current time is ", current_time.Format("01/02/2006"))  //what the heck golang? placeholder???

    start := DateParms{1, 20, 2000}
    end := DateParms{5, 15, 2019}

    test1 := RangeParms{start, DateParms{}}
    test2 := RangeParms{start, end}
    getDataDateRange(test1)
    fmt.Println("---")
    getDataDateRange(test2)

}

//func parseDate(date string) [3]int {


