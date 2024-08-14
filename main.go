package main

import (
	"fmt"
	"prime/constants"
	"prime/finder"
	"time"
)

func main() {
	var min int = 1
	var max int = constants.Giga
	fmt.Printf("Finding prime numbers from %d to %d\n", min, max)

	startTime := time.Now()

	primeFinder := finder.PrimeFinder{}
	primeFinder.SetRange(min, max, 1000)
	finder.WritePrimeFinderRangesToFile(primeFinder.Ranges)

	primeNumbers := primeFinder.Execute()
	finder.WritePrimesToFile(primeNumbers)

	endTime := time.Now()

	numberPerSecond := float64(len(primeNumbers)) / endTime.Sub(startTime).Seconds()
	fmt.Printf("Found prime numbers from %d to %d\n", min, max)
	fmt.Printf("Execution time: %v\n", endTime.Sub(startTime))
	fmt.Printf("Number of prime numbers: %d\n", len(primeNumbers))
	fmt.Printf("Number found per second: %.2f\n", numberPerSecond)
}
