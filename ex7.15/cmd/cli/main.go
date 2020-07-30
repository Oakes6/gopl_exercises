package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"gopl_exercises/ex7.15"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter expression to evaluate: ")
	expr, err := reader.ReadString('\n')

	// attempt to parse expression
	parsedExpr, err := eval.Parse(expr)
	if err != nil {
		// return error message to standard input and exit
		fmt.Println("Cannot parse expression...")
		fmt.Println("Terminating...")
		os.Exit(1)
	}

	// check syntax rules of expression
	vars := make(map[eval.Var]bool)
	err = parsedExpr.Check(vars)
	if err != nil {
		fmt.Println("Error with the syntax of the expression")
		fmt.Println("Please try again...")
		os.Exit(1)
	}

	// prompt user for values to the args provided in the expression and populated in 'vars'
	env := eval.Env(make(map[eval.Var]float64))

	for k := range vars {
		fmt.Printf("Enter value for the %v variable: ", k)
		got, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error with the given argument: %s\n", err)
			fmt.Println("Terminating...")
			os.Exit(1)
		}

		// try convert to float, then to literal
		// if successful add to 'Env'
		parsedLiteral, err := strconv.ParseFloat(got, 64)
		if err != nil {
			fmt.Println("Error parsing the given float value")
			fmt.Println("Terminating...")
			os.Exit(1)
		}
		env[k] = parsedLiteral

	}

	// use accumulated values in 'Env' to evaluate the expression
	result := parsedExpr.Eval(env)

	fmt.Printf("Final answer: %v\n", result)
}
