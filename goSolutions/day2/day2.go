package main

import (
	"fmt"
	"github.com/bsl333/AoC/goSolutions/utils"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	p := utils.GetFilepath("2")
	contents, err := ioutil.ReadFile(p)
	check(err)
	commands := strSliceToIntSlice(strings.Split(string(contents), ","))
	cp := make([]int, len(commands))

	for i := 0; i < 99; i++ {
		for j := 0; j < 99; j++ {
			copy(cp, commands)
			result := runCommands(i, j, cp)
			if i == 12 && j == 2 {
				fmt.Printf("Solution 1: %v\n\n", result)
			}
			if result == 19690720 {
				fmt.Printf("Solution 2: %v\n", i*100+j)
				// return
			}
		}
	}
	t := time.Now()
	fmt.Println("TIME", t.Sub(start))
	fmt.Println(time.Since(start))
}

func runCommands(c1 int, c2 int, commands []int) int {
	commands[1] = c1
	commands[2] = c2
	for i := 0; i < len(commands); i += 4 {
		op := commands[i]
		idx1 := commands[i+1]
		idx2 := commands[i+2]
		location := commands[i+3]
		if op == 1 {
			commands[location] = commands[idx1] + commands[idx2]
		} else if op == 2 {
			commands[location] = commands[idx1] * commands[idx2]
		} else if op == 99 {
			break
		}
	}
	return commands[0]
}

func strSliceToIntSlice(s []string) []int {
	intSlice := make([]int, 0, len(s))
	for _, str := range s {
		val, err := strconv.Atoi(str)
		if err != nil {
			continue
		}
		check(err)
		intSlice = append(intSlice, val)
	}

	return intSlice
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
