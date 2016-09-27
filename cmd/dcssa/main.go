package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/regalskeppet/dcssa"
)

func main() {
	data := dcssa.NewData()
	err := dcssa.ParseDir(".", data)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	enc := json.NewEncoder(os.Stdout)
	err = enc.Encode(data)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
