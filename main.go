package main

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
)

type Article struct {
    Title string `json:"Title"`
    Desc string `json:"desc"`
    Content string `json:"content"`

}

type Articles []Article



func homePage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: HomePage")
}

func handleRequests() {
    myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/", homePage).Queries("surname", "{surname}")
    myRouter.HandleFunc("/all", returnAllArticles).Methods("PUT")
    myRouter.HandleFunc("/article/test{id}test", returnSingleArticle)
    log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r)
    vars := mux.Vars(r)
    key := vars["id"]

    fmt.Fprintf(w, "Key: " + key)
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
    articles := Articles{
        Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
        Article{Title: "Hello2", Desc: "Article Description", Content: "Article Content"},
    }
    fmt.Printf("Endpoint Hit: returnAllArticles")


    json.NewEncoder(w).Encode(articles)
}


func main() {
    fmt.Println("Rest API v2.0 - Mux Routers")
    handleRequests()
}












