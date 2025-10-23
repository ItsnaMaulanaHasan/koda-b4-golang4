package main

import (
	"auth-flow/forgotpassword"
	"auth-flow/login"
	"auth-flow/logout"
	"auth-flow/register"
	"auth-flow/user"
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

func (o menuNotLoggedIn) outputMenu() {
	for i, item := range o.menu {
		fmt.Printf("%v. %v\n", i+1, item)
	}
	fmt.Print("\n0. exit\n\n")
}

type outListMenu interface {
	outputMenu()
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

			switch menu {
			case "1":
				user.ShowListUser()
			case "2":
				logout.Logout()
			case "0":
				loop = false
			default:
				fmt.Print("Invalid menu option, press enter to continue...")
				scanner.Scan()
			}
		} else {
			listMenu = menuNotLoggedIn{[]string{"Register", "Login", "Forgot Password"}}
			listMenu.outputMenu()

			fmt.Print("Choose a menu: ")
			input, _ := reader.ReadString('\n')
			menu := strings.TrimSpace(input)

			switch menu {
			case "1":
				register.Register()
			case "2":
				login.Login()
			case "3":
				forgotpassword.ForgotPassword()
			case "0":
				loop = false
			default:
				fmt.Print("Invalid menu option, press enter to continue...")
				scanner.Scan()
			}
		}
	}
}
