package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	var request map[string]interface{}

	{
		err := json.NewDecoder(os.Stdin).Decode(&request)
		if err != nil {
			fatal("parsing request", err)
		}
	}

	response := make(map[string]interface{})
	response["version"] = request["version"]

	{
		err := json.NewEncoder(os.Stdout).Encode(&response)
		if err != nil {
			fatal("serializing response", err)
		}
	}
}

func fatal(doing string, err error) {
	fmt.Fprintf(os.Stderr, "Error "+doing+": "+err.Error()+"\n")
	os.Exit(1)
}
