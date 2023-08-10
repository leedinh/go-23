package main

import (
	"fmt"
	"strconv"
)

func numDifferentInteger(word string) int {
	var count int
	var num string
	var numMap = make(map[int]bool)
	for _, c := range word {
		if c >= '0' && c <= '9' {
			num += string(c)
		} else {
			if len(num) > 0 {
				tmp, _ := strconv.Atoi(num)
				numMap[tmp] = true
				num = ""
			}
		}
	}
	if len(num) > 0 {
		tmp, _ := strconv.Atoi(num)
		numMap[tmp] = true
	}
	count = len(numMap)
	return count
}

func main() {
	word := "A1b01c001"

	count := numDifferentInteger(word)
	fmt.Println(count)

}
