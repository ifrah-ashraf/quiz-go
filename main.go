package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readCsvFile(filepath string) [][]string {
	f, err := os.Open(filepath)

	if err != nil {
		log.Fatal("Unable to read input file "+filepath, err)
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()

	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filepath, err)
	}

	return records
}

func quizTest(records [][]string) (int, error) {
	scanner := bufio.NewScanner(os.Stdin)
	var result int
	score := 0

	for _, row := range records {
		quest := row[0]

		ans := row[1]

		parsedAns, err := strconv.Atoi(ans)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("what is %s ?\n", quest)

		scanner.Scan()
		input := scanner.Text()

		result, err = strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input, please enter a valid number.")
			continue
		}

		if result == parsedAns {
			score++
		}

	}
	return score, nil
}

func main() {
	records := readCsvFile("./problems.csv")

	score, err := quizTest(records)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Your score is %d\n", score)

}

//here i used bufio to handle the input gracefully like if the user give wrong input (some gibbrish string) it will prompt
//the error message and clear the buffer for the next input more better that scanf
