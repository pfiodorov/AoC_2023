package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	red   = 12
	green = 13
	blue  = 14
)

func get_cubes(input string) (r int, g int, b int) {
	r = 0
	g = 0
	b = 0
	for _, res := range strings.Split(input, ",") {
		num, _ := strconv.Atoi(strings.Split(res[1:], " ")[0])
		if strings.Contains(res, "green") {
			if g < num {
				g = num
			}
		} else if strings.Contains(res, "blue") {
			if b < num {
				b = num
			}
		} else {
			if r < num {
				r = num
			}
		}
	}
	return
}

func main() {
	file, err := os.Open("test_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum_of_games := 0
	sum_of_powers := 0
	for scanner.Scan() {
		text := scanner.Text()
		text_split := strings.Split(text, ":")
		game_str := strings.Split(text_split[0], " ")[1]
		game_id, _ := strconv.Atoi(game_str)
		game_possible := true
		max_r, max_g, max_b := 0, 0, 0
		for _, res := range strings.Split(text_split[1], ";") {
			r, g, b := get_cubes(res)
			if r > max_r {
				max_r = r
			}
			if g > max_g {
				max_g = g
			}
			if b > max_b {
				max_b = b
			}

			if r > red || g > green || b > blue {
				game_possible = false
			}
		}
		sum_of_powers += max_r * max_g * max_b
		if game_possible {
			sum_of_games += game_id
		}
	}
	fmt.Println(sum_of_games," and ", sum_of_powers)
}
