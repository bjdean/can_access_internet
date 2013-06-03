package main

import (
	"fmt"
	"github.com/bjdean/gonetcheck"
	"os"
)

// The default test URLS
var testUrls = []string{
	"http://www.google.com/",
	"http://www.bing.com/",
	"http://www.microsoft.com/",
	"http://www.apple.com/",
	"http://yahoo.com/",
	"http://www.unimelb.edu.au/",
	"http://abc.net.au/",
	"http://www.monash.edu.au/",
	"http://bbb.co.uk/",
}

// The default test TCP addresses
var testTcpAddrs = []string{
	"www.google.com:443",
	"www.example.com:80",
	"tty.freeshell.net:22",
}

// Check if internet access is possible
// If successful exit with 0
// If internet access is detected as down exit with 1
// If some other error occurs print the error and exit with 2
func main() {
	canAccessInternet, errList := gonetcheck.CheckInternetAccess(
		testUrls,
		testTcpAddrs)
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
