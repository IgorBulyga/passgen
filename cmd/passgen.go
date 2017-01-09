package main

import (
	"flag"
	"fmt"
	"os"
	// "strings"
	"github.com/IgorBulyga/passgen"
)

func checkLength(l int) {
	if l <= 0 {
		fmt.Printf("Password length must be greater than 0")
		os.Exit(1)
	}
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

	if flag.NArg() == 0 {
		// flag.Usage()
		// os.Exit(1)
	} else {

	}

	checkLength(l)
	fmt.Println("length = ", l)
	l = 10
	excludedLetters = "abcdefghijklmopqrstvuwxyz"
	var e = []rune(excludedLetters)

	fmt.Println(generator.Generate(l, e, []string{generator.LowerCase}))
}
