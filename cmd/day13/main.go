package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"sort"
)

func check(e error) {
	if e != nil {
		panic(e)
	}

}

func main() {
	if len(os.Args) < 2 {
		panic("Provide the input file as an argument")
	}
	input := readInput(os.Args[1])

	// Part 1
	sum := 0
	for i := 0; i < len(input); i += 2 {
		// fmt.Printf("  %v\n  %v\n", input[i], input[i+1])
		result, resolved := compare_list(input[i], input[i+1])
		if !resolved {
			panic("Unexpected")
		}
		if result {
			sum += (i/2 + 1)
		}

	}
	fmt.Printf("Part 1: %d\n", sum)

	// Part 2
	var div1_idx int
	var div2_idx int
	div1 := []any{[]any{float64(2)}}
	input = append(input, div1)
	div2 := []any{[]any{float64(6)}}
	input = append(input, div2)
	sort.SliceStable(input, func(i, j int) bool {
		result, resolved := compare_list(input[i], input[j])
		if !resolved {
			panic("Unexpected")
		}
		return result
	})
	for k, v := range input {
		if div1_idx == 0 && reflect.DeepEqual(v, div1) {
			div1_idx = k + 1
		} else if div2_idx == 0 && reflect.DeepEqual(v, div2) {
			div2_idx = k + 1
			break
		}
	}
	result := div1_idx * div2_idx
	fmt.Printf("Part 2: %d\n", result)

}

func compare_list(l []any, r []any) (bool, bool) {
	var length int
	if len(l) < len(r) {
		length = len(l)
	} else {
		length = len(r)
	}
	// fmt.Printf("Comparing %v vs %v\n", l, r)
	for k := 0; k < length; k++ {
		// fmt.Printf("Comparing %v vs %v\n", l[k], r[k])
		if !is_list(l[k]) && !is_list(r[k]) {
			lv := int(l[k].(float64))
			rv := int(r[k].(float64))
			if lv < rv {
				return true, true
			}
			if lv > rv {
				return false, true
			}
		} else if is_list(l[k]) && is_list(r[k]) {
			lv := l[k].([]any)
			rv := r[k].([]any)
			result, resolved := compare_list(lv, rv)
			if resolved {
				return result, resolved
			}
		} else if is_list(l[k]) && !is_list(r[k]) {
			lv := l[k].([]any)
			rv := []any{r[k].(float64)}
			result, resolved := compare_list(lv, rv)
			if resolved {
				return result, resolved
			}
		} else if !is_list(l[k]) && is_list(r[k]) {
			lv := []any{l[k].(float64)}
			rv := r[k].([]any)
			result, resolved := compare_list(lv, rv)
			if resolved {
				return result, resolved
			}
		}
		// fmt.Printf("  %d, %T\n", v.(int), v)
	}
	if len(l) < len(r) {
		return true, true
	} else if len(r) < len(l) {
		return false, true
	}

	return true, false
}

func is_list(i any) bool {
	switch i.(type) {
	case []any:
		return true
	default:
		return false
	}
}

type Pair struct {
	l any
	r any
}

func readInput(filename string) [][]any {
	readFile, err := os.Open(filename)
	check(err)
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var input [][]any
	for fileScanner.Scan() {
		var l, r []any
		line := fileScanner.Text()
		json.Unmarshal([]byte(line), &l)
		input = append(input, l)
		fileScanner.Scan()
		line = fileScanner.Text()
		json.Unmarshal([]byte(line), &r)
		input = append(input, r)
		fileScanner.Scan()

	}

	return input
}
