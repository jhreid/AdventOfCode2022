package main

import (
	"fmt"
	"os"
	"strings"
)

type Pos struct {
	x int
	y int
}

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Could not open input file: %v\n", err)
	}
	input := strings.Split(string(file), "\n")

	var starts []Pos
	var end Pos

	hill := [][]byte{}
	for _, row := range input {
		if row != "" {
			hill = append(hill, []byte(row))
		}
	}
	for i := range hill {
		for j := range hill[i] {
			if hill[i][j] == 'S' || hill[i][j] == 'a' {
				starts = append(starts, Pos{i, j})
			} else if hill[i][j] == 'E' {
				end = Pos{i, j}
			}
		}
	}
	hill[end.x][end.y] = 'z' + 1

	shortest := 412
	for _, start := range starts {
		hill[start.x][start.y] = 'a'
		result := astar(hill, start, end)
		if result < shortest {
			shortest = result
		}
	}

	fmt.Println(shortest)

}

func astar(hill [][]byte, start Pos, end Pos) int {
	queue := []Pos{end}
	visited := make(map[Pos]bool)
	visited[end] = true

	pathLength := 0
	for len(queue) > 0 {
		j := len(queue)
		for i := 0; i < j; i++ {
			current := queue[0]
			queue = queue[1:]
			if current == start {
				return pathLength
			}
			neighbours := []Pos{{current.x + 1, current.y}, {current.x - 1, current.y}, {current.x, current.y + 1}, {current.x, current.y - 1}}

			for _, n := range neighbours {
				if n.x < 0 || n.x >= len(hill) || n.y < 0 || n.y >= len(hill[0]) {
					continue
				}
				if visited[n] {
					continue
				}
				if (hill[current.x][current.y] > hill[n.x][n.y]) && (hill[current.x][current.y]-hill[n.x][n.y] > 1) {
					continue
				}
				visited[n] = true
				queue = append(queue, n)
			}
		}

		pathLength++
	}

	return pathLength
}
