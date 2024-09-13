package main

import (
	"fmt"
	"os"

	"github.com/bLittle1996/ledger/internal/cli"
)

func main() {
	err := cli.Run(os.Args[1:])

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
