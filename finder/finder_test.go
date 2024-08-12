package finder

import (
	"fmt"
	"prime/constants"
	"testing"
	"time"
)

func BenchmarkPrimeFinder(b *testing.B) {
	startTime := time.Now()

	finder := PrimeFinder{}

	finder.SetRange(1, constants.Giga, 100)
	WritePrimeFinderRangesToFile(constants.PrimeRangesFilePath, finder.Ranges)

	primeNumbers := finder.Execute()
	WritePrimesToFile(constants.PrimeFilePath, primeNumbers)

	endTime := time.Now()

	numberPerSecond := float64(len(primeNumbers)) / endTime.Sub(startTime).Seconds()
	fmt.Printf("Execution time: %v\n", endTime.Sub(startTime))
	fmt.Printf("Number of prime numbers: %d\n", len(primeNumbers))
	fmt.Printf("Number found per second: %.2f\n", numberPerSecond)
}
