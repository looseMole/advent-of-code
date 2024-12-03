package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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
	DayTwo()
	DayThree()
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

func parseFileToSlicesOfIntegers(inputFileLocation string) [][]int {
	var sliceOfLines = make([][]int, 0, 1000)

	// Open the file
	file, err := os.Open(inputFileLocation)

	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		panic(err)
	}
	defer file.Close() // Close the file when done

	// Create a scanner, to read the file line by line
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		// Split the line by whitespace
		numberStrings := strings.Fields(line)
		var numberSlice []int

		for columnNo := 0; columnNo < len(numberStrings); columnNo++ {
			// Convert the string to integer
			number, err := strconv.Atoi(numberStrings[columnNo])
			if err != nil {
				fmt.Printf("Error converting string to integer: %v\n", err)
				panic(err)
			}

			numberSlice = append(numberSlice, number)

		}

		sliceOfLines = append(sliceOfLines, numberSlice)
	}

	return sliceOfLines
}

func intReportSliceIsSafe(intSlice []int, allowedUnsafeLevelsPerReport int) bool {
	if allowedUnsafeLevelsPerReport < 0 {
		return false
	}

	// Find trend
	positiveDiffCount := 0
	negativeDiffCount := 0

	for i := 1; i < len(intSlice); i++ {
		diff := intSlice[i] - intSlice[i-1]

		if diff < 0 {
			negativeDiffCount++
		} else if diff > 0 {
			positiveDiffCount++
		}
	}

	var ascending bool = positiveDiffCount > negativeDiffCount
	unsafeLevelCount := 0

	for i := 1; i < len(intSlice); i++ {
		diff := intSlice[i] - intSlice[i-1]

		// If any two adjacent numbers are the same, the slope is flat = Strictly speaking no trend.
		// The greatest allowed difference between any two adjacent numbers, is 3.
		if diff == 0 || diff > 3 || diff < -3 {
			unsafeLevelCount++
		}

		// If the trend is not kept
		if ascending && diff < 0 {
			unsafeLevelCount++
		} else if !ascending && diff > 0 {
			unsafeLevelCount++
		}
	}

	if unsafeLevelCount == 0 {
		return true
	}

	if allowedUnsafeLevelsPerReport > 0 {
		// For each level x, try this same function with said level removed.
		for x := 0; x < len(intSlice); x++ {
			var smallerIntSlice []int

			if len(intSlice) >= x+2 {
				smallerIntSlice = append(smallerIntSlice, intSlice[:x]...)
				smallerIntSlice = append(smallerIntSlice, intSlice[x+1:]...)
			} else {
				smallerIntSlice = intSlice[:x]
			}

			if intReportSliceIsSafe(smallerIntSlice, allowedUnsafeLevelsPerReport-1) {
				return true
			}
		}
	}

	return false
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

func DayTwo() {
	fmt.Println("----- Day Two -----")
	var sliceOfLinesOfIntegers [][]int = parseFileToSlicesOfIntegers("dayTwo/input.txt")

	countOfSafeReports := 0
	countOfSafeReportsWithDampener := 0
	for i := 0; i < len(sliceOfLinesOfIntegers); i++ {
		if intReportSliceIsSafe(sliceOfLinesOfIntegers[i], 0) {
			countOfSafeReports++
			countOfSafeReportsWithDampener++
		} else if intReportSliceIsSafe(sliceOfLinesOfIntegers[i], 1) {
			countOfSafeReportsWithDampener++
		}
	}

	fmt.Printf("Amount of safe reports without dampener: %v\n", countOfSafeReports)
	fmt.Printf("Amount of safe reports with dampener: %v\n\n", countOfSafeReportsWithDampener)
}

func parseFileToString(inputFileLocation string) string {
	// Open the file
	file, err := os.Open(inputFileLocation)

	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		panic(err)
	}
	defer file.Close() // Close the file when done

	var fileString string

	// Create a scanner, to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		fileString += line
	}

	return fileString
}

func DayThree() {
	fmt.Println("----- Day Three -----")

	fileString := parseFileToString("dayThree/input.txt")

	exp, _ := regexp.Compile("mul\\((\\d{1,3}),(\\d{1,3})\\)")
	// Int parameter = -1, because if positive, returns up to that many slices of the result.
	matches := exp.FindAllStringSubmatch(fileString, -1)

	sumOfMultiplications := 0

	for i := 0; i < len(matches); i++ {
		integerOne, err := strconv.Atoi(matches[i][1])
		if err != nil {
			fmt.Printf("Error converting string to integer: %v\n", err)
			panic(err)
		}

		integerTwo, err := strconv.Atoi(matches[i][2])
		if err != nil {
			fmt.Printf("Error converting string to integer: %v\n", err)
			panic(err)
		}

		sumOfMultiplications += integerOne * integerTwo
	}

	fmt.Printf("Counted: %v instances of the mul(int, int) command, in file).\n", len(matches))
	fmt.Printf("The sum of the correctly-formed multiplications should be: %v.\n", sumOfMultiplications)
}
