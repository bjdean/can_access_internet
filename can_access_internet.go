package main

import (
	"fmt"
	"github.com/bjdean/gonetcheck"
	"os"
)

// Check if internet access is possible
// If successful exit with 0
// If internet access is detected as down exit with 1
// If some other error occurs print the error and exit with 2
func main() {
	canAccessInternet, errList := gonetcheck.CheckInternetAccess()
	switch errList {
	case nil:
		switch canAccessInternet {
		case true:
			os.Exit(0)
		default:
			os.Exit(1)
		}
	default:
		fmt.Println(
			"Error(s) returned by gonetcheck.CheckInternetAccess:",
			errList)
		os.Exit(2)
	}
}
