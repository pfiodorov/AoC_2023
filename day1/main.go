package main

import (
    "log"
    "bufio"
    "fmt"
    "os"
    "regexp"
    "strconv"
    "strings"
)

func main() {
    file, err := os.Open("test_input2.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    
    scanner := bufio.NewScanner(file)
    sum := 0
    for scanner.Scan() {
        text := scanner.Text()
        re := regexp.MustCompile("[0-9]+")
        clean_str_arr := re.FindAllString(text, -1)
	clean_str := strings.Join(clean_str_arr, "")
        res, _ := strconv.Atoi(string(clean_str[0]) + string(clean_str[len(clean_str)-1]))
// Debug output
//	fmt.Println(clean_str, " eq ", res)
        sum = sum + res
    }
    fmt.Println("Done")
    fmt.Println(sum)
}
