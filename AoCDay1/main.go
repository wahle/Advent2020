package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func convertToInt(stringInputs []string) ([]int, error) {
	var intInputs = []int{}

	for _, stringValue := range stringInputs {
		intValue, err := strconv.Atoi(stringValue)
		if err != nil {
			return intInputs, err
		}
		intInputs = append(intInputs, intValue)
	}
	return intInputs, nil
}

func addValuesTwo(vals []int, desiredVal int) (int, int, error) {
	for startPoint := 0; startPoint < len(vals); startPoint++ {
		for endPoint := len(vals) - 1; endPoint > 0; endPoint-- {
			addedValue := vals[startPoint] + vals[endPoint]
			if addedValue == 2020 {
				return vals[startPoint], vals[endPoint], nil
			}
		}
	}
	return 0, 0, nil
}

func addValuesThree(vals []int, desiredVal int) (int, int, int, error) {
	for startPoint := 0; startPoint < len(vals); startPoint++ {
		for endPoint := len(vals) - 1; endPoint > 0; endPoint-- {
			for thirdPoint := 1; thirdPoint < len(vals); thirdPoint++ {
				addedValue := vals[startPoint] + vals[endPoint] + vals[thirdPoint]
				if addedValue == 2020 {
					return vals[startPoint], vals[endPoint], vals[thirdPoint], nil
				}
			}
		}
	}
	return 0, 0, 0, nil
}

func main() {
	input, err := readLines("./input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	intInput, err := convertToInt(input)
	if err != nil {
		log.Fatalf("convertToInt: %s", err)
	}
	sort.Ints(intInput)

	var splitPoint = len(intInput) / 2
	fmt.Println("splitPoint:", splitPoint)

	foundVal1, foundVal2, err := addValuesTwo(intInput, 2020)
	if err != nil {
		log.Fatalf("addValuesTwo: %s", err)
	}
	fmt.Printf("Found Value 1: %d, Found Value 2: %d\n", foundVal1, foundVal2)
	fmt.Printf("Answer is: %d\n", foundVal1*foundVal2)
	foundVal1, foundVal2, foundVal3, err := addValuesThree(intInput, 2020)
	if err != nil {
		log.Fatalf("addValuesThree: %s", err)
	}
	fmt.Printf("Found Value 1: %d, Found Value 2: %d, Found Value 3: %d\n", foundVal1, foundVal2, foundVal3)
	fmt.Printf("Answer is: %d\n", foundVal1*foundVal2*foundVal3)
}
