package adventofcode2022

import "fmt"

func Day25(lines []string) Snafu {
	sum := Snafu("0")
	for _, line := range lines {
		sum = AddSnafu(sum, Snafu(line))
	}
	return sum
}

// day25ImplViaDec convertes each snafu -> dec, adds all, and returns dec ->
// snafu.
/*
func day25ImplViaDec(lines []string) Snafu {
	var sum int
	for _, line := range lines {
		sum += SnafuToDec(Snafu(line))
	}
	return DecToSnafu(sum)
}
*/

type Snafu string

/*
func DecToSnafu(n int) Snafu {
}
*/

func SnafuToDec(s Snafu) int {
	var n int
	digit := 1
	for i := len(s) - 1; i >= 0; i-- {
		var val int
		switch s[i] {
		case '0':
			val = 0
		case '1':
			val = 1
		case '2':
			val = 2
		case '=':
			val = -2
		case '-':
			val = -1

		default:
			panic(fmt.Sprintf("illegal snafu character '%c'", s[i]))
		}
		n += digit * val
		digit *= 5
	}
	return n
}

// AddSnafu adds two snafu numbers in the snafu domain, i.e. without conversion
// to decimal and then using the decimal '+' operator.
func AddSnafu(a, b Snafu) Snafu {
	var (
		min, max = len(a), len(a)
		larger   = a
	)
	if len(a) < len(b) {
		max = len(b)
		larger = b
	} else if len(a) > len(b) {
		min = len(b)
	}
	digit := func(s Snafu, n int) byte { // return nth digit from right
		return s[len(s)-1-n]
	}
	var buf = make([]byte, max+1) // +1 for optional carry
	var carry byte = 48           // '0'
	for i := 0; i < min; i++ {
		var c1, c2 byte
		n := carry
		da := digit(a, i)
		c1, n = AddSnafuDigit(n, da)
		db := digit(b, i)
		c2, n = AddSnafuDigit(n, db)
		_, carry = AddSnafuDigit(c1, c2)
		buf[i] = n
	}

	// copy everything from the larger number, considering carry
	for i := min; i < max; i++ {
		var n byte
		carry, n = AddSnafuDigit(carry, digit(larger, i))
		buf[i] = n
	}
	buf[max] = carry // final carry may extend original len
	Reverse(buf)
	if buf[0] == '0' {
		buf = buf[1:]
	}
	return Snafu(buf)
}

// AddSnafu1 adds two snafu digits and returns carry ('0' or '1') and value.
func AddSnafuDigit(a, b byte) (byte, byte) {
	switch a {
	case '2':
		switch b {
		case '2':
			return '1', '-'
		case '1':
			return '1', '='
		case '0':
			return '0', '2'
		case '-':
			return '0', '1'
		case '=':
			return '0', '0'
		}
	case '1':
		switch b {
		case '2':
			return '1', '='
		case '1':
			return '0', '2'
		case '0':
			return '0', '1'
		case '-':
			return '0', '0'
		case '=':
			return '0', '-'
		}
	case '0':
		return '0', b
	case '-':
		switch b {
		case '2':
			return '0', '1'
		case '1':
			return '0', '0'
		case '0':
			return '0', '-'
		case '-':
			return '0', '='
		case '=':
			return '-', '2'
		}
	case '=':
		switch b {
		case '2':
			return '0', '0'
		case '1':
			return '0', '-'
		case '0':
			return '0', '='
		case '-':
			return '-', '2'
		case '=':
			return '-', '1'
		}
	}
	panic(fmt.Sprintf("illegal snafu %v or %v", a, b))
}
