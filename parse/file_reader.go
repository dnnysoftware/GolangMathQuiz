package parse

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

type question struct {
	problem string
	answer  string
}

// Read and get records of quiz problems and answers from csv file
func ReadQuizProblems(csvFile *os.File) []question {
	reader := csv.NewReader(csvFile)
	questions := make([]question, 0)
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if line[0] != "" && line[1] != "" {
			questions = append(questions, question{line[0], line[1]})
		}
	}
	return questions
}
