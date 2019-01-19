package main

import (
	"fmt"
	"os"

	"github.com/SteveCastle/openmob/packages/shrike/src/pkg/cmd"
)

func main() {
	if err := cmd.RunServer(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
