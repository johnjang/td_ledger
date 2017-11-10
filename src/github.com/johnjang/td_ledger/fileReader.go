package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
    "strconv"
    "regexp"
)


//Given a string in mm/dd/yyy format, convert into mm-dd-yyy
func parseDateSimple(date string) string {
    return strings.Replace(date, "/", "-", -1)
}


//returns sqllike query hopefully..
func parseLine(line string) {
    //MM/DD/YYYY, NAME,withrdaw,Deposit,total
    s_line := strings.Split(line, ",")
    if len(s_line) != 5 {
        fmt.Println("ERROR")
        //throw an exception?
    }

    dates := parseDateSimple(s_line[0])

    name := s_line[1]
    re := regexp.MustCompile(".*(#.*)")
    name = re.ReplaceAllString(name, "")

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

    insertItem(false, dates, name, withdraw, deposit)

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



