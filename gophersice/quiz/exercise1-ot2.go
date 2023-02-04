package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strconv"
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
	// reader := bufio.NewReader(os.Stdin)s
	records := csvReader(*wordPtr)
	theProblems := parseProblems(records)

	var right int
	var answer string

	timer_secs, err := strconv.Atoi(*timePtr)

	if err != nil {
		fmt.Println("Invalid value for the timer")
	}

	timer := time.NewTimer(time.Duration(timer_secs) * time.Second)
	questions := make(chan string, len(theProblems))

	// for i := range theProblems {

	// }

	i := 0
	func() {
		for {
			select {
			case <-timer.C:
				fmt.Println("Game over")
				return

			case <-questions:

				fmt.Println(questions)
				fmt.Print("-> ")
				fmt.Scanln(&answer)

				if theProblems[i].answer == answer {
					right += 1
				}
				i++
			default:
				questions <- theProblems[i].question
			}

		}

	}()

	fmt.Println("Total correct", right, "Out of", i)

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
