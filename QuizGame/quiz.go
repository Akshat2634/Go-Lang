package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to the Quiz Game!")
	fmt.Println("Enter your name:")
	var name string
	fmt.Scan(&name)
	fmt.Println("Hello, " + name + "! Welcome to the Quiz Game!")

	fmt.Println("What is your age?")
	var age uint	
	fmt.Scan(&age)

	if age >= 10 {
		fmt.Println("You are old enough to play!")
	} else {
		fmt.Println("You are not old enough to play!")
		return
	}
	
	score := 0
	totalQuestions := 5
	
	// Question 1
	fmt.Printf("Question 1: What is better, the RTX 4090 or the RTX 4080? ")
	var answer1 string
	fmt.Scan(&answer1)
	fmt.Println("You answered: " + answer1)
	if strings.ToLower(answer1) == "rtx" || strings.Contains(strings.ToLower(answer1), "4090") {
		fmt.Println("Correct!")
		score++
	} else {
		fmt.Println("Incorrect! The correct answer is RTX 4090.")
	}
	
	// Question 2
	fmt.Printf("Question 2: What is the capital of France? ")
	var answer2 string
	fmt.Scan(&answer2)
	fmt.Println("You answered: " + answer2)
	if strings.ToLower(answer2) == "paris" {
		fmt.Println("Correct!")
		score++
	} else {
		fmt.Println("Incorrect! The correct answer is Paris.")
	}
	
	// Question 3
	fmt.Printf("Question 3: What is 15 + 27? ")
	var answer3 string
	fmt.Scan(&answer3)
	fmt.Println("You answered: " + answer3)
	if answer3 == "42" {
		fmt.Println("Correct!")
		score++
	} else {
		fmt.Println("Incorrect! The correct answer is 42.")
	}
	
	// Question 4
	fmt.Printf("Question 4: Which programming language is this quiz written in? ")
	var answer4 string
	fmt.Scan(&answer4)
	fmt.Println("You answered: " + answer4)
	if strings.ToLower(answer4) == "go" || strings.ToLower(answer4) == "golang" {
		fmt.Println("Correct!")
		score++
	} else {
		fmt.Println("Incorrect! The correct answer is Go (or Golang).")
	}
	
	// Question 5
	fmt.Printf("Question 5: What year was Go programming language first released? ")
	var answer5 string
	fmt.Scan(&answer5)
	fmt.Println("You answered: " + answer5)
	if answer5 == "2009" {
		fmt.Println("Correct!")
		score++
	} else {
		fmt.Println("Incorrect! The correct answer is 2009.")
	}

	// Final score and feedback
	fmt.Println("\n" + strings.Repeat("=", 30))
	fmt.Println("Quiz Complete!")
	fmt.Println("You scored: " + strconv.Itoa(score) + " out of " + strconv.Itoa(totalQuestions) + " points")
	
	percentage := float64(score) / float64(totalQuestions) * 100
	fmt.Printf("Percentage: %.1f%%\n", percentage)
	
	if percentage >= 80 {
		fmt.Println("Excellent work, " + name + "! 🎉")
	} else if percentage >= 60 {
		fmt.Println("Good job, " + name + "! 👍")
	} else if percentage >= 40 {
		fmt.Println("Not bad, " + name + ". Keep practicing! 📚")
	} else {
		fmt.Println("Better luck next time, " + name + "! 💪")
	}
	
	fmt.Println("Thanks for playing!")
}