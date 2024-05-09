package main

import (
	edit "editedtext/edited"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Error usage: go run main.go <input.txt> <output.txt>")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	file, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error occurred during file opening:", err)
		return
	}

	content := strings.Split(string(file), "\n")

	result := edit.EditedText(content)

	finalOutPut := result

	output := os.WriteFile(outputFile, []byte(finalOutPut), 0o644)
	if output != nil {
		fmt.Println("Error occurred during writting file", output)
		return
	}
}
