package login

import (
	"auth-flow/user"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Login() {
	loop := true
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(os.Stdin)
	validLogin := false

	for loop {
		fmt.Printf("\x1bc")
		fmt.Print("--- Login ---\n\n")

		fmt.Print("Enter your email: ")
		inputEmail, _ := reader.ReadString('\n')
		email := strings.TrimSpace(inputEmail)

		fmt.Print("Enter your password: ")
		inputPassword, _ := reader.ReadString('\n')
		password := strings.TrimSpace(inputPassword)

		for _, item := range user.Users {
			if item.Email == email {
				if item.Password == password {
					validLogin = true
					user.UserLogin = item
				}
			}
		}

		if !validLogin {
			fmt.Print("Invalid email or password, try again.. ")
			scanner.Scan()
			continue
		}

		fmt.Print("Login success, enter to back to home.. ")
		scanner.Scan()
		loop = false
	}
}
