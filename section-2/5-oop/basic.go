package main

import "fmt"

type Account struct {
	name   string
	Username string
}

func (a *Account) SetName(name string) {
	a.name = name
}

func (a *Account) GetName() string {
	return a.name
}

func main() {
	account := &Account{Username: "cahyonadi"}
	account.SetName("Antonius Cahyo")

	fmt.Println("Username:", account.Username)
	fmt.Println("Name:", account.GetName())
	fmt.Println(account.name)
}