package main

import "fmt"

func main() {
	input1 := []uint{2, 3, 1, 1, 0}
	fmt.Println(CanJump(input1))

	input2 := []uint{3, 2, 1, 0, 4}
	fmt.Println(CanJump(input2))

	input3 := []uint{0}
	fmt.Println(CanJump(input3))

	/*
		 	tests := []struct {
				args []uint
				want bool
			}{
				{args: []uint{2, 3, 1, 1, 4}, want: true},
				{args: []uint{1, 1, 1, 1, 0}, want: true},
				{args: []uint{5, 4, 3, 2, 1, 0}, want: true},
				{args: []uint{0}, want: true},
				{args: []uint{5}, want: true},
				{args: []uint{}, want: false},
				{args: []uint{1, 2, 3, 0, 2}, want: false},
				{args: []uint{3, 2, 1, 0, 4}, want: false},
				{args: []uint{0, 0, 0, 0, 0}, want: false},
				{args: []uint{1, 2, 3}, want: false},
				{args: []uint{1, 2, 3, 0, 1}, want: false},
				{args: []uint{1, 0, 0, 0, 0}, want: false},
				{args: []uint{1, 0, 1, 0, 1}, want: false},
				{args: []uint{10, 20, 30, 40, 0}, want: false},
			}

			for _, tc := range tests {
				got := CanJump(tc.args)
				if !reflect.DeepEqual(got, tc.want) {
					log.Fatalf("%s(%+v) == %+v instead of %+v\n",
						"CanJump",
						tc.args,
						got,
						tc.want,
					)
				}
			}
	*/
}

func CanJump(nums []uint) bool {

	index := 0
	prevIndex := -1

	if len(nums) == 1 {
		return true
	}

	if len(nums) == 0 {
		return false
	}

	for index < len(nums) {
		index += int(nums[index])

		if index >= len(nums) {
			return false
		}

		if index == prevIndex {
			return false
		}

		if index == len(nums)-1 {
			return true
		}

		prevIndex = index
	}

	return true
}
