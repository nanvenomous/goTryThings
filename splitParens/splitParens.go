package splitParens

import "math"

type symTracker map[rune]int

func newSymTracker() symTracker {
	st := make(symTracker)
	st['['] = 0
	st[']'] = 0
	st['('] = 0
	st[')'] = 0
	st['?'] = 0
	return st
}

// checkParenCombo takes a sym tracker and returns the remaining number of check marks
func checkParenCombo(st symTracker, cmarks int, lp rune, rp rune) (bool, int) {
	lpn := st[lp]
	rpn := st[rp]
	diff := int(math.Abs(float64(lpn - rpn)))
	diff = cmarks - diff
	if diff < 0 { // ran out of check marks
		return false, -1
	}
	return true, diff
}

func sidesMatches(ra []rune) bool {
	st := newSymTracker()
	for _, c := range ra { // O(s)
		st[c] += 1
	}
	cmarks := st['?']
	success, cmarks := checkParenCombo(st, cmarks, '(', ')')
	if !success {
		return false
	}
	success, cmarks = checkParenCombo(st, cmarks, '[', ']')
	if !success || cmarks > 0 {
		return false
	}
	return true
}

func splitRuneArray(ra []rune, idx int) ([]rune, []rune) {
	return ra[0:idx], ra[idx:]
}

func matchParens(s string) int {
	matchCount := 0
	ra := []rune(s)                  // O(S)
	for i := 2; i < len(s); i += 2 { // O(S/2)
		ra1, ra2 := splitRuneArray(ra, i)
		if sidesMatches(ra1) && sidesMatches(ra2) { // O(s1 + s2) = O(S)
			matchCount++
		}
	}
	return matchCount
}

func numberOfSplits(s string) int {
	if len(s)%2 != 0 { // if odd there will be no possible splits
		return 0
	}
	return matchParens(s)
}

/*
S: lenght of string given
s[n]: is the length of a substring from string given
Time Complexity:
O(S + (S/2)*(s1 + s2))
O(S + (S/2)*S)
O(S + S^2) or O(S*(1 + S))
*/
