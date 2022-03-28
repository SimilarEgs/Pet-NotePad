package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	records := []string{}
	scanner := bufio.NewScanner(os.Stdin)
	Exit := 0
	
	for Exit == 0 {
		fmt.Print("\nEnter a command and data: ")
		scanner.Scan()

		line := strings.SplitAfterN(scanner.Text(), " ", 2)
		command := strings.ReplaceAll(line[0], " ", "")
		command = strings.ToLower(command)

		switch command {
		case "create":
			if len(records) < 5 {
				records = append(records, line[1])
				fmt.Println("[OK] The note was successfully created")
			} else {
				fmt.Println("[Error] Notepad is full")
			}
		case "list":
			if len(records) >= 1 {
				for i := 0; i < len(records); i++ {
					fmt.Printf("[Info] %d: %s\n", i+1, records[i])
				}
			} else {
				fmt.Println("[Info] No records yet")
			}
		case "clear":
			records = nil
			fmt.Println("[OK] All notes were successfully deleted")
		case "exit":
			fmt.Println("[Info] Bye!")
			Exit++	
		default:
			fmt.Println("[Error] Wrong command")
		}

	}

}
