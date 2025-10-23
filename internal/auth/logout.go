package auth

import (
	"auth-flow/internal/user"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Logout() {
	scanner := bufio.NewScanner(os.Stdin)
	reader := bufio.NewReader(os.Stdin)
	loop := true

	for loop {
		func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Printf("\nError: %v\n", r)
					fmt.Print("Press enter to try again...")
					scanner.Scan()
				}
			}()

			fmt.Printf("\x1bc")
			fmt.Print("Are you sure you want to logout (y/n)? ")
			inputConfirm, _ := reader.ReadString('\n')
			confirm := strings.TrimSpace(inputConfirm)

			if strings.ToLower(confirm) == "y" {
				currentUser := &user.UserLogin
				defer func() {
					fmt.Printf("User %v logged out successfully\n", currentUser.Email)
				}()

				user.UserLogin = user.User{}

				fmt.Print("Logout success, enter to back to home.. ")
				scanner.Scan()
				loop = false
			} else if strings.ToLower(confirm) == "n" {
				fmt.Print("Logout cancelled, returning to menu.. ")
				scanner.Scan()
				loop = false
			} else {
				panic("Invalid input! Please enter 'y' or 'n'")
			}
		}()
	}
}
