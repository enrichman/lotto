package main

import (
	"fmt"
	"math/rand"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().Unix()))

func main() {
	wheels := []string{
		"Bari", "Cagliari", "Firenze", "Genova", "Milano",
		"Napoli", "Palermo", "Roma", "Torino", "Venezia", "Nazionale",
	}

	for _, w := range wheels {
		numbers := getUniqueRandomNumbers(1, 90, 6)
		fmt.Printf("%-8s\t%+v\n", w, numbers)
	}
}

func getUniqueRandomNumbers(lowerBound, upperBound, numbers int) []int {
	vals := map[int]struct{}{}
	for len(vals) < 6 {
		vals[lowerBound+r.Intn(upperBound)] = struct{}{}
	}

	values := []int{}
	for k := range vals {
		values = append(values, k)
	}
	return values
}
