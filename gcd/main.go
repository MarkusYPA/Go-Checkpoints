package main

import (
	"fmt"
	"os"
)

func main() {
	testCases := []struct {
		a    uint
		b    uint
		want uint
	}{
		{42, 10, 2},
		{42, 12, 6},
		{14, 77, 7},
		{17, 3, 1},
		{12, 23, 1},
		{25, 15, 5},
		{23043, 122, 1},
		{11, 77, 11},
	}

	for _, tc := range testCases {
		got := Gcd(tc.a, tc.b)
		if got != tc.want {
			fmt.Printf("Gcd(%d, %d) = %d instead of %d\n", tc.a, tc.b, got, tc.want)
			os.Exit(1)
		}
	}
}

/* func Gcd(a, b uint) uint {
var guess uint

	   	if a > b {
	   		guess = b
	   	} else {
	   		guess = a
	   	}

	for guess > 1 {
		if a%guess == 0 && b%guess == 0 {
			break
		}
		guess--
	}

	return guess
} */

func Gcd(a, b uint) uint {
	if b == 0 {
		return a
	}
	return Gcd(b, a%b)
}
