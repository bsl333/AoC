package main

import "fmt"

const (
	min = 256310
	max = 732736
)

func main() {
	validPasswordCount := 0
	for i := min + 1; i < max; i++ {
		if validInput(i) {
			validPasswordCount++
		}
	}

	fmt.Println(validPasswordCount)

}

func validInput(num int) bool {
	revNum := intToSliceRev(num)
	hasAdjency := false
	for i := 0; i < len(revNum)-1; i++ {
		if revNum[i] < revNum[i+1] {
			return false
		}
		if revNum[i] == revNum[i+1] {
			hasAdjency = true
		}
	}
	return hasAdjency
}

func intToSliceRev(i int) []int {
	revNum := make([]int, 0, 6)
	for i > 0 {
		revNum = append(revNum, i%10)
		i /= 10
	}
	return revNum
}
