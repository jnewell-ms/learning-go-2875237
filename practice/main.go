package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	//NOTE: Characters are wrapped in single quotes, strings are wrapped in double quotes
	//error output is ignored using '_'
	input, _ := reader.ReadString('\n')
	fmt.Println("You entered:", input)

	fmt.Print("Enter a number: ")
	numInput, _ := reader.ReadString('\n')
	//Expect there to possibly be empty space before the number, so use strings.TrimSpace to deal with that, also using 64 bits ~ since 64 bit processor
	aFloat, err := strconv.ParseFloat(strings.TrimSpace(numInput), 64)

	//This is checking if there is an error
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Value of number:", aFloat)
	}

}
