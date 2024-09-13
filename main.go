package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

var correct_answers = 0
var items_len = 0

func exitProgram() int {
	fmt.Printf("\nYou scored %v out of %v.\n", correct_answers, items_len)

	return 3
}

func prog_signals() {
	ch := make(chan os.Signal, 1)

	signal.Notify(ch, os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)

	<-ch
	defer os.Exit(exitProgram())
}

func main() {
	go prog_signals()

	file, err := os.Open("./problems.csv")

	if err != nil {
		fmt.Println("Error Opening File: ", err)
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

	fmt.Printf("You scored %v out of %v.\n", correct_answers, len(items))
}
