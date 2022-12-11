package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	number         int
	items          []int
	operation      operationFunc
	test           testFunc
	trueTarget     int
	falseTarget    int
	trues          []int
	falses         []int
	inspectedItems int
}

type testFunc func(int) bool

type operationFunc func(int) int

func (m *Monkey) addItems(items []int) {
	m.items = append(m.items, items...)
}

func (m *Monkey) round() {
	m.trues = []int{}
	m.falses = []int{}
	for _, item := range m.items {
		m.inspectedItems++
		worryLevel := m.operation(item)
		worryLevel = worryLevel / 3
		if m.test(worryLevel) {
			m.trues = append(m.trues, worryLevel)
		} else {
			m.falses = append(m.falses, worryLevel)
		}
	}
	m.items = []int{}
}

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Could not open input file: %v\n", err)
	}
	input := strings.Split(string(file), "\n\n")

	var monkeys []*Monkey
	for _, m := range input {
		monkeys = append(monkeys, makeMonkey(m))
	}

	for r := 0; r < 20; r++ {
		for _, m := range monkeys {
			m.round()
			monkeys[m.trueTarget].addItems(m.trues)
			monkeys[m.falseTarget].addItems(m.falses)
		}
	}

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].inspectedItems > monkeys[j].inspectedItems
	})

	for _, m := range monkeys {
		fmt.Printf("Monkey %d made %d inspections\n", m.number, m.inspectedItems)
	}
	fmt.Printf("Part one: %d\n", monkeys[0].inspectedItems*monkeys[1].inspectedItems)

}

func makeMonkey(m string) *Monkey {
	lines := strings.Split(m, "\n")

	number, _ := strconv.Atoi(string(lines[0][7]))
	items := []int{}
	for _, i := range strings.Split(lines[1][18:], ", ") {
		item, _ := strconv.Atoi(i)
		items = append(items, item)
	}
	operand := lines[2][23]
	multiplicand := lines[2][25:]
	var operation operationFunc
	if operand == '*' {
		if multiplicand == "old" {
			operation = func(val int) int { return val * val }
		} else {
			m, _ := strconv.Atoi(multiplicand)
			operation = func(val int) int { return val * m }
		}
	} else if operand == '+' {
		if multiplicand == "old" {
			operation = func(val int) int { return val + val }
		} else {
			m, _ := strconv.Atoi(multiplicand)
			operation = func(val int) int { return val + m }
		}
	}
	divisor, _ := strconv.Atoi(lines[3][21:])
	test := func(val int) bool { return val%divisor == 0 }
	trueTarget, _ := strconv.Atoi(string(lines[4][29]))
	falseTarget, _ := strconv.Atoi(string(lines[5][30]))

	return &Monkey{number: number, items: items, operation: operation, test: test, trueTarget: trueTarget, falseTarget: falseTarget}
}
