// Learned how to open and scan a text file and check errors
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("AdventInput.txt")

	if err != nil {
		fmt.Println("No file found")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		values := scanner.Text()
		fmt.Println(values)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}
}
