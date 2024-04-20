package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func main() {
	file, err := os.Open("input_1.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		s := scanner.Text()

		nums := []int{}
		for _, char := range s {
			if unicode.IsDigit(char) {
				nums = append(nums, int(char-'0'))
			}
		}
		fmt.Println(s, nums)

		first, last := nums[0], nums[len(nums)-1]
		f, l := strconv.Itoa(first), strconv.Itoa(last)

		num, _ := strconv.Atoi((f + l))

		sum += num
	}

	fmt.Println(sum)

}
