package main

import (
	"fmt"
	"os"

	"github.com/renan-campos/json-config/configuration"
)

func main() {
	configuration, err := configuration.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load configuration: %v\n", err)
		return
	}
	fmt.Printf("Program Name: %s\n", configuration.ProgramName)
	fmt.Println(configuration.HelloMessage)
	fmt.Println(configuration.ByeMessage)
}
