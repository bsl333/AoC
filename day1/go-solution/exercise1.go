package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./input.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalFuel := 0

	for scanner.Scan() {
		mass, err := strconv.Atoi(scanner.Text())
		check(err)
		totalFuel += calcFuelNeeded(mass)
	}
	fmt.Println(totalFuel)
}

func calcFuelNeeded(mass int) int {
	return mass/3 - 2
}

func check(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}
