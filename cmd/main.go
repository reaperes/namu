package main

import (
	"fmt"
	"github.com/reaperes/namu/pkg/cmd"
	"os"
)

func main() {
	namu := cmd.Root()

	if err := namu.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
