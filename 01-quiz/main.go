package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	timeout := flag.Int("timeout", 30, "how much time you have")
	flag.Parse()

	problems, err := readProblems(*csvFilename)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	correct := 0
	fmt.Println("Riddle me this, press return to start")
	fmt.Scanf("\n")

	timer := time.NewTimer(time.Duration(*timeout) * time.Second)

	go func() {
		for _, p := range problems {
			fmt.Printf("%s >> ", p[0])

			var input string
			_, err := fmt.Scanf("%s\n", &input)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}

			input = strings.TrimSuffix(input, "\n")
			if input == p[1] {
				correct += 1
			}
		}
	}()

	<-timer.C
	fmt.Println("Sorry, time is up!")
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
