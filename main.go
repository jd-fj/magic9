package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	// forces new random seed value
	rand.Seed(time.Now().UnixNano())

	// derives valid number from array
	randomSelection := (rand.Intn(len(adventureTimeAnswers)))

	// selects an answer
	finalAnswer := (adventureTimeAnswers[randomSelection])

	// slowly reveal the answer to user
	printScrambledAnswer(finalAnswer)
}

func shuffle(s string) string {
	r := []rune(s)
	rand.Seed(time.Now().UnixNano())

	for i := range r {
		j := rand.Intn(i + 1)
		r[i], r[j] = r[j], r[i]
	}

	return string(r)
}

func printScrambledAnswer(finalAnswer string) {
	// hide cursor
	fmt.Print("\033[?25l")

	sleepDuration := time.Duration(1 * time.Microsecond)
	incrementer := 1.0

	for i := 0; i < 7; i++ {
		shuffled := shuffle(finalAnswer)
		incrementer *= 3
		sleepDuration += time.Duration(math.Log10(incrementer) * float64(time.Millisecond))

		// print out the scrambled string
		for j := 0; j <= len(shuffled); j++ {
			fmt.Printf("\r%s", shuffled[:j])
			time.Sleep(sleepDuration)
		}
	}

	pauseOneSec()
	fmt.Printf("\r%s\n", finalAnswer)
	pauseOneSec()

	// show the cursor again
	fmt.Print("\033[?25h")
}

func pauseOneSec() { time.Sleep(time.Second) }
