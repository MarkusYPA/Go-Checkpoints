package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		z01.PrintRune('\n')
		return
	}

	out := ""
	space := false
	foundLetter := false
	for _, r := range args[0] {

		if !foundLetter && (r == ' ' || r == '\t') {
			continue
		}

		if !(r == ' ' || r == '\t') {
			out += string(r)
			space = false
			foundLetter = true
		}

		if (r == ' ' || r == '\t') && !space {
			out += " "
			space = true
		}

	}

	if out[len(out)-1] == ' ' {
		out = out[:len(out)-1]
	}

	prtStr(out)
}

func prtStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
