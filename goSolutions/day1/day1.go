package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("/Users/bsl333/Documents/personal-code/Udemy/golang/goworkspace/src/github.com/bsl333/AoC/data/day1.txt")
	check(err)
	defer file.Close()

	fuelNeededForModules, fuelNeededForFuel := part1(file)
	fmt.Printf("Total fuel for part 1 %v\n\n", fuelNeededForModules)
	fmt.Printf("Total fuel for part 2 %v", fuelNeededForModules+fuelNeededForFuel)
}

func part1(f *os.File) (int, int) {
	scanner := bufio.NewScanner(f)
	fuelNeededForModules, fuelNeededForFuel := 0, 0

	for scanner.Scan() {
		mass, err := strconv.Atoi(scanner.Text())
		check(err)
		fuelForModule := calcFuelNeeded(mass)
		fuelNeededForModules += fuelForModule
		fuelNeededForFuel += calcAdditionalFuelForFuelWeight(fuelForModule)
	}

	if err := scanner.Err(); err != nil {
		check(err)
	}

	return fuelNeededForModules, fuelNeededForFuel
}

func calcFuelNeeded(mass int) int {
	return mass/3 - 2
}

func calcAdditionalFuelForFuelWeight(mass int) int {
	if mass > 0 {
		additFuel := calcFuelNeeded(mass)
		return additFuel + calcAdditionalFuelForFuelWeight(additFuel)
	}
	return 0
}

func check(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}
