package adventofcode2022

// In absence of a generic Reverse() in Go's stdlib...
// https://github.com/golang/go/issues/36887
// https://github.com/golang/go/issues/47988
func Reverse[T any](s []T) {
        for a1, a2 := 0, len(s)-1; a1 < a2; a1, a2 = a1+1, a2-1 {
                s[a1], s[a2] = s[a2], s[a1]
        }
}
