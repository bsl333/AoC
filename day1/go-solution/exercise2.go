package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	f, e := os.Open("./input.txt")
	check(e)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var totalFuel int

	for scanner.Scan() {
		mass, e := strconv.Atoi(scanner.Text())
		check(e)
		totalFuel += fuelNeeded(mass)
	}

	fmt.Println(totalFuel)
}

func fuelNeeded(mass int) int {
	fuel := mass/3 - 2
	if fuel > 0 {
		return fuel + fuelNeeded(fuel)
	}
	return 0
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
