package adventofcode2022

import (
	"fmt"
	"go/constant"
	"go/token"
	"go/types"
	"math"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	ID          int
	Items       []float64
	Operation   string
	DivisibleBy float64
	IfTrue      int
	IfFalse     int
}

func NewMonkey(lines []string) (Monkey, error) {
	var m Monkey
	var n int
	var err error

	// ID
	m.ID, err = strconv.Atoi(lines[0][len("Monkey ") : len(lines[0])-1])
	if err != nil {
		return m, fmt.Errorf("error parsing %s: %w", lines[0], err)
	}

	// Items
	items := strings.Split(lines[1], ": ")[1]
	ss := strings.Split(strings.TrimSpace(items), ", ")
	for i := range ss {
		n, err := strconv.Atoi(ss[i])
		if err != nil {
			return m, fmt.Errorf("error parsing monkey %d: items "+
				"%q: %w", m.ID, ss[i], err)
		}
		m.Items = append(m.Items, float64(n))
	}

	// Operation, only keep statement without assignment
	m.Operation = strings.TrimSpace(lines[2][len("  Operation: old = "):])

	// Divisible
	ss = strings.Fields(lines[3])
	n, err = strconv.Atoi(ss[3])
	if err != nil {
		return m, fmt.Errorf("error parsing monkey %d: divisible %q: "+
			"%w", m.ID, ss[3], err)
	}
	m.DivisibleBy = float64(n)

	// If true
	ss = strings.Fields(lines[4])
	n, err = strconv.Atoi(ss[5])
	if err != nil {
		return m, fmt.Errorf("error parsing monkey %d: true %q: %w",
			m.ID, ss[5], err)
	}
	m.IfTrue = n

	// If false
	ss = strings.Fields(lines[5])
	n, err = strconv.Atoi(ss[5])
	if err != nil {
		return m, fmt.Errorf("error parsing monkey %d: true %q: %w",
			m.ID, ss[5], err)
	}
	m.IfFalse = n

	return m, nil
}

func NewMonkeys(lines []string) ([]Monkey, error) {
	const oneMonkey = 7 // one monkey every 7 lines
	var ms []Monkey
	for i := 0; i < len(lines); i += oneMonkey {
		m, err := NewMonkey(lines[i : i+oneMonkey])
		if err != nil {
			return ms, fmt.Errorf("error parsing line %d: %w",
				i, err)
		}
		ms = append(ms, m)
	}
	return ms, nil
}

func Day11(lines []string, part1 bool) uint {
	monkeys, err := NewMonkeys(lines)
	if err != nil {
		return 0
	}

	divisible := func(x, y float64) bool {
		// return x%n == 0
		// return math.Mod(x, y) == 0 // not working
		// https://github.com/golang/go/issues/26181
		div := x / y
		b := math.Floor(div) == div
		return b
	}

	move := func(from, into int) {
		monkeys[into].Items = append(monkeys[into].Items,
			monkeys[from].Items[0])
		monkeys[from].Items = monkeys[from].Items[1:]
	}

	rounds := 10_000
	if part1 {
		rounds = 20
	}
	inspections := make([]uint, len(monkeys))
	for range rounds {
		for j := range monkeys {
			for len(monkeys[j].Items) > 0 {
				inspections[j]++
				// apply operation
				m := map[string]float64{"old": monkeys[j].Items[0]}
				monkeys[j].Items[0] = Eval(monkeys[j].Operation, m)

				if part1 {
					monkeys[j].Items[0] = math.Floor(monkeys[j].Items[0] / 3)
				}

				b := divisible(monkeys[j].Items[0],
					monkeys[j].DivisibleBy)
				nextMonkey := monkeys[j].IfTrue
				if !b {
					nextMonkey = monkeys[j].IfFalse
				}
				move(j, nextMonkey)
			}
		}
	}

	sort.Slice(inspections, func(i, j int) bool {
		return inspections[i] > inspections[j]
	})
	return inspections[0] * inspections[1]
}

// Eval uses Go's internal compiler to evaluate a formula.
func Eval(formula string, m map[string]float64) float64 {
	// redeclare global scope
	types.Universe = types.NewScope(nil, token.NoPos, token.NoPos, "universe")

	for k, v := range m {
		c := types.NewConst(
			token.NoPos,
			nil,
			k,
			types.Typ[types.Float64],
			constant.MakeFloat64(float64(v)))
		types.Universe.Insert(c)
	}

	fs := token.NewFileSet()
	tv, err := types.Eval(fs, nil, token.NoPos, formula)
	if err != nil {
		return 0
	}
	n, exact := constant.Float64Val(tv.Value)
	if !exact {
		return 0
	}
	return n
}
