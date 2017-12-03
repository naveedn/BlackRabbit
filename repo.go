package main

import (
	"fmt"
)

var currentId int
var accounts Accounts

func init() {
	RepoCreate(Account{Name: "Robbie Wagner", Age: 26, Email: "robbie.wagner@gmail.com"})
	RepoCreate(Account{Name: "Naveed Nadjmabadi", Age: 25, Email: "naveed.nadjmabadi@gmail.com"})
}

func RepoCreate(a Account) Account {
	currentId += 1
	a.Id = currentId
	accounts = append(accounts, a)
	return a
}

func RepoFind(id int) Account {
	for _, a := range accounts {
		if a.Id == id {
			return a
		}
	}

	return Account{}
}

func RepoDestroy(id int) error {
	for i, a := range accounts {
		if a.Id == id {
			accounts = append(accounts[:i], accounts[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("Could not find account with Id %d", id)
}
