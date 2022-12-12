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
	file, err := os.ReadFile("test.txt")
	if err != nil {
		fmt.Printf("Could not open input file: %v\n", err)
	}
	input := strings.Split(string(file), "\n")

	var start Pos
	var end Pos

	hill := [][]byte{}
	for _, row := range input {
		if row != "" {
			hill = append(hill, []byte(row))
		}
	}
	for i := range hill {
		for j := range hill[i] {
			if hill[i][j] == 'S' {
				start = Pos{i, j}
			} else if hill[i][j] == 'E' {
				end = Pos{i, j}
			}
		}
	}
	hill[end.x][end.y] = 'z' + 1
	hill[start.x][start.y] = 'a' - 1

	shortest := breadthFirstSearch(hill, start, end)

	fmt.Printf("Starting at %v and ending at %v\n\n", start, end)
	fmt.Println(shortest)
}

func breadthFirstSearch(hill [][]byte, start Pos, end Pos) int {
	visited := make(map[Pos]bool)
	visited[start] = true
	queue := []Pos{start}
	steps := 0

	for len(queue) > 0 {
		k := len(queue)
		for i := 0; i < k; i++ {
			cur := queue[0]
			queue = queue[1:]
			neighbours := []Pos{{cur.x + 1, cur.y}, {cur.x - 1, cur.y}, {cur.x, cur.y + 1}, {cur.x, cur.y - 1}}

			if cur == end {
				return steps
			}

			for _, n := range neighbours {
				if n.x < 0 || n.x >= len(hill) || n.y < 0 || n.y >= len(hill[0]) {
					continue
				}
				if visited[n] {
					continue
				}
				if hill[n.x][n.y] > hill[cur.x][cur.y] && hill[n.x][n.y]-hill[cur.x][cur.y] > 1 {
					continue
				}
				visited[n] = true
				queue = append(queue, n)
			}
		}
		steps++
	}

	return steps
}
