package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	csvFileName := flag.String("csv", "problems.csv", "A csv in formate of question,answer")
	timeLimit := flag.Int("limit", 30, "The time limit for the quiz in seconds.")
	flag.Parse()
	file, err := os.Open(*csvFileName)
	if err != nil {
		exit(fmt.Sprintf("Failed to open csv file: %s\n", *csvFileName))
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()

	if err != nil {
		exit("Failed to parse the provided csv file.")
	}

	problems := parseLines(lines)
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	correct := 0
	for i, prob := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, prob.q)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()
		select {
		case <-timer.C:
			fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
			return
		case answer := <-answerCh:
			if answer == prob.a {
				correct++
			}
		}
	}
}

// for i, prob := range problems {
// 	select {
// 	case <-timer.C:
// 		fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
// 		return
// 	default:
// 		fmt.Printf("Problems #%d: %s = \n", i+1, prob.q)
// 		var answer string
// 		fmt.Scanf("%s\n", &answer)
// 		if answer == prob.a {
// 			correct++
// 		}
// 	}
// }

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: line[1],
		}
	}
	return ret
}

type problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Printf(msg)
	os.Exit(1)
}
