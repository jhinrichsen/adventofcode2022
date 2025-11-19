package adventofcode2022

import (
	"testing"
)

func TestDay20Part1Example(t *testing.T) {
	const want = 3
	ns, err := numbersFromFilename(exampleFilename(20))
	if err != nil {
		t.Fatal(err)
	}
	got := Day20(ns, 1, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay20Part1(t *testing.T) {
	const want = 11616
	ns, err := numbersFromFilename(filename(20))
	if err != nil {
		t.Fatal(err)
	}
	got := Day20(ns, 1, true)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay20Part2ExampleRounds(t *testing.T) {
	ns, err := numbersFromFilename(exampleFilename(20))
	if err != nil {
		t.Fatal(err)
	}

	// Test each round of mixing
	tests := []struct {
		rounds int
		want   []int // Expected arrangement starting from 0
	}{
		{0, []int{0, 3246356612, 811589153, 1623178306, -2434767459, 2434767459, -1623178306}},  // Initial (after key applied, not mixed)
		{1, []int{0, -2434767459, 3246356612, -1623178306, 2434767459, 1623178306, 811589153}},  // After 1 round
		{2, []int{0, 2434767459, 1623178306, 3246356612, -2434767459, -1623178306, 811589153}},  // After 2 rounds
		{3, []int{0, 811589153, 2434767459, 3246356612, 1623178306, -1623178306, -2434767459}},  // After 3 rounds
		{4, []int{0, 1623178306, -2434767459, 811589153, 2434767459, 3246356612, -1623178306}},  // After 4 rounds
		{5, []int{0, 811589153, -1623178306, 1623178306, -2434767459, 3246356612, 2434767459}},  // After 5 rounds
		{6, []int{0, 811589153, -1623178306, 3246356612, -2434767459, 1623178306, 2434767459}},  // After 6 rounds
		{7, []int{0, -2434767459, 2434767459, 1623178306, -1623178306, 811589153, 3246356612}},  // After 7 rounds
		{8, []int{0, 1623178306, 3246356612, 811589153, -2434767459, 2434767459, -1623178306}},  // After 8 rounds
		{9, []int{0, 811589153, 1623178306, -2434767459, 3246356612, 2434767459, -1623178306}},  // After 9 rounds
		{10, []int{0, -2434767459, 1623178306, 3246356612, -1623178306, 2434767459, 811589153}}, // After 10 rounds
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := mixAndGetArrangement(ns, tt.rounds, false)
			if len(got) != len(tt.want) {
				t.Fatalf("round %d: length mismatch: got %d, want %d", tt.rounds, len(got), len(tt.want))
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Fatalf("round %d: position %d: got %v, want %v\nfull got:  %v\nfull want: %v",
						tt.rounds, i, got[i], tt.want[i], got, tt.want)
				}
			}
		})
	}
}

func TestDay20Part2Example(t *testing.T) {
	const want = 1623178306
	ns, err := numbersFromFilename(exampleFilename(20))
	if err != nil {
		t.Fatal(err)
	}
	got := Day20(ns, 10, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay20Part2(t *testing.T) {
	const want = 9937909178485
	ns, err := numbersFromFilename(filename(20))
	if err != nil {
		t.Fatal(err)
	}
	got := Day20(ns, 10, false)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

// Helper function to get the arrangement after mixing
func mixAndGetArrangement(srcs []int, mixRounds int, part1 bool) []int {
	const decryptionKey = 811589153
	values := make([]int, len(srcs))
	for i := range srcs {
		if part1 {
			values[i] = srcs[i]
		} else {
			values[i] = srcs[i] * decryptionKey
		}
	}

	list := make([]Node, len(values))
	for i := range values {
		list[i] = Node{Value: values[i], OriginalIndex: i}
	}

	// Mix the specified number of times
	for range mixRounds {
		for origIdx := range values {
			currIdx := -1
			for i := range list {
				if list[i].OriginalIndex == origIdx {
					currIdx = i
					break
				}
			}

			node := list[currIdx]
			list = append(list[:currIdx], list[currIdx+1:]...)

			newIdx := currIdx + node.Value
			if len(list) > 0 {
				newIdx = ((newIdx % len(list)) + len(list)) % len(list)
			}

			list = append(list[:newIdx], append([]Node{node}, list[newIdx:]...)...)
		}
	}

	// Find index of 0 and return arrangement starting from 0
	zeroIdx := -1
	for i := range list {
		if list[i].Value == 0 {
			zeroIdx = i
			break
		}
	}

	result := make([]int, len(list))
	for i := range list {
		result[i] = list[(zeroIdx+i)%len(list)].Value
	}
	return result
}

func BenchmarkDay20Part1(b *testing.B) {
	ns, err := numbersFromFilename(filename(20))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Day20(ns, 1, true)
	}
}

func BenchmarkDay20Part2(b *testing.B) {
	ns, err := numbersFromFilename(filename(20))
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Day20(ns, 10, false)
	}
}
