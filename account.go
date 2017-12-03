package main

type Account struct {
	Name  string `json:"name"`
	Age   int8   `json:"age"`
	Email string `json:"email"`
}

type Accounts []Account