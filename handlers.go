package main

import (
	"encoding/json"
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
)

func AccountIndex(w http.ResponseWriter, r *http.Request) {
	accounts := Accounts{
		Account{Name: "Robbie Wagner", Age: 26, Email: "robbie.wagner@gmail.com"},
		Account{Name: "Naveed Nadjmabadi", Age: 25, Email: "naveed.nadjmabadi@gmail.com"},
	}

	if err := json.NewEncoder(w).Encode(accounts); err != nil {
		panic(err)
	}
}

func AccountId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountId := vars["accountID"]
	fmt.Fprintln(w, "Account Show:", accountId)
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}
