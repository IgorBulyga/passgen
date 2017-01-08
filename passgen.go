package passgen

import (
	"math/rand"
	"time"
)

// Format for generate password
const (
	Number    = "numbers"
	LowerCase = "lower"
	UpperCase = "upper"
)

// LatinLetters - uppercased latin letters
var LatinLetters = [26]rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}

// Generate Password generation functions\n\n
// l Length of password
// e Excluded symbols
// f Password format: Numbers, LowerCase, UpperCase
func Generate(l int, e []rune, f []string) string {
	var s = availableSymbols(f)
	s = excludeElements(s, e)

	rs := rand.NewSource(time.Now().UnixNano())
	rg := rand.New(rs)

	var password string
	for i := 0; i < l; i++ {
		rn := rg.Intn(len(s))
		password += string(s[rn])
	}
	return password
}

func availableSymbols(f []string) []rune {
	var s = []rune{}
	for _, v := range f {
		switch v {
		case Number:
			s = append(s, []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}...)
		case LowerCase:
			for _, v := range LatinLetters {
				l := v + 32
				s = append(s, l)
			}
		case UpperCase:
			s = append(s, LatinLetters[:]...)
		}
	}
	return s
}

func excludeElements(s []rune, e []rune) []rune {
	for i, v := range s {
		if containsr(e, v) {
			s = s[:i+copy(s[i:], s[i+1:])]
		}
	}
	return s
}

func contains(s []string, e string) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

func containsr(s []rune, e rune) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}
