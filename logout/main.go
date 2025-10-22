package logout

import (
	"auth-flow/user"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Logout() {
	fmt.Printf("\x1bc")
	scanner := bufio.NewScanner(os.Stdin)
	reader := bufio.NewReader(os.Stdin)
	loop := true

	for loop {
		fmt.Print("Are you sure you want to logout (y/n)? ")
		inputConfirm, _ := reader.ReadString('\n')
		confirm := strings.TrimSpace(inputConfirm)
		if strings.ToLower(confirm) == "y" {
			user.UserLogin = user.User{}
			fmt.Print("Logout success, enter to back to home.. ")
			scanner.Scan()
			loop = false
		} else if strings.ToLower(confirm) != "n" {
			fmt.Print("Invalid input, please try again.. ")
			scanner.Scan()
		}
	}
}
