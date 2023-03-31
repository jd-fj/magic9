package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	answers := []string{
		"It is certain",
		"It is decidedly so",
		"Without a doubt",
		"Yes definitely",
		"You may rely on it",
		"As I see it yes",
		"Most likely",
		"Outlook good",
		"Yes",
		"Signs point to yes",
		"Reply hazy try again",
		"Ask again later",
		"Better not tell you now",
		"Cannot predict now",
		"Concentrate and ask again",
		"Don't count on it",
		"My reply is no",
		"My sources say no",
		"Outlook not so good",
		"Very doubtful",
	}

	finalAnswer := (answers[rand.Intn(len(answers))])
	printScrambledAnswer(finalAnswer)
}

func shuffle(s string) string {
	r := []rune(s)
	rand.Seed(time.Now().UnixNano())

	for i := len(r) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		r[i], r[j] = r[j], r[i]
	}

	return string(r)
}

func printScrambledAnswer(finalAnswer string) {
	// hide cursor
	fmt.Print("\033[?25l")

	sleepDuration := time.Duration(1 * time.Microsecond)
	incrementer := 5

	for i := 0; i < 7; i++ {
		shuffled := shuffle(finalAnswer)
		incrementer *= 3

		sleepDuration += time.Duration(math.Log(float64(incrementer)) * float64(time.Millisecond))
		// print out the scrambled string
		for j := 0; j < len(shuffled); j++ {
			fmt.Printf("\r%s", shuffled[:j])
			time.Sleep(sleepDuration)
		}
	}


	time.Sleep(1 + time.Second)
	fmt.Printf("\r%s\n", finalAnswer)
	time.Sleep(1 + time.Second)
	// show the cursor again
	fmt.Print("\033[?25h")
}
