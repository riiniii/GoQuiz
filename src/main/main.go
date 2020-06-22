package main

import (
	"encoding/csv"
	"flag"
	"os"
	"fmt"
	"strings"
	"time"
	"math/rand"
)

type problem struct{
	q string
	a []string
} 

func main(){
	// my flags
	csvFilename, timeLimit, shuffle := setFlags()
	lines := readFile(csvFilename)
	runProblems(lines, timeLimit, shuffle)
}

func setFlags() (*string, *int, *bool){
	csvFilename := flag.String("csv", "questions.csv", "a csv file in the form of 'question, answer'")
	timeLimit := flag.Int("timeLimit", 20, "the time limit for the whole quiz")
	shuffle := flag.Bool("shuffle", false, "shuffle the questions")
	flag.Parse()
	return csvFilename, timeLimit, shuffle
}
func readFile(csvFilename *string) [][]string{
	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the csv file: %s\n", *csvFilename))
	}

	// open file 
	r := csv.NewReader(file)
	lines, err := r.ReadAll()

	if err != nil {
		exit(fmt.Sprintf("Failed to parse the the provided csv file: %s", err))
	}
	return lines
}
func shuffle(problems []problem) []problem{
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(problems), func(i, j int) { problems[i],problems[j] = problems[j], problems[i] }) 
	return problems
}
func runProblems(lines [][]string, timeLimit *int, toShuffle *bool){
	problems := parseLines(lines)
	
	if *toShuffle {
		problems = shuffle(problems)
	}

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	score :=  0
	problemloop: 
	for i, p := range problems { // for each problem
		fmt.Printf("Problem #%d: %s", i+1, p.q)
		answerCh := make(chan string)
		go waitInput(answerCh) 
		select { // select is like switch, except for goroutines
			case <- timer.C:
				printScoreOutOf(score, len(problems))
				break problemloop
				// Scanf gets rid of trailing spaces, captures new line with \n
				// Scanf blockes our timer, so may not here from timer until after we respond 
				// create a goroutine for answer itself
			case answer := <-answerCh: 
 				if isCorrectAnswer(answer, p.a[0]) {
					score++
					printScore(1, score)
				} else {
					printScore(0, score)
				}
		}
		
	}
	printScoreOutOf(score, len(problems))
}

func waitInput(answerCh chan string) {
	var answer string
	fmt.Scanf("%s\n", &answer) 
	answerCh <- answer // send our answer string to our channel 
}

func isCorrectAnswer(usersAns string, correctAns string) bool {
	return strings.ToLower(usersAns) == strings.ToLower(strings.TrimSpace(correctAns))
}

func printScore(pointsEarned int, score int){
	fmt.Printf("+%d point... (Total Score: %d)\n", pointsEarned, score)
}

func printScoreOutOf(score int, totalPossible int){
	fmt.Printf("You scored %d out of %d\n", score, totalPossible)
}

func parseLines(lines [][]string) []problem{
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem {
			q: line[0],
			a: line[1:],
		}
	}
	return ret
}
func exit(msg string){
	fmt.Println(msg)
	os.Exit(1)
}


