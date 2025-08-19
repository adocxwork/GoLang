package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func getSplitInput() (string, string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	words := strings.Fields(input)
	if len(words) > 1 {
		restOfString := strings.Join(words[1:], " ")
		return words[0], restOfString
	}
	return words[0], ""
}

func removeElementInPlace(slice []string, index int) ([]string, error) {
	if index < 0 || index >= len(slice) {
		return slice, fmt.Errorf("index out of range")
	}
	copy(slice[index:], slice[index+1:])
	return slice[:len(slice)-1], nil
}

func displayTasks(tasks []string) {
	if len(tasks) == 0 {
		fmt.Println("No tasks to display.")
		return
	}
	for i, task := range tasks {
		fmt.Printf("%d. %v\n", i+1, task)
	}
}

func main() {
	var tasks []string
	fmt.Println("---ToDo App---")

	for {
		// Display menu options
		fmt.Println("\n---Options---")
		fmt.Println("add <task>")
		fmt.Println("remove <task no>")
		fmt.Println("view")
		fmt.Println("exit")
		fmt.Print("> ")

		// Read and parse user input
		input1, input2 := getSplitInput()

		switch input1 {
		case "add":
			// Add task
			if input2 != "" {
				tasks = append(tasks, input2)
				fmt.Println("Task added!")
			} else {
				fmt.Println("Error: Task description cannot be empty.")
			}
		case "remove":
			// Remove task
			idx, err := strconv.Atoi(input2)
			if err != nil || idx <= 0 {
				fmt.Println("Error: Invalid task number.")
				continue
			}
			tasks, err = removeElementInPlace(tasks, idx-1)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Task removed!")
			}
		case "view":
			// View tasks
			displayTasks(tasks)
		case "exit":
			// Exit the application
			fmt.Println("Exiting...")
			return
		default:
			// Invalid command
			fmt.Println("Invalid input. Type 'add', 'remove', 'view' or 'exit'.")
		}
	}
}
