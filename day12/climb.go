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

type Node struct {
	p     Pos
	count int
}

type Queue []Node

func (q Queue) contains(node Node) bool {
	for _, n := range q {
		if n.p == node.p {
			return true
		}
	}
	return false
}

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Could not open input file: %v\n", err)
	}
	input := strings.Split(string(file), "\n")

	var start Pos
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
			if hill[i][j] == 'S' {
				start = Pos{i, j}
			}
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

	astar(hill, Node{start, 0}, Node{end, 0})
}

func astar(hill [][]byte, start Node, end Node) {
	queue := Queue{end}

	i := 0
	for i < len(queue) {
		current := queue[i]
		if current.p == start.p {
			break
		}
		neighbours := []Node{{Pos{current.p.x + 1, current.p.y}, i + 1}, {Pos{current.p.x - 1, current.p.y}, i + 1}, {Pos{current.p.x, current.p.y + 1}, i + 1}, {Pos{current.p.x, current.p.y - 1}, i + 1}}
		println(neighbours)

		for _, n := range neighbours {
			if n.p.x < 0 || n.p.x >= len(hill) || n.p.y < 0 || n.p.y >= len(hill[0]) {
				continue
			}
			if queue.contains(n) {
				continue
			}
			if (hill[current.p.x][current.p.y] > hill[n.p.x][n.p.y]) && (hill[current.p.x][current.p.y]-hill[n.p.x][n.p.y] > 1) {
				continue
			}
			queue = append(queue, n)
		}

		i++
	}

	fmt.Println(i)

	visited := make(map[Pos]string)
	for _, n := range queue {
		visited[n.p] = string(hill[n.p.x][n.p.y])
	}
	for x := 0; x < len(hill); x++ {
		for y := 0; y < len(hill[0]); y++ {
			n, exist := visited[Pos{x, y}]
			if exist {
				fmt.Print(n)
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}

	fmt.Println()
	fmt.Println()

	fmt.Println()
	fmt.Println()
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
