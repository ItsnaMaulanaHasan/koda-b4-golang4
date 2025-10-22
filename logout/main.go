package logout

import (
	"auth-flow/user"
	"bufio"
	"fmt"
	"os"
)

func Logout() {
	fmt.Printf("\x1bc")
	scanner := bufio.NewScanner(os.Stdin)
	user.UserLogin = user.User{}
	fmt.Print("Logout succes, enter to back to home.. ")
	scanner.Scan()
}
