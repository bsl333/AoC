package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/bsl333/AoC/goSolutions/utils"
)

func main() {
	start := time.Now()
	p := utils.GetFilepath("3")
	contents, err := ioutil.ReadFile(p)
	check(err)
	data := strings.Split(string(contents), "\n")
	d1 := strings.Split(data[0], ",")
	d2 := strings.Split(data[1], ",")

	part1(d1, d2)
	part2(d1, d2)
	fmt.Println(time.Since(start))
}

func part1(d1, d2 []string) {
	wirePath1 := drawPathCountSteps(d1)
	wirePath2 := drawPathCountSteps(d2)

	minDistance := math.MaxFloat64
	for key := range wirePath1 {
		if steps := wirePath2[key]; steps > 0 {
			curDist := math.Abs(float64(key[0])) + math.Abs(float64(key[1]))
			if curDist < minDistance {
				minDistance = curDist
			}
		}
	}

	fmt.Println(minDistance)
}

func part2(d1, d2 []string) {
	wirePath1 := drawPathCountSteps(d1)
	wirePath2 := drawPathCountSteps(d2)

	minSteps := math.MaxInt64
	for coordinates, steps1 := range wirePath1 {
		if steps2 := wirePath2[coordinates]; steps2 > 0 {
			numSteps := steps1 + steps2
			if numSteps < minSteps {
				minSteps = numSteps
			}
		}
	}

	fmt.Println(minSteps)
}

// iteration determines direcition
// prev is the previous cooridinates (x, y)
// eg.,
// if positive, increment prev position at idx by 1
// if negative, decrement prev position at idx by 1
func move(path map[[2]int]int, prev [2]int, idx, count, iterations int) ([2]int, int) {
	if iterations == 0 {
		return prev, count
	}
	count++
	if iterations > 0 {
		prev[idx]++
		path[prev] = count
		return move(path, prev, idx, count, iterations-1)
	}
	prev[idx]--
	path[prev] = count
	return move(path, prev, idx, count, iterations+1)
}

func drawPathCountSteps(directions []string) map[[2]int]int {
	prev := [2]int{0, 0}
	path := map[[2]int]int{{0, 0}: 0}
	count := 0
	for _, direction := range directions {
		iterations, err := strconv.Atoi(string(direction[1:]))
		check(err)
		switch direction[0] {
		case 'R':
			prev, count = move(path, prev, 0, count, iterations)
		case 'U':
			prev, count = move(path, prev, 1, count, iterations)
		case 'L':
			prev, count = move(path, prev, 0, count, -1*iterations)
		case 'D':
			prev, count = move(path, prev, 1, count, -1*iterations)
		}
	}
	return path
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// func drawPath(directions []string) map[[2]int]bool {
// 	prev := [2]int{0, 0}
// 	path := map[[2]int]bool{{0, 0}: true}
// 	for _, direction := range directions {
// 		n, err := strconv.Atoi(string(direction[1:]))
// 		check(err)
// 		if direction[0] == 'R' {
// 			for i := 0; i < n; i++ {
// 				prev[0]++
// 				path[prev] = true
// 			}
// 		} else if direction[0] == 'U' {
// 			for i := 0; i < n; i++ {
// 				prev[1]++
// 				path[prev] = true
// 			}
// 		} else if direction[0] == 'L' {
// 			for i := 0; i < n; i++ {
// 				prev[0]--
// 				path[prev] = true
// 			}
// 		} else {
// 			for i := 0; i < n; i++ {
// 				prev[1]--
// 				path[prev] = true
// 			}
// 		}
// 	}

// 	return path
// }
