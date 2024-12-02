package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	//// URL to fetch the file from
	//url := "https://adventofcode.com/2024/day/1/input"
	//
	//// Get the data
	//resp, err := http.Get(url)
	//if err != nil {
	//	fmt.Printf("Error fetching URL: %v\n", err)
	//	return
	//}
	//defer resp.Body.Close() // Close the response body when done
	//
	//// Check if the response status code is 200
	//if resp.StatusCode != http.StatusOK {
	//	fmt.Printf("Error: %v\n", resp.Status)
	//	return
	//}
	//
	//// Read the response body
	//data, err := io.ReadAll(resp.Body)
	//if err != nil {
	//	fmt.Printf("Error reading response body: %v\n", err)
	//	return
	//}
	//
	//inputIds := string(data)
	//fmt.Println(inputIds)

	// Run the dayOne function
	DayOne()
}

func parseListsToSlices(inputFileLocation string) ([]int, []int) {
	// Open the file
	file, err := os.Open(inputFileLocation)

	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		panic(err)
	}
	defer file.Close() // Close the file when done

	inputNumbersLeft := []int{}
	inputNumbersRight := []int{}

	// Create a scanner, to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Split the line into numbers
		numberStrings := strings.Fields(line)

		// Convert the strings to integers, and add them to the inputNumbers slice
		numberLeft, err := strconv.Atoi(numberStrings[0])
		if err != nil {
			fmt.Printf("Error converting string to integer: %v\n", err)
			panic(err)
		}
		inputNumbersLeft = append(inputNumbersLeft, numberLeft)

		numberRight, err := strconv.Atoi(numberStrings[1])
		if err != nil {
			fmt.Printf("Error converting string to integer: %v\n", err)
			panic(err)
		}
		inputNumbersRight = append(inputNumbersRight, numberRight)
	}

	// Sort the slices
	sort.Ints(inputNumbersLeft)
	sort.Ints(inputNumbersRight)
	return inputNumbersLeft, inputNumbersRight
}

func DayOne() {
	fmt.Println("----- Day One -----")
	inputNumbersLeft, inputNumbersRight := parseListsToSlices("dayOne/input.txt")

	// Calculate difference between every pair of numbers
	sumOfDifferences := 0

	for i := 0; i < len(inputNumbersLeft); i++ {
		// Subtract the smaller number from the larger number, and add the result to the sumOfDifferences.
		difference := 0

		if inputNumbersLeft[i] > inputNumbersRight[i] {
			difference = inputNumbersLeft[i] - inputNumbersRight[i]
		} else {
			difference = inputNumbersRight[i] - inputNumbersLeft[i]
		}

		sumOfDifferences += difference
	}

	fmt.Printf("Sum of differences: %v\n", sumOfDifferences)

	similarityScore := 0

	// Count occurrences of numbers from left, in right
	for i := 0; i < len(inputNumbersLeft); i++ {
		// Count the number of occurrences of the number in the right list
		count := 0
		for j := 0; j < len(inputNumbersRight); j++ {
			if inputNumbersLeft[i] == inputNumbersRight[j] {
				count++
			}
		}

		// Calculate the added similarity score
		similarityScore += count * inputNumbersLeft[i]
	}

	fmt.Printf("Similarity score: %v\n\n", similarityScore)
}
}
