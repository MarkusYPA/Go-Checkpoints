package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(ZipString("YouuungFellllas"))
	fmt.Println(ZipString("Thee quuick browwn fox juumps over the laaazy dog"))
	fmt.Println(ZipString("Helloo Therre!"))
}

func ZipString(s string) string {

	result := ""

	for i := 0; i < len(s); i++ {
		char := s[i]
		count := 1
		for i != len(s)-1 && char == s[i+count] {
			count++
		}
		result += strconv.Itoa(count) + string(char)
		i += count - 1
	}
	return result
}
