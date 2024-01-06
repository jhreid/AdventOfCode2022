package main

import (
	"fmt"
	"os"
	"strings"
)

type Rock struct {
	shape [][]int
}

func (r Rock) printRock() {
	for _, row := range r.shape {
		for _, p := range row {
			if p == 0 {
				fmt.Print(".")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
}

func main() {
	file, err := os.ReadFile("test.txt")
	if err != nil {
		fmt.Printf("Could not open input file: %v\n", err)
	}
	jets := strings.Split(string(file), "")

	first := Rock{[][]int{{1, 1, 1, 1}}}
	second := Rock{[][]int{{0, 1, 0}, {1, 1, 1}, {0, 1, 0}}}
	third := Rock{[][]int{{0, 0, 1}, {0, 0, 1}, {1, 1, 1}}}
	fourth := Rock{[][]int{{1}, {1}, {1}, {1}}}
	fifth := Rock{[][]int{{1, 1}, {1, 1}}}

	first.printRock()
	second.printRock()
	third.printRock()
	fourth.printRock()
	fifth.printRock()

	fmt.Println(jets)
}
