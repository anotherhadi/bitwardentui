package main

import (
	"fmt"
	"os"
	"time"

	"github.com/anotherhadi/bitwarden_tui/ui/loading"
	"github.com/anotherhadi/bitwarden_tui/ui/login"
	"github.com/anotherhadi/bitwarden_tui/ui/unlock"
	"github.com/anotherhadi/bitwarden_tui/vault"
)

func exitOnError(err error) {
	if err != nil {
		fmt.Println("Bitwarden TUI exited with error:\n", err)
		os.Exit(1)
	}
}

func main() {
	quit := make(chan struct{})
	go loading.Loading(quit)
	v, err := vault.LoadVault()
	exitOnError(err)
	quit <- struct{}{}
	time.Sleep(100000000)

	// var DEBUG bool = false
	// if DEBUG {
	// 	fmt.Println("DEBUG: ", v)
	//
	// 	os.Exit(0)
	// }

	if v.Status == "unauthenticated" {
		serverUrl, username, password, err := login.Login()
		exitOnError(err)
		err = v.Login(serverUrl, username, password)
		exitOnError(err)
		fmt.Println("Login successful !")
	} else if v.Status == "locked" {
		password, err := unlock.Unlock()
		exitOnError(err)
		err = v.Unlock(password)
		exitOnError(err)
		fmt.Println("Unlock successful !")
	} else {
		os.Exit(0)
	}
}
