package main

import (
	"github.com/bjdean/gonetcheck"
	"os"
	"fmt"
)

// Check if internet access is possible
// If successful exit with 0
// If internet access is detected as down exit with 1
// If some other error occurs print the error and exit with 2
func main () {
	can_access_internet, err := gonetcheck.CheckInternetAccess()
	switch err {
		case nil:
			switch can_access_internet {
				case true: os.Exit(0)
				default: os.Exit(1)
			}
		default:
			fmt.Println(
				"Error returned by gonetcheck.CheckInternetAccess:",
				err)
			os.Exit(2)
	}
}
