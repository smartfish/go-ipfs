package main

import (
	"os"

	"encoding/json"
	"fmt"

	"golang.org/x/tools/benchmark/parse"
)

func main() {
	if err := mainErr(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}

func mainErr() error {
	benches, err := parse.ParseSet(os.Stdin)
	if err != nil {
		return err
	}

	for k, bench := range benches {
		fmt.Printf("%s:\n", k)
		b, err := json.Marshal(bench)
		if err != nil {
			return err
		}

		fmt.Printf("    %s\n", b)
	}

	return nil
}
