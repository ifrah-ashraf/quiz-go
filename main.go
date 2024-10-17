package main

import (
	"bufio"
	"context"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
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

// the io.Reader (third param in func quiztest) interface give proper control on the type of input we are passing to the quiz function such as -:

//Files: The *os.File type (like os.Stdin) implements io.Reader.
//Strings: strings.NewReader wraps a string to implement io.Reader.
//Network Connections: net.Conn implements io.Reader for reading data from network connections.

func quizTest(ctx context.Context, records [][]string, input io.Reader) (int, error) {
	scanner := bufio.NewScanner(input)
	var result int
	score := 0

	for _, row := range records {
		quest := row[0]
		ans := row[1]

		select {
		case <-ctx.Done():
			return score, ctx.Err()

		default:

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

	}
	return score, nil
}

func main() {
	records := readCsvFile("./problems.csv")

	timePtr := flag.Int("time", 2, "a timer for quiz")

	flag.Parse()

	fmt.Printf("Your time for quiz is %v sec\n", *timePtr)

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(*timePtr)*time.Second)
	defer cancelFunc()

	score, err := quizTest(ctx, records, os.Stdin)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Your score is %d\n", score)

}

//here i used bufio to handle the input gracefully like if the user give wrong input (some gibbrish string) it will prompt
//the error message and clear the buffer for the next input more better that scanf

//how to use flag ?
// flag declared should have format like flag.format_type(name , value , description)

// the context
