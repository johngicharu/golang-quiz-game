package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

var correct_answers = 0
var items_len = 0

type problem struct {
	q string
	a string
}

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

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))

	for i, line := range lines {
		ret[i] = problem{q: line[0], a: line[1]}
	}

	return ret
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

	// Read the entire csv file entirely
	items, err := file_reader.ReadAll()
	items_len = len(items)

	if err != nil {
		exitProgram(fmt.Sprintln("File Reader Read Error: ", err))
	}

	problems := parseLines(items)

	for i, p := range problems {
		var answer string

		fmt.Printf("Problem #%v: %s = ", i+1, p.q)
		_, err := fmt.Scanf("%s\n", &answer)

		if err != nil {
			exitProgram("Sorry, seems an invalid value was entered")
			break
		}

		if answer == strings.TrimSpace(p.a) {
			correct_answers++
		}
	}

	exitProgram(fmt.Sprintf("You scored %v out of %v.\n", correct_answers, len(items)))
}
