package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func parse_ints(text string) int {
	val, err := strconv.Atoi(text)
	if err != nil {
		log.Fatal(err)
	}
	return val
}

func part_1(text string) int {
	pattern := `mul\((\d{1,3}),(\d{1,3})\)`
	rp := regexp.MustCompile(pattern)
	matches := rp.FindAllStringSubmatch(text, -1)
	ans := 0
	for _, vals := range matches {
		ans += parse_ints(vals[1]) * parse_ints(vals[2])
	}
	return ans

}

func part_2(text string) int {
	pattern := `(?s)don't\(\).*?(?:$|do\(\))`
	rp := regexp.MustCompile(pattern)
	matches := rp.ReplaceAllString(text, "")
	return part_1(matches)
}

func process_inputs(file_name string) string {
	reader, err := os.ReadFile(file_name)
	if err != nil {
		log.Fatal(err)
	}
	value := string(reader)
	return value

}
func main() {
	file := "inputs.txt"
	text := process_inputs(file)
	// fmt.Println(part_1(text))
	fmt.Println(part_2(text))
	// part_2(text)
}
