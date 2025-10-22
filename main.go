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

func main() {
	defer func() {
		fmt.Println("\nExiting program")
		os.Exit(0)
	}()
	loop := true
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(os.Stdin)
	for loop {
		fmt.Printf("\x1bc")
		fmt.Print("--- Welcome to System Authentication Flow ---\n\n")

		if user.UserLogin.Email != "" {
			fmt.Printf("Hello %v!\n", user.UserLogin.Email)
			fmt.Println("1. List All Users")
			fmt.Println("2. Logout")

			fmt.Print("\n0. exit\n\n")

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
			fmt.Println("1. Register")
			fmt.Println("2. Login")
			fmt.Println("3. Forgot Password")

			fmt.Print("\n0. exit\n\n")

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
