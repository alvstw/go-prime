package finder

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func WritePrimesToFile(path string, primeNumbers []int) {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error closing file:", err)
		}
	}(file)

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

func WritePrimeFinderRangesToFile(path string, primeFinderRanges []PrimeFinderRange) {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error closing file:", err)
		}
	}(file)

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
