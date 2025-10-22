package forgotpassword

import (
	"auth-flow/user"
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

func ForgotPassword() {
	// Defer untuk cleanup
	defer func() {
		fmt.Println("Exiting forgot password process...")
	}()

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
			fmt.Print("--- Forgot Password ---\n\n")

			fmt.Print("Enter your email: ")
			inputEmail, _ := reader.ReadString('\n')
			email := strings.TrimSpace(inputEmail)

			if email == "" {
				panic("Email cannot be empty!")
			}

			if !strings.Contains(email, "@") || !strings.Contains(email, ".") {
				panic("Invalid email format!")
			}

			fmt.Print("Enter new password: ")
			inputPassword, _ := reader.ReadString('\n')
			password := strings.TrimSpace(inputPassword)

			if password == "" {
				panic("Password cannot be empty!")
			}

			fmt.Print("Confirm password: ")
			inputConfirmPassword, _ := reader.ReadString('\n')
			confirmPassword := strings.TrimSpace(inputConfirmPassword)

			if password != confirmPassword {
				panic("Password confirmation does not match!")
			}

			var hashedPassword string
			func() {
				defer func() {
					password = ""
					confirmPassword = ""
				}()
				hash := md5.Sum([]byte(password))
				hashedPassword = hex.EncodeToString(hash[:])
			}()

			userFound := false
			var foundUser *user.User

			for i := range user.Users {
				if user.Users[i].Email == email {
					foundUser = &user.Users[i]
					userFound = true
					break
				}
			}

			if !userFound {
				panic("Email not registered!")
			}

			foundUser.Password = hashedPassword

			fmt.Print("Password updated successfully, enter to back to home.. ")
			scanner.Scan()
			loop = false
		}()
	}
}
