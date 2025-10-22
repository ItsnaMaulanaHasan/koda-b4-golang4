package login

import (
	"auth-flow/user"
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

func Login() {
	loop := true
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(os.Stdin)

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
			fmt.Print("--- Login ---\n\n")

			fmt.Print("Enter your email: ")
			inputEmail, _ := reader.ReadString('\n')
			email := strings.TrimSpace(inputEmail)

			if email == "" {
				panic("Email cannot be empty!")
			}

			if !strings.Contains(email, "@") || !strings.Contains(email, ".") {
				panic("Invalid email format!")
			}

			fmt.Print("Enter your password: ")
			inputPassword, _ := reader.ReadString('\n')
			password := strings.TrimSpace(inputPassword)

			if password == "" {
				panic("Password cannot be empty!")
			}

			var hashedPassword string
			func() {
				defer func() {
					password = ""
				}()
				hash := md5.Sum([]byte(password))
				hashedPassword = hex.EncodeToString(hash[:])
			}()

			validLogin := false
			var foundUser *user.User

			for i := range user.Users {
				if user.Users[i].Email == email {
					if user.Users[i].Password == hashedPassword {
						validLogin = true
						foundUser = &user.Users[i]
						break
					}
				}
			}

			if !validLogin {
				panic("Invalid email or password!")
			}

			if foundUser != nil {
				user.UserLogin = *foundUser
			}

			fmt.Print("Login success, enter to back to home.. ")
			scanner.Scan()
			loop = false
		}()
	}
}
