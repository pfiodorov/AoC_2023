package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func isInt(input string) (num int, is_int bool) {
	number_strings := [9]string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	// Comment this out to pass day 1 part 1
	written_strings := [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	for index, a := range number_strings {
		if strings.Contains(input, a) {
			num = index + 1
			is_int = true
			return
		}
	}

	for index, a := range written_strings {
		if strings.Contains(input, a) {
			num = index + 1
			is_int = true
			return
		}
	}
	return 0, false
}

func findFirstInt(input string) (num string) {
	var buf string
	for i := 0; i < len(input); i++ {
		buf += string(input[i])
		num, is_int := isInt(string(buf))
		if is_int {
			return strconv.Itoa(num)
		}
	}
	return "0"
}

func findLastInt(input string) (num string) {
	var buf string
	for i := len(input) - 1; i >= 0; i-- {
		buf = string(input[i]) + buf
		num, is_int := isInt(string(buf))
		if is_int {
			return strconv.Itoa(num)
		}
	}
	return "0"
}

func main() {
	file, err := os.Open("test_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		text := scanner.Text()
		firstInt := findFirstInt(text)
		lastInt := findLastInt(text)
		res, _ := strconv.Atoi(firstInt + lastInt)
		sum = sum + res
	}
	fmt.Println("Done")
	fmt.Println(sum)
}
