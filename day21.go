package adventofcode2022

import (
	"strconv"
	"strings"
)

func Day21(lines []string, part1 bool) int {
	if part1 {
		return day21Part1(lines)
	}
	return day21Part2(lines)
}

func day21Part1(lines []string) int {
	var (
		chans    = make(map[string]chan int, len(lines))
		formulas = make(map[string]Formula, len(lines))

		mkc = func(name string) chan int {
			c := make(chan int, 1)
			chans[strings.TrimSpace(name)] = c
			return c
		}
	)

	// phase 1: create immediate channels and formulas
	for _, line := range lines {
		ss := strings.Split(line, ":")

		c := mkc(ss[0])
		// number or formula?
		n, err := strconv.Atoi(strings.TrimSpace(ss[1]))
		if err == nil {
			c <- n
		} else {
			ops := strings.Fields(ss[1])
			formulas[ss[0]] = Formula{
				left:      ops[0],
				operation: ops[1][0],
				right:     ops[2],
			}
		}
	}

	// phase 2: channel plumbing
	for k, v := range formulas {
		go func(name string, f Formula) {
			var a, b int
			for i := 0; i < 2; i++ {
				select {
				case a = <-chans[f.left]:
				case b = <-chans[f.right]:
				}
			}
			c := chans[name]
			switch f.operation {
			case '+':
				c <- a + b
			case '-':
				c <- a - b
			case '*':
				c <- a * b
			case '/':
				c <- a / b
			default:
				c <- 0 // Unknown operation, return 0
			}
		}(k, v)
	}
	n := <-chans["root"]
	return n
}

func day21Part2(lines []string) int {
	monkeys := parseDay21Monkeys(lines)

	// Root should perform equality check
	rootFormula := monkeys["root"].formula

	// Determine which side contains humn
	leftHasHumn := containsHumn(rootFormula.left, monkeys)

	var target int
	var exprWithHumn string

	if leftHasHumn {
		target = evaluate(rootFormula.right, monkeys)
		exprWithHumn = rootFormula.left
	} else {
		target = evaluate(rootFormula.left, monkeys)
		exprWithHumn = rootFormula.right
	}

	// Solve for humn by working backwards
	return solveForHumn(exprWithHumn, target, monkeys)
}

func parseDay21Monkeys(lines []string) map[string]Day21Monkey {
	monkeys := make(map[string]Day21Monkey, len(lines))

	for _, line := range lines {
		parts := strings.Split(line, ":")
		name := strings.TrimSpace(parts[0])
		expr := strings.TrimSpace(parts[1])

		n, err := strconv.Atoi(expr)
		if err == nil {
			monkeys[name] = Day21Monkey{value: n, hasValue: true}
		} else {
			ops := strings.Fields(expr)
			monkeys[name] = Day21Monkey{
				hasValue: false,
				formula: Formula{
					left:      ops[0],
					operation: ops[1][0],
					right:     ops[2],
				},
			}
		}
	}

	return monkeys
}

func containsHumn(name string, monkeys map[string]Day21Monkey) bool {
	if name == "humn" {
		return true
	}

	monkey := monkeys[name]
	if monkey.hasValue {
		return false
	}

	return containsHumn(monkey.formula.left, monkeys) || containsHumn(monkey.formula.right, monkeys)
}

func evaluate(name string, monkeys map[string]Day21Monkey) int {
	monkey := monkeys[name]

	if monkey.hasValue {
		return monkey.value
	}

	left := evaluate(monkey.formula.left, monkeys)
	right := evaluate(monkey.formula.right, monkeys)

	switch monkey.formula.operation {
	case '+':
		return left + right
	case '-':
		return left - right
	case '*':
		return left * right
	case '/':
		return left / right
	default:
		return 0
	}
}

func solveForHumn(name string, target int, monkeys map[string]Day21Monkey) int {
	if name == "humn" {
		return target
	}

	monkey := monkeys[name]
	leftHasHumn := containsHumn(monkey.formula.left, monkeys)

	var newTarget int
	var nextName string

	if leftHasHumn {
		rightValue := evaluate(monkey.formula.right, monkeys)
		switch monkey.formula.operation {
		case '+':
			// target = left + right => left = target - right
			newTarget = target - rightValue
		case '-':
			// target = left - right => left = target + right
			newTarget = target + rightValue
		case '*':
			// target = left * right => left = target / right
			newTarget = target / rightValue
		case '/':
			// target = left / right => left = target * right
			newTarget = target * rightValue
		}
		nextName = monkey.formula.left
	} else {
		leftValue := evaluate(monkey.formula.left, monkeys)
		switch monkey.formula.operation {
		case '+':
			// target = left + right => right = target - left
			newTarget = target - leftValue
		case '-':
			// target = left - right => right = left - target
			newTarget = leftValue - target
		case '*':
			// target = left * right => right = target / left
			newTarget = target / leftValue
		case '/':
			// target = left / right => right = left / target
			newTarget = leftValue / target
		}
		nextName = monkey.formula.right
	}

	return solveForHumn(nextName, newTarget, monkeys)
}

type Day21Monkey struct {
	hasValue bool
	value    int
	formula  Formula
}

type Formula struct {
	left, right string
	operation   byte
}
