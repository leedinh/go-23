package util

import (
	"fmt"
	"sort"
	"strconv"
)

func SortInt(input []string) ([]int, error) {
	var intSlice []int
	for _, str := range input {
		if i, err := strconv.Atoi(str); err == nil {
			intSlice = append(intSlice, i)
		} else {
			return nil, err // Return an error if parsing fails
		}
	}
	sort.Ints(intSlice)
	return intSlice, nil
}

func SortFloat(input []string) ([]float64, error) {
	var floatSlice []float64
	for _, str := range input {
		if f, err := strconv.ParseFloat(str, 64); err == nil {
			floatSlice = append(floatSlice, f)
		} else {
			return nil, err // Return an error if parsing fails
		}
	}
	sort.Float64s(floatSlice)
	return floatSlice, nil
}

func SortString(input []string) []string {
	sort.Strings(input)
	return input
}

func SortMix(input []string) ([]float64, []string) {
	var numberSlice []float64
	var stringSlice []string

	for _, str := range input {
		if f, err := strconv.ParseFloat(str, 64); err == nil {
			numberSlice = append(numberSlice, f)
		} else {
			stringSlice = append(stringSlice, str)
		}
	}
	sort.Float64s(numberSlice)
	sort.Strings(stringSlice)
	return numberSlice, stringSlice
}

func PrintSlice(input ...[]string) {
	for _, str := range input {
		println(str)
	}
}

func IntOrFloatSliceToStringSlice(input interface{}) ([]string, error) {
	switch slice := input.(type) {
	case []int:
		var output []string
		for _, num := range slice {
			output = append(output, strconv.Itoa(num))
		}
		return output, nil
	case []float64:
		var output []string
		for _, num := range slice {
			output = append(output, strconv.FormatFloat(num, 'f', -1, 64))
		}
		return output, nil
	default:
		return nil, fmt.Errorf("unsupported type %T", slice)
	}
}
