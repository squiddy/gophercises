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
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	flag.Parse()

	problems, err := readProblems(*csvFilename)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	correct := 0

	fmt.Println("Riddle me this")

	reader := bufio.NewReader(os.Stdin)
	for _, p := range problems {
		fmt.Printf("%s >> ", p[0])
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		input = strings.TrimSuffix(input, "\n")
		if input == p[1] {
			correct += 1
		}
	}

	fmt.Printf("You got %d out of %d questions correct.", correct, len(problems))
}

func readProblems(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("Error reading file: %w", err)
	}

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("Error parsing file: %w", err)
	}

	return records, nil
}
