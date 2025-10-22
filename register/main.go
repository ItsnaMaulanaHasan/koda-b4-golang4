package register

import (
	"auth-flow/user"
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

func Register() {
	defer func() {
		fmt.Println("Exiting registration process...")
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

			dataRegister := &user.User{}

			fmt.Printf("\x1bc")
			fmt.Print("--- Register ---\n\n")

			fmt.Print("What is your first name: ")
			inputFirstName, _ := reader.ReadString('\n')
			firstName := strings.TrimSpace(inputFirstName)

			if firstName == "" {
				panic("First name cannot be empty!")
			}
			dataRegister.FirstName = firstName

			fmt.Print("What is your last name: ")
			inputLastName, _ := reader.ReadString('\n')
			lastName := strings.TrimSpace(inputLastName)

			if lastName == "" {
				panic("Last name cannot be empty!")
			}
			dataRegister.LastName = lastName

			fmt.Print("What is your email: ")
			inputEmail, _ := reader.ReadString('\n')
			email := strings.TrimSpace(inputEmail)

			if !strings.Contains(email, "@") || !strings.Contains(email, ".") {
				panic("Invalid email format!")
			}

			for i := range user.Users {
				if user.Users[i].Email == email {
					panic("Email already registered!")
				}
			}
			dataRegister.Email = email

			fmt.Print("Input password: ")
			inputPassword, _ := reader.ReadString('\n')
			password := strings.TrimSpace(inputPassword)

			fmt.Print("Confirm password: ")
			inputConfirmPassword, _ := reader.ReadString('\n')
			confirmPassword := strings.TrimSpace(inputConfirmPassword)

			if password != confirmPassword {
				panic("Password confirmation does not match!")
			}

			func() {
				defer func() {
					password = ""
					confirmPassword = ""
				}()

				hash := md5.Sum([]byte(password))
				hashedPassword := hex.EncodeToString(hash[:])
				dataRegister.Password = hashedPassword
			}()

			fmt.Println("\n\nIs the data correct?")
			fmt.Printf("First Name: %v\n", dataRegister.FirstName)
			fmt.Printf("Last Name: %v\n", dataRegister.LastName)
			fmt.Printf("Email: %v\n\n", dataRegister.Email)
			fmt.Print("Continue (y/n): ")
			inputConfirm, _ := reader.ReadString('\n')
			confirm := strings.TrimSpace(inputConfirm)

			if strings.ToLower(confirm) == "y" {
				fmt.Printf("\x1bc")
				user.Users = append(user.Users, *dataRegister)
				fmt.Print("Register success, enter to back to home.. ")
				scanner.Scan()
				loop = false
			} else if strings.ToLower(confirm) == "n" {
				fmt.Print("Registration cancelled, press enter to restart...")
				scanner.Scan()
			} else {
				panic("Invalid input! Please enter 'y' or 'n'")
			}
		}()
	}
}
