package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func process_inputs(filename string) ([]int, []int) {
	myFile, err := os.Open(filename)
	left := make([]int, 0)
	right := make([]int, 0)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(myFile)
	line := 0
	for scanner.Scan() {
		vals := strings.Fields(scanner.Text())
		line++
		if l, err := strconv.Atoi(vals[0]); err == nil {

			left = append(left, l)
		}
		if r, err := strconv.Atoi(vals[1]); err == nil {

			right = append(right, r)
		}
	}
	return left, right

}
func main() {
	left, right := process_inputs("day1.txt")
	left_lookup := make(map[int]int)
	right_lookup := make(map[int]int)
	for _, k := range left {
		_, ok := left_lookup[k]
		if !ok {
			left_lookup[k] = 1
		} else {
			left_lookup[k] += 1
		}
	}
	for _, k := range right {
		_, ok := right_lookup[k]
		if !ok {
			right_lookup[k] = 1
		} else {
			right_lookup[k] += 1
		}
	}
	ans := 0
	for k, v := range left_lookup {
		val, ok := right_lookup[k]
		if !ok {
			ans += 0
		} else {
			ans += val * k * v
		}
	}
	fmt.Println(ans)
}
