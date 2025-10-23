package main

import (
	"auth-flow/internal/auth"
	"auth-flow/internal/user"
	"bufio"
	"fmt"
	"os"
	"strings"
)

type menuLoggedIn struct {
	emailLoggedIn string
	menu          []string
}

type menuNotLoggedIn struct {
	menu []string
}

func (o menuLoggedIn) outputMenu() {
	fmt.Printf("Hello %v!\n", o.emailLoggedIn)
	for i, item := range o.menu {
		fmt.Printf("%v. %v\n", i+1, item)
	}
	fmt.Print("\n0. exit\n\n")
}

func (o menuLoggedIn) chooseMenu(menu *string, loop *bool, scanner *bufio.Scanner) {
	switch *menu {
	case "1":
		user.ShowListUser()
	case "2":
		auth.Logout()
	case "0":
		*loop = false
	default:
		fmt.Print("Invalid menu option, press enter to continue...")
		scanner.Scan()
	}
}

func (o menuNotLoggedIn) outputMenu() {
	for i, item := range o.menu {
		fmt.Printf("%v. %v\n", i+1, item)
	}
	fmt.Print("\n0. exit\n\n")
}

func (o menuNotLoggedIn) chooseMenu(menu *string, loop *bool, scanner *bufio.Scanner) {
	switch *menu {
	case "1":
		auth.Register()
	case "2":
		auth.Login()
	case "3":
		auth.ForgotPassword()
	case "0":
		*loop = false
	default:
		fmt.Print("Invalid menu option, press enter to continue...")
		scanner.Scan()
	}
}

type outListMenu interface {
	outputMenu()
	chooseMenu(menu *string, loop *bool, scanner *bufio.Scanner)
}

func main() {
	defer func() {
		fmt.Println("\nExiting program")
		os.Exit(0)
	}()
	loop := true
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(os.Stdin)

	var listMenu outListMenu
	for loop {
		fmt.Printf("\x1bc")
		fmt.Print("--- Welcome to System Authentication Flow ---\n\n")

		if user.UserLogin.Email != "" {
			listMenu = menuLoggedIn{user.UserLogin.Email, []string{"List All Users", "Logout"}}
			listMenu.outputMenu()

			fmt.Print("Choose a menu: ")
			input, _ := reader.ReadString('\n')
			menu := strings.TrimSpace(input)

			listMenu.chooseMenu(&menu, &loop, scanner)
		} else {
			listMenu = menuNotLoggedIn{[]string{"Register", "Login", "Forgot Password"}}
			listMenu.outputMenu()

			fmt.Print("Choose a menu: ")
			input, _ := reader.ReadString('\n')
			menu := strings.TrimSpace(input)

			listMenu.chooseMenu(&menu, &loop, scanner)
		}
	}
}

// go build -o auth-flow .
// file auth-flow
// GOOS=windows GOARCH=amd64 go build -o auth-flow.exe .
