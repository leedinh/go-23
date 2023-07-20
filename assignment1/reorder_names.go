package main

import (
	"fmt"
	"os"
	"strings"
)

func do_reorder(firstname string, lastname string, lang string) (first string, last string) {
	switch lang {
	case "vn", "cn", "jp", "kr", "tw":
		{
			first = lastname
			last = firstname
		}
	case "us":
		{
			first = firstname
			last = lastname
		}
	default:
		{
			panic("Language not supported")
		}
	}
	return
}

func parse_input(args []string) (firstname string, lastname string, middle string) {
	firstname = args[1]
	lastname = args[2]
	middle = strings.Join(args[3:len(args)-1], " ")
	return
}

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

	// parse the input
	firstname, lastname, middle := parse_input(os.Args)
	first, last := do_reorder(firstname, lastname, lang)

	// Print the output
	if middle == "" {
		fmt.Printf("Output: %s %s", first, last)
	} else {
		fmt.Printf("Output: %s %s %s", first, middle, last)
	}
}
