package main

import (
	"log"
	"os"
)

func main() {

	file, err := os.Open("input_3.txt")
	if err != nil {
		log.Fatal(err)
	}
	solution(file)
}

func solution(f *os.File) int {

}
