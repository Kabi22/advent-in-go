package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	file_name := "input_2.txt"
	fmt.Printf("%d", solution(file_name))
}

func solution(fname string) int {
	f, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	// bag contained only 12 red cubes, 13 green cubes, and 14 blue cubes?

	result := 0

	for scanner.Scan() {
		s := scanner.Text()
		split := strings.Split(s, ":")
		if len(split) < 2 {
			log.Fatal(": split failed, array too small")
		}
		game_no, err := strconv.Atoi((strings.Split(split[0], " "))[1])
		if err != nil {
			log.Fatal(err)
		}
		hands := split[1]

		// count balls
		fmt.Println(s)
		red_cnt, blue_cnt, green_cnt := 0, 0, 0
		for _, k := range strings.Split(hands, ";") {
			// fmt.Println(k)
			for _, b := range strings.Split(k, ",") {
				j := strings.Split(b, " ")[1:]

				num, _ := strconv.Atoi(j[0])
				// fmt.Printf("%q\n", j)
				switch j[1] {
				case ("red"):
					red_cnt = Max(num, red_cnt)
				case ("green"):
					green_cnt = Max(num, green_cnt)
				case ("blue"):
					blue_cnt = Max(num, blue_cnt)
				default:
					log.Fatal(b[1])
				}

				// fmt.Printf("num: %d -- %d %d %d\n", num, red_cnt, green_cnt, blue_cnt)

			}

		}
		if red_cnt <= 12 && green_cnt <= 13 && blue_cnt <= 14 {
			result += game_no
			fmt.Println(game_no)
			fmt.Printf("%d %d %d\n", red_cnt, blue_cnt, green_cnt)
		}

	}
	return result
}

func Max(n1 int, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

func solution_part_2() int {

}
