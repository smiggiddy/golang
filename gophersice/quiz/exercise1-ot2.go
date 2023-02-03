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

	ch := make(chan string)

	fmt.Println("Welcome to Go Quiz")
	flag.Parse()
	reader := bufio.NewReader(os.Stdin)
	records := csvReader(*wordPtr)
	theProblems := parseProblems(records)

	var right int

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
				ch <- "test"
			}
			fmt.Println(theProblems[i].question)
			fmt.Print("-> ")

			text, _ := reader.ReadString('\n')
			// convert CRLF to LF
			text = strings.Replace(text, "\n", "", -1)

			if theProblems[i].answer == text {
				right += 1
			}
			i++

		}

	}()
	hello := <-ch
	fmt.Println("Thanks for playing", hello)
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
