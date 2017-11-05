package main

import (
    "fmt"
    "time"
)


func main() {

    current_time := time.Now()
    fmt.Println("current time is ", current_time.Format("01/02/2006"))  //what the heck golang? placeholder???

}



