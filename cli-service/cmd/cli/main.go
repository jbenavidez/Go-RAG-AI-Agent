package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to the NYC Capital Project RAG AI CLI")
	fmt.Println("Ask any question related to nyc capitals projects from 2023-2025.")
	fmt.Println("The AI agent will provide answers based on the data provided by data.cityofnewyork.us")
	fmt.Println("Type your question and press Enter. Type 'exit' to quit.\n")

	for {
		fmt.Print("\nEnter question: ")
		question, _ := reader.ReadString('\n')
		question = strings.TrimSpace(question)
		if question == "exit" {
			break
		}

		//answer, err := cliService.Ask(context.Background(), question)

		fmt.Println("\nAnswer:")

	}
}
