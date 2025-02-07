package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	numberSlice := []int{}

	file, err := os.Open("AdventInput2020-#1.txt")
	if err != nil {
		fmt.Println("Unable to open")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		number := scanner.Text()
		convertedNumber, err := strconv.Atoi(number)
		if err != nil {
			fmt.Println("Cannot convert string")
		}
		numberSlice = append(numberSlice, convertedNumber)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	answer := getAnswer(numberSlice)
	fmt.Println(answer)

}

func getAnswer(numberSlice []int) int32 {

	for i := 0; i < len(numberSlice)-1; i++ {
		for j := 0; j < len(numberSlice)-1; j++ {
			for k := 0; k < len(numberSlice); k++ {
				if numberSlice[i]+numberSlice[j]+numberSlice[k] == 2020 {
					return int32((numberSlice[i] * numberSlice[j] * numberSlice[k]))
				}
			}
		}
	}

	return 0
}
