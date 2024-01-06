package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Cave struct {
	label    string
	rate     int
	open     bool
	children map[string]*Cave
}

type Node struct {
	label     string
	distance  int
	totaldist int
	rate      int
	children  []*Node
}

func (n Node) calcPressure(tick int) int {
	tick += n.distance

	pressure := n.rate * (30 - tick)
	maxChildPressure := 0

	for _, c := range n.children {
		childPressure := c.calcPressure(tick + 1)
		if childPressure > maxChildPressure {
			maxChildPressure = childPressure
		}
	}

	return pressure + maxChildPressure
}

// assign 'valuable' caves
// walk this list - get shortest distance to other valuable nodes - choose most valuable
// valuable caves are worth (time_left - distance) * rate
// remove that cave from valuable list and start process again from there

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Could not open input file: %v\n", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	caves := make(map[string]*Cave)
	re := regexp.MustCompile(`^Valve (\w{2}) .* rate=(\d*); .* to valve[s]* (.*)`)
	for scanner.Scan() {
		line := scanner.Text()
		found := re.FindStringSubmatch(line)

		// fmt.Printf("Valve: %s, rate: %s, next: %s\n", found[1], found[2], found[3])
		rate, _ := strconv.Atoi(found[2])
		childLabels := strings.Split(found[3], ", ")
		children := make(map[string]*Cave)
		for _, c := range childLabels {
			children[c] = &Cave{}
		}
		caves[found[1]] = &Cave{found[1], rate, false, children}
	}

	assignChildren(caves)

	valuable := findValuableCaves(caves)

	resultOne := walkTunnels(caves, valuable, Node{"AA", 0, 0, 0, []*Node{}}, make(map[string]bool))

	fmt.Println(resultOne.calcPressure(0))

}

func assignChildren(caves map[string]*Cave) {
	for _, c := range caves {
		for _, k := range caves {
			if _, exists := k.children[c.label]; exists {
				c.children[k.label] = k
			}
		}
	}
}

func findValuableCaves(caves map[string]*Cave) map[string]*Cave {
	valuable := make(map[string]*Cave)
	for _, c := range caves {
		if c.rate > 0 {
			valuable[c.label] = c
		}
	}
	return valuable
}

func walkTunnels(caves map[string]*Cave, valuable map[string]*Cave, start Node, visited map[string]bool) Node {
	for _, v := range valuable {
		if visited[v.label] {
			continue
		}
		distance := bfs(caves, start.label, v.label)
		// fmt.Printf("%s to %s: %d\n", start.label, v.label, distance)
		newValuable := make(map[string]*Cave)
		for key, val := range valuable {
			if key != v.label {
				newValuable[key] = val
			}
		}
		newVisited := make(map[string]bool)
		for key, val := range visited {
			newVisited[key] = val
		}
		newVisited[start.label] = true
		distancehere := start.totaldist + distance
		if distancehere < 30 {
			thisNode := walkTunnels(caves, newValuable, Node{v.label, distance, distancehere, v.rate, []*Node{}}, newVisited)
			start.children = append(start.children, &thisNode)
		}
	}

	return start
}

func bfs(caves map[string]*Cave, start string, end string) int {
	queue := []Cave{*caves[start]}
	visited := make(map[string]bool)
	visited[start] = true

	pathLength := 0
	for len(queue) > 0 {
		j := len(queue)
		for i := 0; i < j; i++ {
			current := queue[0]
			queue = queue[1:]
			if current.label == end {
				return pathLength
			}
			neighbours := current.children

			for _, n := range neighbours {
				if visited[n.label] {
					continue
				}
				visited[n.label] = true
				queue = append(queue, *n)
			}
		}

		pathLength++
	}

	return pathLength
}
