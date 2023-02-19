package parse

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strings"
)

type Question struct {
	Problem string
	Answer  string
}

// Read and get records of quiz problems and answers from csv file
func ReadQuizProblems(csvFile *os.File) []Question {
	reader := csv.NewReader(csvFile)
	questions := make([]Question, 0)
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if line[0] != "" && line[1] != "" {
			questions = append(questions, Question{Problem: line[0], Answer: strings.TrimSpace(line[1])})
		}
	}
	return questions
}
