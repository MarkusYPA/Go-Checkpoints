package main

import (
	"unicode"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"

	"github.com/01-edu/z01"
)

func main() {
	/* 	bytes := [10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'}
	   	printMemory(bytes) */

	table := [10]byte{}

	for j := 0; j < 5; j++ {
		for i := 0; i < random.IntBetween(7, 10); i++ {
			table[i] = byte(random.IntBetween(13, 126))
		}
		challenge.Function("PrintMemory", PrintMemory, solutions.PrintMemory, table)
	}
	table2 := [10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'}
	challenge.Function("PrintMemory", PrintMemory, solutions.PrintMemory, table2)

}

func PrintMemory(bts [10]byte) {

	for i, b := range bts {
		str := toHex(b)

		if str == "" {
			z01.PrintRune('0')
			z01.PrintRune('0')
		}

		if len(str) == 1 {
			z01.PrintRune('0')
		}

		for _, r := range str {
			z01.PrintRune(r)
		}

		// This does what all the above does, but may not be allowed
		//fmt.Printf("%.2x", b)

		// Only print space or line feed if not at the last byte
		if i != len(bts)-1 {
			if (i+1)%4 != 0 {
				z01.PrintRune(' ')
			} else {
				z01.PrintRune('\n')
			}
		}
	}

	z01.PrintRune('\n')

	for _, b := range bts {
		if unicode.IsGraphic(rune(b)) && rune(b) != ' ' {
			z01.PrintRune(rune(b))
		} else {
			z01.PrintRune('.')
		}
	}

	z01.PrintRune('\n')
}

func toHex(b byte) string {
	n := int(b)
	base := "0123456789abcdef"
	out := ""

	for n > 0 {
		charIn := n % 16
		char := base[charIn]
		out = string(char) + out
		n /= 16
	}

	return out
}
