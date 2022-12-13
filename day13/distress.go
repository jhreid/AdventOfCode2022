package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

type Pair struct {
	a string
	b string
}

func (p Pair) inOrder() bool {
	result := leftSmaller(p.a, p.b)
	return result
}

func leftSmaller(left string, right string) bool {
	if rune(left[0]) == rune(right[0]) {
		return leftSmaller(left[1:], right[1:])
	}
	if rune(left[0]) == ']' && (rune(right[0]) == '[' || rune(right[0]) == ',' || unicode.IsDigit(rune(right[0]))) {
		return true
	}
	if unicode.IsDigit(rune(left[0])) && rune(right[0]) == '[' {
		if strings.ContainsRune(left, ',') {
			left = "[" + string(left[0:strings.IndexRune(left, ',')]) + "]" + left[strings.IndexRune(left, ','):]
		} else {
			left = "[" + string(left[0]) + "]" + left[1:]
		}
		return leftSmaller(left[1:], right[1:])
	}
	if (unicode.IsDigit(rune(left[0])) || rune(left[0]) == '[' || rune(left[0]) == ',') && rune(right[0]) == ']' {
		return false
	}
	if unicode.IsDigit(rune(right[0])) && rune(left[0]) == '[' {
		if strings.ContainsRune(right, ',') {
			right = "[" + string(right[0:strings.IndexRune(right, ',')]) + "]" + right[strings.IndexRune(right, ','):]
		} else {
			right = "[" + string(right[0]) + "]" + right[1:]
		}
		return leftSmaller(left[1:], right[1:])
	}
	if unicode.IsDigit(rune(left[0])) && unicode.IsDigit(rune(right[0])) {
		re := regexp.MustCompile("^([0-9]+)")
		found := re.FindAllString(left, -1)
		a, _ := strconv.Atoi(found[0])
		found = re.FindAllString(right, -1)
		b, _ := strconv.Atoi(found[0])
		if a < b {
			return true
		} else {
			return false
		}
	}

	return false
}

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Could not open input file: %v\n", err)
	}
	input := strings.Split(string(file), "\n\n")

	pairs := make([]Pair, 0)
	for _, line := range input {
		rows := strings.Split(line, "\n")
		pairs = append(pairs, Pair{rows[0], rows[1]})
	}

	partOne := 0
	for i, pair := range pairs {
		if pair.inOrder() {
			partOne += (i + 1)
		}
	}

	println(partOne)

	input2 := make([]string, 0)
	for _, s := range input {
		pair := strings.Split(s, "\n")
		input2 = append(input2, pair[0])
		input2 = append(input2, pair[1])
	}
	input2 = append(input2, "[[2]]")
	input2 = append(input2, "[[6]]")

	sort.Slice(input2, func(i, j int) bool {
		return leftSmaller(input2[i], input2[j])
	})

	partTwo := 1
	for i, s := range input2 {
		if s == "[[2]]" || s == "[[6]]" {
			partTwo = partTwo * (i + 1)
		}
	}

	fmt.Println(partTwo)
}
