package finder

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func WritePrimesToFile(primeNumbers []int) {
	file, err := os.OpenFile("primes.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	for _, val := range primeNumbers {
		_, err := writer.WriteString(fmt.Sprintf("%d\n", val))
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}

	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing data:", err)
		return
	}
}

func WritePrimeFinderRangesToFile(primeFinderRanges []PrimeFinderRange) {
	file, err := os.OpenFile("primeFinderRanges.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	sort.Slice(primeFinderRanges, func(i, j int) bool {
		return primeFinderRanges[i].Start < primeFinderRanges[j].Start
	})

	writer := bufio.NewWriter(file)

	for _, val := range primeFinderRanges {
		_, err := writer.WriteString(fmt.Sprintf("%d,%d\n", val.Start, val.End))
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}

	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing data:", err)
		return
	}
}
