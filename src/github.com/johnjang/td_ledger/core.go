package main

import (
    "fmt"
    "bytes"
    "strconv"
)

/*
func main() {
    current_time := time.Now()
    fmt.Println("current time is ", current_time.Format("01/02/2006"))  //what the heck golang? placeholder???
    start := "10-10-2017"
    end := "10-25-2017"
    dateRangeTable(queryItemDateRange(start, end))
}
*/

func addValue(a, b float64, err error) float64{
    if err == nil {
        return a+b
    } else {
        fmt.Println(err)
    }
    return a
}

func inputDateRange(start, end string) string {
    return dateRangeTable(queryItemDateRange(start, end))
}

func inputDateStart(start string) string {
    return dateRangeTable(queryItemDateStart(start))
}

func dateRangeTable(result [][5]string) string {
    var table, header, body bytes.Buffer

    if len(result) == 0 {
        return "No data..."
    }

    //Calculate the body 
    totalWithdraw, totalDeposit := 0.0, 0.0
    for i:=0; i<len(result)-1; i++ {
        if result[i][3] == "0" { //withdraw flag
            body.WriteString(fmt.Sprintf("%d) %s\t%s\t\t - $%s\n", i, result[i][0], result[i][1], result[i][2]))
            withdraw, err := strconv.ParseFloat(result[i][2], 64)
            totalWithdraw = addValue(totalWithdraw, withdraw, err)
        } else {
            body.WriteString(fmt.Sprintf("%d) %s\t%s\t\t + $%s\n", i, result[i][0], result[i][1], result[i][3]))
            deposit, err := strconv.ParseFloat(result[i][3], 64)
            totalDeposit = addValue(totalDeposit, deposit, err)
        }
    }

    header.WriteString(fmt.Sprintf("Date: %s ~ %s\n", result[len(result)-1][0], result[0][0]))
    header.WriteString(fmt.Sprintf("Total Withdraw: %.2f\tTotal Deposit: %.2f\n", totalWithdraw, totalDeposit))
    header.WriteString("=====================================\n")

    table.WriteString(header.String())
    table.WriteString(body.String())
    table.WriteString("\n\n")
    table.WriteString(categoryTable(result))

    return table.String()
}



func categoryTable(result [][5]string) string {
    cmap := make(map[string]float64)
    var table, header, body bytes.Buffer

    for i:=0; i<len(result); i++ {
        category := queryCategoriesForItem(result[i][1])
        _, ok := cmap[category]

        amount := result[i][3]
        if amount == "0" {
            amount = result[i][2]
        }
        amount_float, err := strconv.ParseFloat(amount, 64)
        if err != nil {
            amount_float = 0.0
        }

        if !ok {
            cmap[category] = amount_float
        } else {
            cmap[category] += amount_float
        }
    }

    for k, v := range cmap {
        body.WriteString(fmt.Sprintf("Category: %s\t\tTotal: %.2f\n", k, v))
    }
    header.WriteString("Category spending table...\n")
    header.WriteString("===============================\n")

    table.WriteString(header.String())
    table.WriteString(body.String())

    return table.String()

}



