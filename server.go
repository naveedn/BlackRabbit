package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Account struct {
	Name  string
	Age   int8
	Email string
}

type Accounts []Account

func handler(w http.ResponseWriter, r *http.Request) {
	accounts := Accounts{
		Account{Name: "Robbie Wagner", Age: 26, Email: "robbie.wagner@gmail.com"},
		Account{Name: "Naveed Nadjmabadi", Age: 25, Email: "naveed.nadjmabadi@gmail.com"},
	}

	json.NewEncoder(w).Encode(accounts)
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/Accounts", handler)
	log.Fatal(http.ListenAndServe(":8080", router))
}
