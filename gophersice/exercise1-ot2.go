package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// load cli flags

	wordPtr := flag.String("file", "./exercises.csv", "file path")
	flag.Parse()
	reader := bufio.NewReader(os.Stdin)

	records := make([][]string, 0)

	file := *wordPtr
	records = csvReader(file)

	fmt.Println("Welcome to Go Quiz")
	totalQuestions := len(records)

	var right, wrong int

	for i := 0; i < totalQuestions; i++ {

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

	}

	fmt.Println("Thanks for playing")
	fmt.Println("Totals:", right, "Right", wrong, "Wrong")

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
	// fmt.Println(records)

	return records
}
