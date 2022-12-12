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
	hill[end.x][end.y] = 'z'

	shortest := 412
	for _, start := range starts {
		hill[start.x][start.y] = 'a'
		result := breadthFirstSearch(hill, start, end)
		if result < shortest {
			shortest = result
		}
	}

	fmt.Println(shortest)
}

func breadthFirstSearch(hill [][]byte, start Pos, end Pos) int {
	visited := make(map[Pos]bool)
	visited[start] = true
	queue := []Pos{start}
	steps := 0
	found := false

	for len(queue) > 0 {
		k := len(queue)
		for i := 0; i < k; i++ {
			cur := queue[0]
			queue = queue[1:]
			neighbours := []Pos{{cur.x + 1, cur.y}, {cur.x - 1, cur.y}, {cur.x, cur.y + 1}, {cur.x, cur.y - 1}}

			if cur == end {
				found = true
				return steps
			}

			for _, n := range neighbours {
				if n.x < 0 || n.x >= len(hill) || n.y < 0 || n.y >= len(hill[0]) {
					continue
				}
				if visited[n] {
					continue
				}
				if (hill[n.x][n.y] > hill[cur.x][cur.y]) && (hill[n.x][n.y]-hill[cur.x][cur.y] > 1) {
					continue
				}
				visited[n] = true
				queue = append(queue, n)
			}
		}
		steps++
	}

	if !found {
		return 500
	}

	return steps
}
