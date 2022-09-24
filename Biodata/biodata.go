package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	if len(args) <= 1 {
		fmt.Println("There is no argument passed.")
		return
	}
	arg, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	if arg < 0 {
		fmt.Println("Argument may not be negative.")
		return
	}
}
