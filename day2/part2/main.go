package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"

	"strconv"
	"strings"
)

func check_increasing(level []int) bool {
	i := 0
	j := 1
	// fmt.Println(level)
	for i < j && j < len(level) {
		if level[j] > level[i] && level[j]-level[i] <= 3 && level[j]-level[i] >= 1 {
			i++
			j++
		} else {
			// return false
			return false
		}
	}
	return true
}

func check_decreasing(level []int) bool {
	i := 0
	j := 1
	// fmt.Println(level)
	for i < j && j < len(level) {
		if level[j] < level[i] && level[i]-level[j] <= 3 && level[i]-level[j] >= 1 {
			i++
			j++
		} else {
			return false
		}
	}
	return true
}

func count_valid_levels(levels [][]int) int {
	count := 0
	for _, level := range levels {
		if check_decreasing(level) || check_increasing(level) {
			count++
		} else {
			for i := range level {
				new_lvl := slices.Delete(slices.Clone(level), i, i+1)
				if check_decreasing(new_lvl) || check_increasing(new_lvl) {
					count++
					break
				}
			}
		}
	}
	return count
}

func strings_slice_int_slice(lists []string) []int {
	ans := make([]int, len(lists))
	var err error
	for i := 0; i < len(lists); i++ {
		ans[i], err = strconv.Atoi(lists[i])
		if err != nil {
			log.Fatal(err)
		}
	}
	return ans
}
func process_inputs(filename string) [][]int {
	buffer, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer buffer.Close()
	lines := bufio.NewScanner(buffer)
	levels := make([][]int, 0)
	for lines.Scan() {
		level := strings.Fields(lines.Text())
		levels = append(levels, strings_slice_int_slice(level))
	}
	return levels
}
func main() {
	filename := "inputs.txt"
	levels := process_inputs(filename)
	ans := count_valid_levels(levels)
	fmt.Println(ans)
}
