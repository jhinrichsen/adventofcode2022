package adventofcode2022

import (
	"strconv"
	"strings"
)

func Day21(lines []string, part1 bool) int {
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

type Formula struct {
	left, right string
	operation   byte
}
