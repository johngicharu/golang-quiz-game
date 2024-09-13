package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

var correct_answers = 0
var items_len = 0

func exitProgram(msg string) int {
	fmt.Printf(msg)

	return 3
}

func prog_signals() {
	ch := make(chan os.Signal, 1)

	signal.Notify(ch, os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)

	<-ch
	defer os.Exit(exitProgram(fmt.Sprintf("\nYou scored %v out of %v.\n", correct_answers, items_len)))
}

func main() {
	go prog_signals()

	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")

	flag.Parse()

	file, err := os.Open(*csvFilename)

	if err != nil {
		exitProgram(fmt.Sprintf("Error Opening the CSV File: %s\n", *csvFilename))
	}

	// Close file once the function is done running
	defer file.Close()

	file_reader := csv.NewReader(file)

	items, err := file_reader.ReadAll()
	items_len = len(items)

	if err != nil {
		fmt.Println("File Reader Read Error: ", err)
	}

	for i, item := range items {
		var answer string
		fmt.Printf("Problem #%v: %v = ", i+1, item[0])
		_, err := fmt.Scanln(&answer)

		if err != nil {
			fmt.Println("Sorry, seems an invalid value was entered")
			break
		}

		if answer == item[1] {
			correct_answers += 1
		}
	}

	exitProgram(fmt.Sprintf("You scored %v out of %v.\n", correct_answers, len(items)))
}
