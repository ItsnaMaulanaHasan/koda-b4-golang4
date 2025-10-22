package main

import (
	"auth-flow/register"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	loop := true
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(os.Stdin)
	for loop {
		fmt.Printf("\x1bc")
		fmt.Print("--- Welcome to System Authentication Flow ---\n\n")

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
			fmt.Print("2")
		case "3":
			fmt.Print("3")
		case "0":
			loop = false
		default:
			fmt.Print("error")
			scanner.Scan()
		}
	}
}
