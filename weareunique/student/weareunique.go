package student

// I guess we misunderstood in the end
/* func WeAreUnique(str1, str2 string) int {
	if str1 == "" && str2 == "" {
		return -1
	}

	uniques := 0
	comb := str1 + str2

	for _, r1 := range comb {
		counter := 0

		for _, r2 := range comb {
			if r1 == r2 {
				counter++
			}
		}
		if counter == 1 {
			uniques++
		}
	}
	return uniques
} */

func WeAreUnique(str1, str2 string) int {
	if str1 == "" && str2 == "" {
		return -1
	}

	comb := str1 + str2
	uniques := ""

	for _, r0 := range comb {
		onStr1 := false
		onStr2 := false
		onUniques := false

		for _, r1 := range str1 {
			if r0 == r1 {
				onStr1 = true
			}
		}

		for _, r2 := range str2 {
			if r0 == r2 {
				onStr2 = true
			}
		}

		for _, rU := range uniques {
			if r0 == rU {
				onUniques = true
			}
		}

		// If not already found, and not on the other string (it's going to be on one of them)
		if !onUniques && (!onStr1 || !onStr2) {
			uniques += string(r0)
		}
	}

	return len(uniques)
}
