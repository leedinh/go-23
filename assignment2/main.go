package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/leedinh/go-23/assignment2/util"
)

func main() {
	flag := os.Args[1]
	switch flag {
	case "-int":
		{
			sorted, err := util.SortInt(os.Args[2:])
			if err != nil {
				panic(err)
			}
			res, _ := util.IntOrFloatSliceToStringSlice(sorted)
			fmt.Println("Output:", strings.Join(res, " "))
		}
	case "-float":
		{
			sorted, err := util.SortFloat(os.Args[2:])
			if err != nil {
				panic(err)
			}
			res, _ := util.IntOrFloatSliceToStringSlice(sorted)
			fmt.Println("Output:", strings.Join(res, " "))
		}
	case "-string":
		{
			sorted := util.SortString(os.Args[2:])
			fmt.Println("Output:", strings.Join(sorted, " "))
		}
	case "-mix":
		{
			sortednum, sortedstring := util.SortMix(os.Args[2:])
			numbers_str, _ := util.IntOrFloatSliceToStringSlice(sortednum)
			fmt.Println("Output:", strings.Join(numbers_str, " "), strings.Join(sortedstring, " "))

		}
	default:
		{
			panic("Invalid flag")
		}
	}

}
