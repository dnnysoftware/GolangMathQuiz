package operate

import (
	"fmt"
	"math/rand"
	"quiz/parse"
	"time"
)

func SolveQuestions(questions *[]parse.Question, timeLimit *int) float32 {

	shuffleQuestions(*questions)
	var correct int8 = 0
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	for i, q := range *questions {
		fmt.Printf("\nQuestion #%d: %s = ", i+1, q.Problem)
		ansCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			ansCh <- answer
		}()
		select {
		case <-timer.C:
			fmt.Println("Ran Out Of Time Heres How Many Are Correct:")
			return calculateCorrectPercentage(correct, int8(len(*questions)))
		case answer := <-ansCh:
			if answer == q.Answer {
				fmt.Println("Correct!")
				correct++
			} else {
				fmt.Println("Incorrect!")
			}
		}
	}

	return calculateCorrectPercentage(correct, int8(len(*questions)))
}

func shuffleQuestions(questions []parse.Question) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for i := 0; i < len(questions); i++ {
		j := r.Intn(len(questions))

		questionA := questions[i]
		questionB := questions[j]

		questions[i] = questionB
		questions[j] = questionA
	}
}

func calculateCorrectPercentage(correct int8, total int8) float32 {
	return (float32(correct) / float32(total)) * 100
}
