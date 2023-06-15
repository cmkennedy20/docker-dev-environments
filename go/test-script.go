package main

import (
	"fmt"
	"os"
)

func main() {
	// Specify the name of the environment variable you want to print
	envVariableName := "INPUT_ENV_VAR"

	// Retrieve the value of the environment variable
	envVariableValue := os.Getenv(envVariableName)

	// Check if the environment variable is set
	fmt.Printf("Hello world\n Environment var: %s \n",envVariableValue)
}
