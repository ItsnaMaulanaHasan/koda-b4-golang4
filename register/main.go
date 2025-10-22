package register

import (
	"auth-flow/user"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Register() {
	loop := true
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(os.Stdin)

	dataRegister := user.User{}
	for loop {
		fmt.Printf("\x1bc")
		fmt.Print("--- Register ---\n\n")

		fmt.Print("What is your first name: ")
		inputFirstName, _ := reader.ReadString('\n')
		firstName := strings.TrimSpace(inputFirstName)
		dataRegister.FirstName = firstName

		fmt.Print("What is your last name: ")
		inputLastName, _ := reader.ReadString('\n')
		lastName := strings.TrimSpace(inputLastName)
		dataRegister.LastName = lastName

		fmt.Print("What is your email: ")
		inputEmail, _ := reader.ReadString('\n')
		email := strings.TrimSpace(inputEmail)
		dataRegister.Email = email

		fmt.Print("Input password: ")
		inputPassword, _ := reader.ReadString('\n')
		password := strings.TrimSpace(inputPassword)

		fmt.Print("Confirm password: ")
		inputConfirmPassword, _ := reader.ReadString('\n')
		confirmPassword := strings.TrimSpace(inputConfirmPassword)

		if password != confirmPassword {
			fmt.Print("Incorrect password confirmation, please try again ")
			scanner.Scan()
			continue
		}
		dataRegister.Password = password

		fmt.Println("Is the data correct?")
		fmt.Printf("First Name: %v\n", dataRegister.FirstName)
		fmt.Printf("Last Name: %v\n", dataRegister.LastName)
		fmt.Printf("Email: %v\n", dataRegister.Email)
		fmt.Print("Continue (y/n): ")
		inputConfirm, _ := reader.ReadString('\n')
		confirm := strings.TrimSpace(inputConfirm)

		if strings.ToLower(confirm) == "y" {
			user.Users = append(user.Users, dataRegister)
			fmt.Print("Register success, enter to back.. ")
			scanner.Scan()
			loop = false
		}
	}
}
