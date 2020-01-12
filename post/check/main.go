package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	var request map[string]interface{}

	if err := json.NewDecoder(os.Stdin).Decode(&request); err != nil {
		fatal("parsing request", err)
	}

	response := []interface{}{}
	if request["version"] != nil {
		response = append(response, request["version"])
	}

	if err := json.NewEncoder(os.Stdout).Encode(&response); err != nil {
		fatal("serializing response", err)
	}
}

func fatal(doing string, err error) {
	fmt.Fprintf(os.Stderr, "Error "+doing+": "+err.Error()+"\n")
	os.Exit(1)
}
