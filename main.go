package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"quiz/operate"
	"quiz/parse"
	"strings"
)

var (
	quizQuestions = flag.String("csv", "problems.csv", "CSV file that constains problems with solutions")
	timeLimit     = flag.Int("limit", 60, "Time limit for the quiz in seconds")
)

func main() {
	flag.Parse()

	if !strings.HasSuffix(*quizQuestions, ".csv") {
		log.Fatalf("The file %s is not a CSV", *quizQuestions)
	}

	csvFile, err := os.Open(*quizQuestions)
	if err != nil {
		log.Fatalf("Could not open file %s", *quizQuestions)
	}
	questions := parse.ReadQuizProblems(csvFile)
	percentSolved := operate.SolveQuestions(&questions, timeLimit)
	fmt.Printf("Correct Questions: %.2f%%\n", percentSolved)

}
