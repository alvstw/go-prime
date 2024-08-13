package main

import (
	"fmt"
	"time"
)

func main() {
	var min int = 1
	var max int = giga
	fmt.Printf("Finding prime numbers from %d to %d\n", min, max)

	startTime := time.Now()

	finder := PrimeFinder{}
	finder.setRange(min, max, 1000)
	writePrimeFinderRangesToFile(finder.Ranges)

	primeNumbers := finder.execute()
	writePrimesToFile(primeNumbers)

	endTime := time.Now()

	numberPerSecond := float64(len(primeNumbers)) / endTime.Sub(startTime).Seconds()
	fmt.Printf("Found prime numbers from %d to %d\n", min, max)
	fmt.Printf("Execution time: %v\n", endTime.Sub(startTime))
	fmt.Printf("Number of prime numbers: %d\n", len(primeNumbers))
	fmt.Printf("Number found per second: %.2f\n", numberPerSecond)
}
