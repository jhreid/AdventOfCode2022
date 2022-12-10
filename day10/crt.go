package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Could not open input file: %v\n", err)
	}
	defer file.Close()

	x := 1
	cycles := 0
	var values []int
	screen := [240]int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), " ")
		switch input[0] {
		case "noop":
			{
				cycles += 1
				values = checkCycle(x, cycles, values)
				screen = checkScreen(x, cycles, screen)
			}
		case "addx":
			{
				cycles += 1
				values = checkCycle(x, cycles, values)
				screen = checkScreen(x, cycles, screen)
				cycles += 1
				values = checkCycle(x, cycles, values)
				screen = checkScreen(x, cycles, screen)
				val, _ := strconv.Atoi(input[1])
				x += val
			}
		}
	}

	fmt.Printf("x is %d\n", x)
	fmt.Printf("values %v\n", values)

	signal := sum(values)
	fmt.Printf("Total strength: %d\n", signal)

	drawScreen(screen)
}

func checkScreen(x int, cycle int, screen [240]int) [240]int {
	beamPos := cycle%40 - 1
	if beamPos >= x-1 && beamPos <= x+1 {
		screen[cycle-1] = 1
	}
	return screen
}

func checkCycle(x int, cycle int, values []int) []int {
	if (cycle+20)%40 == 0 {
		values = append(values, cycle*x)
	}
	return values
}

func sum(s []int) int {
	result := 0
	for _, v := range s {
		result += v
	}
	return result
}

func drawScreen(screen [240]int) {
	for line := 0; line < 6; line++ {
		for _, pixel := range screen[line*40 : line*40+40] {
			if pixel == 1 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
