package main

import (
	"fmt"
	"time"
)

const (
	min = 256310
	max = 732736
)

func main() {
	start := time.Now()
	validPasswordCount := 0
	validPasswordCountAdv := 0
	for i := min; i <= max; i++ {
		if validInputPart1(i) {
			validPasswordCount++
		}
		if validInputPart2(i) {
			validPasswordCountAdv++
		}
	}

	fmt.Println("part 1:", validPasswordCount)
	fmt.Println("part 2:", validPasswordCountAdv)
	fmt.Println(time.Since(start))

}

func validInputPart1(num int) bool {
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

func validInputPart2(num int) bool {
	revNum := intToSliceRev(num)
	var prev int
	count := 0
	hasAdjency := false
	for i, n := range revNum {
		if i == 0 {
			prev = n
			count = 1
			continue
		}
		if n > prev {
			return false
		} else if prev == n {
			count++
		} else {
			if count == 2 {
				hasAdjency = true
			}
			prev = n
			count = 1
		}
	}
	if count == 2 {
		hasAdjency = true
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
