package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type tasks struct {
	ID      int    `json:ID`
	Name    string `json:Name`
	Content string `json:Content`
}

type allTask []tasks

var task = allTask{
	{
		ID:      1,
		Name:    "Task one",
		Content: "Some Content",
	},
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my API 2022")

}
func main() {

	router := mux.NewRouter().StrictSlash(true) //strinct mode (/route/<--- incomplet)

	router.HandleFunc("/", indexRoute)
	log.Fatal(http.ListenAndServe(":5000", router))

}
