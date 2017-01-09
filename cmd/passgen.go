package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/IgorBulyga/passgen"
)

func checkLength(l int) {
	if l <= 0 {
		fmt.Printf("Password length must be greater than 0")
		os.Exit(1)
	}
}

func configureOptions(o string) []string {
	var os = strings.Split(o, ",")
	if len(os) > 3 || len(os) == 0 {
		fmt.Println("Invalid number of Options, used default (U, L, N)")
		return []string{generator.LowerCase, generator.UpperCase, generator.Number}
	}

	var om = make(map[string]string)
	for _, v := range os {
		v = strings.Trim(v, " ")
		switch v {
		case "U":
			om[v] = generator.UpperCase
		case "L":
			om[v] = generator.LowerCase
		case "N":
			om[v] = generator.Number
		}
	}

	var r = make([]string, len(om))
	for _, v := range om {
		r = append(r, v)
	}
	return r
}

func main() {
	flag.Usage = func() {
		fmt.Printf("Usage of passgen - generate random password\n")
		flag.PrintDefaults()
	}

	var l int
	flag.IntVar(&l, "l", 1, "Password length")
	var excludedLetters string
	flag.StringVar(&excludedLetters, "e", "", "Excluded symbols in one line, eg: aA12")
	var options string
	flag.StringVar(&options, "o", "U,L,N", "Options of which symbols types will be in password.\n\tPossible values:\n\t\tU - Uppercase letters\n\t\tL - Lowercase letters\n\t\tN - numbers")

	flag.Parse()

	checkLength(l)
	var e = []rune(excludedLetters)
	var o = configureOptions(options)
	fmt.Println(generator.Generate(l, e, o))
	os.Exit(0)
}
