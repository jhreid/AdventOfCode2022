package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type elfDir struct {
	parent *elfDir
	name   string
	files  map[string]*elfFile
	dirs   map[string]*elfDir
}

type elfFile struct {
	name string
	size int64
}

func (ed elfDir) calc_size() int64 {
	total := int64(0)
	for _, f := range ed.files {
		total += f.size
	}
	for _, d := range ed.dirs {
		total += d.calc_size()
	}
	return total
}

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Could not open input file: %v\n", err)
	}
	input := strings.Split(string(file), "\n")

	allFolders := []*elfDir{}

	root := new(elfDir)
	root.name = "/"
	root.dirs = make(map[string]*elfDir)
	root.files = make(map[string]*elfFile)
	allFolders = append(allFolders, root)
	currentDir := root
outerLoop:
	for _, line := range input {
		if line == "" {
			for {
				if currentDir.name == "/" {
					break outerLoop
				}
				currentDir = currentDir.parent
			}
		}
		tokens := strings.Split(line, " ")
		if tokens[0] == "dir" {
			newDir := new(elfDir)
			newDir.name = tokens[1]
			newDir.parent = currentDir
			newDir.dirs = make(map[string]*elfDir)
			newDir.files = make(map[string]*elfFile)
			currentDir.dirs[newDir.name] = newDir
			allFolders = append(allFolders, newDir)
		} else if tokens[1] == "cd" {
			if tokens[2] == "/" {
				continue
			}
			if tokens[2] == ".." {
				currentDir = currentDir.parent
			} else {
				currentDir = currentDir.dirs[tokens[2]]
			}
		} else {
			size, err := strconv.ParseInt(tokens[0], 10, 64)
			if err == nil {
				newFile := new(elfFile)
				newFile.name = tokens[1]
				newFile.size = size
				currentDir.files[tokens[1]] = newFile
			}
		}
	}

	var total int64 = 0
	for _, d := range allFolders {
		size := d.calc_size()
		if size <= 100000 {
			total += size
		}
	}

	fmt.Println(total)

	filtered := []int64{}
	for _, d := range allFolders {
		size := d.calc_size()
		if size >= 30000000-(70000000-root.calc_size()) {
			filtered = append(filtered, size)
		}
	}
	smallest := filtered[0]
	for _, s := range filtered {
		if s < smallest {
			smallest = s

		}
	}

	fmt.Println(smallest)
}
