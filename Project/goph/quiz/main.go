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
	// get the file
	filename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	// timer parse
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	flag.Parse()

	// open this file
	file, err := os.Open(*filename)
	if err != nil {
		fmt.Println("Failed to open file")
		os.Exit(1)
	}

	// reader
	r := csv.NewReader(file)
	lines, err := r.ReadAll() // read all line in file
	if err != nil {
		fmt.Println("failed to read csv file")
	}
	problems := parseLines(lines)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	// fmt.Print(problems)
	correct := 0

	for i, p := range problems {
		fmt.Printf("Problem %d: %s = ", i+1, p.question)
		answerCh := make(chan string)
		go func() {
			var storeans string
			fmt.Scanf("%s\n", &storeans)
			answerCh <- storeans
		}()
		select {
		case <-timer.C:
			fmt.Printf("%d / %d\n", correct, len(problems))
			return
		case storeanswer := <-answerCh:
			if storeanswer == p.answer {
				correct++
			}
			/* default:
			fmt.Printf("Problem %d: %s = ", i+1, p.question)
			var storeans string
			fmt.Scanf("%s\n", &storeans)
			if storeans == p.answer {
				correct++
			} */
		}

	}

	fmt.Printf("%d / %d\n", correct, len(problems))
}

// storing values in struct
func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}
	return ret
}

type problem struct {
	question string
	answer   string
}
