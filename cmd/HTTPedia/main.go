package main

import (
	"encoding/json"
	"fmt"
	"os"

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
}

func usage() {
	pflag.PrintDefaults()
}

func main() {
	pflag.Parse()

	c := config{
		statusCode: statusCode,
	}

	result, err := run(c)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}

	jsonStr, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}

	fmt.Fprintln(os.Stdout, jsonStr)
}

func run(c config) (string, error) {
	description, err := HTTPedia.GetStatusCode(c.statusCode)
	if err != nil {
		return "", err
	}

	return description, nil
}
