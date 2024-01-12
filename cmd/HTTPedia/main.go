package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/rnemeth90/HTTPedia"
	"github.com/spf13/pflag"
)

var (
	statusCode int
)

type config struct {
	statusCode int
}

func init() {
	pflag.IntVarP(&statusCode, "statuscode", "s", 0, "Status Code")
	pflag.Usage = usage
	pflag.ErrHelp = nil
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	fmt.Fprintln(os.Stderr, "This program returns the description of an HTTP status code.")
	fmt.Fprintln(os.Stderr, "Flags:")
	pflag.PrintDefaults()
	fmt.Fprintln(os.Stderr, "\nExamples:")
	fmt.Fprintf(os.Stderr, "  %s -s 404\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "  %s --statuscode=200\n", os.Args[0])
}

func main() {
	pflag.Parse()
	var statusCode int

	nonPflagArgs := pflag.Args()
	if len(nonPflagArgs) >= 1 {
		s, err := strconv.Atoi(nonPflagArgs[0])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		statusCode = s
	}

	c := config{
		statusCode: statusCode,
	}

	result, err := run(c)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	fmt.Fprintln(os.Stdout, result)
}

func run(c config) (string, error) {
	statusCode, err := HTTPedia.GetStatusCode(c.statusCode)
	if err != nil {
		return "", err
	}

	return statusCode.Description, nil
}
