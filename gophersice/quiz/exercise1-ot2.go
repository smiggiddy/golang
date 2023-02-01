package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type problems struct {
	question string
	answer   string
}

// load cli flags
var (
	wordPtr = flag.String("file", "./exercises.csv", "file path")
	timePtr = flag.String("timer", "30", "Used to set the default timer for the program")
)

func main() {
	fmt.Println("Welcome to Go Quiz")
	flag.Parse()
	reader := bufio.NewReader(os.Stdin)
	records := csvReader(*wordPtr)
	theProblems := parseProblems(records)

	fmt.Println(theProblems)

	totalQuestions := len(records)

	var right, wrong int

	timer_secs, err := strconv.Atoi(*timePtr)

	if err != nil {
		fmt.Println("Invalid value for the timer")
	}

	timer := time.NewTimer(time.Duration(timer_secs) * time.Second)
	i := 0

	// timer_done := timer.Stop()
	func() {
		for {
			select {
			case <-timer.C:
				fmt.Println("Game over")
				return

			default:
			}
			if i < totalQuestions {
				fmt.Println("What is the answer:", records[i][0])
				fmt.Print("-> ")

				text, _ := reader.ReadString('\n')
				// convert CRLF to LF
				text = strings.Replace(text, "\n", "", -1)

				if records[i][1] == text {
					fmt.Println("Correct")
					right += 1
				} else {
					fmt.Println("Wrong")
					wrong += 1
				}
				i++
				continue
			}

		}

	}()

	fmt.Println("Thanks for playing")
	fmt.Println("Totals:", right, "Right", wrong, "Wrong")

}

func parseProblems(d [][]string) []problems {
	formated := make([]problems, len(d))
	// ret := make([]problems, len(d))

	// my  method
	for i, line := range d {
		formated[i].question = line[0]
		formated[i].answer = line[1]
	}

	// gophercise method
	// for i, line := range d {
	// 	ret[i] = problems{
	// 		question: line[0],
	// 		answer:   line[1],
	// 	}
	// }

	return formated
}

func csvReader(file string) [][]string {
	// Function reads csv file and returns array

	// 1. Open the file
	recordFile, err := os.Open(file)

	if err != nil {
		fmt.Println("An error encountered ::", err)
	} // 2. Initialize the reader
	reader := csv.NewReader(recordFile) // 3. Read all the records
	records, _ := reader.ReadAll()      // 4. Iterate through the records as you wish

	return records
}
