package user

import (
	"bufio"
	"fmt"
	"os"
)

type User struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}

func (fn User) GetFullName() string {
	return fmt.Sprintf("%v %v", fn.FirstName, fn.LastName)
}

var Users []User

var UserLogin User

func ShowListUser() {
	loop := true
	scanner := bufio.NewScanner(os.Stdin)

	for loop {
		fmt.Printf("\x1bc")
		fmt.Print("--- List All User ---\n\n")
		for i, item := range Users {
			fmt.Printf("%v.\n", i+1)
			fmt.Printf("Full Name: %v\n", item.GetFullName())
			fmt.Printf("Password: %v\n", item.Password)
		}

		fmt.Print("\n\n Enter to back to home.. ")
		scanner.Scan()
		loop = false
	}
}
