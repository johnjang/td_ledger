package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
    "strconv"
)


//Given a string in mm/dd/yyy format, it will return a list of integers
// containing above values in the order
func parseDate(date string) [3]int {
    str_dates := strings.Split(date, "/")
    var int_dates [3]int

    for i:=0; i<len(str_dates); i++ {
        int_dates[i], _ = strconv.Atoi(str_dates[i])
    }

    return int_dates

}


//returns sqllike query hopefully..
func parseLine(line string) {
    //MM/DD/YYYY, NAME,withrdaw,Deposit,total
    s_line := strings.Split(line, ",")
    if len(s_line) != 5 {
        fmt.Println("ERROR")
        //throw an exception?
    }

    int_dates := parseDate(s_line[0])

    name := s_line[1]
    var withdraw float64
    var deposit float64

    if s_line[2] != "" {
        val, err := strconv.ParseFloat(s_line[2], 64)
        if err != nil {
            fmt.Println(err)
        }
        withdraw = val
    }
    if s_line[3] != "" {
        val, err := strconv.ParseFloat(s_line[3], 64)
        if err != nil {
            fmt.Println(err)
        }
        deposit = val;
    }

    fmt.Println(int_dates, name, withdraw, deposit)
    //insertItem(false, int_dates[0], int_dates[1], int_dates[2], name, withdraw, deposit)



}

func readFile(file_location string) {
    file, err := os.Open(file_location)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        parseLine(scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}


/*

func main() {
    stat, err := os.Stat(os.Args[1])
    if err != nil {
        fmt.Println(err)
        return
    }
    switch mode := stat.Mode(); {
    case mode.IsRegular():
        //do nothing. 
    default:
        return
    }

    readFile(os.Args[1])
}

*/


