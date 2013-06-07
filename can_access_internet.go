package main

import (
	"fmt"
	"github.com/bjdean/gonetcheck"
	"os"
	"flag"
	"encoding/json"
	"time"
)

// Command-line flags
var progCfgFilepath string
var verbose bool
var timeoutSeconds int
func init() {
	flag.StringVar(&progCfgFilepath, "cfg", "", "configuration override")
	flag.StringVar(&progCfgFilepath, "c", "", "configuration override")

	flag.BoolVar(&verbose, "verbose", false, "verbose output")
	flag.BoolVar(&verbose, "v", false, "verbose output")

	flag.IntVar(&timeoutSeconds, "timeout", 10, "timeout (in seconds)")
	flag.IntVar(&timeoutSeconds, "t", 10, "timeout (in seconds)")
}

// Check if internet access is possible
// If successful exit with 0
// If internet access is detected as down exit with 1
// If some other error occurs print the error and exit with 2
func main() {
	flag.Parse()

	if verbose {
		gonetcheck.DEBUG = gonetcheck.DBG_VERBOSE
	}

	var cfg progCfg
	switch progCfgFilepath {
		case "":
			cfg = defaultProcCfg
		default:
			cfg = readJsonCfg(progCfgFilepath)
	}

	canAccessInternet, errList := gonetcheck.CheckInternetAccess(
		time.Duration( time.Duration(timeoutSeconds) * time.Second ),
		cfg.Urls,
		cfg.TcpAddrs)
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

func readJsonCfg (progCfgFilepath string) progCfg {
	jsonFh, jsonFhErr := os.Open(progCfgFilepath)
	if jsonFhErr != nil { panic(jsonFhErr) }

	fhStat, fhStatErr := jsonFh.Stat()
	if fhStatErr != nil { panic(fhStatErr) }

	jsonBytes := make( []byte, fhStat.Size() )
	_, readErr := jsonFh.Read(jsonBytes)
	if readErr != nil { panic(readErr) }
	jsonFh.Close()

	var cfg progCfg
	unmarshalErr := json.Unmarshal(jsonBytes, &cfg)
	if unmarshalErr != nil { panic(unmarshalErr) }

	return cfg
}

// Script configuration struct
type progCfg struct {
	Urls, TcpAddrs []string
}

// Default configuration
var defaultProcCfg = progCfg{
	Urls : []string{
		"http://www.google.com/",
		"http://www.bing.com/",
		"http://www.microsoft.com/",
		"http://www.apple.com/",
		"http://yahoo.com/",
		"http://www.unimelb.edu.au/",
		"http://abc.net.au/",
		"http://www.monash.edu.au/",
		"http://bbb.co.uk/",
	},
	TcpAddrs : []string{
		"www.google.com:443",
		"www.example.com:80",
		"tty.freeshell.net:22",
	},
}
