package main

import (
	"fmt"

	"github.com/anotherhadi/bitwarden_tui/vault"
)

func main() {
	v, err := vault.LoadVault()
	if err != nil {
		panic(err)
	}
	fmt.Println(v)

	fmt.Println("Logout:")
	err = v.Logout()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Login:")
	err = v.Login("", "test@proton.me", "test")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Token:", v.SessionKey)
}
