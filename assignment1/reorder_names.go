package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args_len := len(os.Args)
	// Check if the number of arguments is valid
	if args_len < 4 {
		panic("Invalid number of arguments")
	}

	// Check if the language is valid
	lang := strings.ToLower(os.Args[args_len-1])
	if len(lang) != 2 {
		panic("Invalid language")
	}

	firstname := os.Args[1]
	lastname := os.Args[2]
	var first string
	var second string
	switch lang {
	case "vn", "cn", "jp", "kr", "tw":
		{
			first = lastname
			second = firstname
		}
	default:
		{
			first = firstname
			second = lastname
		}
	}

	// Print the output
	if args_len == 4 {
		fmt.Printf("Output: %s %s", first, second)
	} else {
		fmt.Printf("Output: %s %s %s", first, second, strings.Join(os.Args[3:args_len-1], " "))
	}
}
