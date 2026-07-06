package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	"golang.org/x/term"
)

func main() {
	isJSON := false
	for _, arg := range os.Args {
		if arg == "--json" {
			isJSON = true
			break
		}
	}

	isTTY := term.IsTerminal(int(os.Stdout.Fd()))

	var progressWriter io.Writer
	if isJSON || !isTTY {
		progressWriter = os.Stderr
	} else {
		progressWriter = os.Stdout
	}

	fmt.Fprintln(progressWriter, "Loading...")
	time.Sleep(100 * time.Millisecond)
	fmt.Fprintln(progressWriter, "Processing...")
	time.Sleep(100 * time.Millisecond)

	if isJSON {
		resp := map[string]string{
			"message": "Hello, Bounty Hunter!",
			"status":  "success",
		}
		jsonData, err := json.MarshalIndent(resp, "", "  ")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(string(jsonData))
	} else {
		fmt.Println("Hello, Bounty Hunter!")
	}
}
