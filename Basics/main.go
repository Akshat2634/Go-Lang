package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func printTaskList(taskItems []string) {
	for idx,task := range taskItems {
		fmt.Printf("%d. %s\n",idx+1,task)	
	}
}

func addTask(taskItems []string, task string) []string {
	taskItems = append(taskItems, task)
	return taskItems
}

func getUserInput() string {
	// Example 1: Using bufio.Scanner for reading a line of input
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter your input: ")
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func getUserInputWithFmt() string {
	// Example 2: Using fmt.Scanln for simple input
	var input string
	fmt.Print("Enter your input: ")
	fmt.Scanln(&input)
	return input
}

func getUserInputMultipleWords() string {
	// Example 3: Using bufio.Reader for reading input with spaces
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your input (can contain spaces): ")
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	var greeting = "Hello user. Welcome to our TodoList App!"
	fmt.Fprintf(w,greeting)
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	var taskItems = []string{
		"Do Leetcode problems daily to improve your skills!",
		"Apply for jobs daily to improve your chances of getting hired!",
		"Learn new technologies daily to improve your skills!",
	}

	for idx,task := range taskItems {
		fmt.Fprintf(w,"%d. %s\n",idx+1,task)
	}
}

func main() {
	var taskItems = []string{
		"Do Leetcode problems daily to improve your skills!",
		"Apply for jobs daily to improve your chances of getting hired!",
		"Learn new technologies daily to improve your skills!",
	}

	fmt.Println("#### Welcome to our TodoList App! ####")
	fmt.Println("List of all my todos:")
	printTaskList(taskItems)

	// Example of taking user input to add a new task
	fmt.Println("\n--- User Input Examples ---")
	
	// Example 1: Simple input using bufio.Scanner
	userTask := getUserInput()
	if userTask != "" {
		taskItems = addTask(taskItems, userTask)
		fmt.Println("\nAfter adding your task:")
		printTaskList(taskItems)
	}

	// Example 2: Input using fmt.Scanln
	fmt.Println("\nUsing fmt.Scanln:")
	userTask2 := getUserInputWithFmt()
	if userTask2 != "" {
		taskItems = addTask(taskItems, userTask2)
	}

	// Example 3: Input with spaces using bufio.Reader
	fmt.Println("\nUsing bufio.Reader (supports spaces):")
	userTask3 := getUserInputMultipleWords()
	if userTask3 != "" {
		taskItems = addTask(taskItems, userTask3)
	}

	fmt.Println("\nFinal task list:")
	printTaskList(taskItems)

	fmt.Println("\nStarting web server on :8080...")
	http.HandleFunc("/",helloHandler) 
	http.HandleFunc("/tasks",tasksHandler)

	http.ListenAndServe(":8080",nil)
}