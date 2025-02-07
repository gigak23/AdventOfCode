package main

import (
	"bufio"
	"fmt"
	"os"

	//"strings"
	"strconv"
	"unicode/utf8"
)

type Separations struct {
	minNum int
	maxNum int
	text   string
	letter string
}

func main() {

	s := Separations{}
	count := 0
	file, err := os.Open("AdventInput2020-#2.txt")
	if err != nil {
		fmt.Println("Cannot Open File")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		var data Separations = splitText(line, s)
		valid := countValidPassword(data)
		if valid == true {
			count += 1
		}
		fmt.Println(data)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(count)

}

func splitText(line string, s Separations) Separations {
	stringLength := utf8.RuneCountInString(line)
	for i := 0; i < stringLength; i++ {
		if string(line[1]) == "-" {
			minNum, err := strconv.Atoi(string(line[0]))
			s.minNum = minNum
			if err != nil {
				fmt.Println("Not successful")
			}
			if string(line[3]) == " " {
				maxNum, err := strconv.Atoi(string(line[2]))
				s.maxNum = maxNum
				if err != nil {
					fmt.Println("Not successful")
				}
				s.letter = string(line[4])
				if i >= 7 {
					s.text += string(line[i])
				}
			} else if string(line[4]) == " " {
				maxNum, err := strconv.Atoi(string(line[2]) + string(line[3]))
				s.maxNum = maxNum
				if err != nil {
					fmt.Println("Not successful")
				}
				s.letter = string(line[5])
				if i >= 8 {
					s.text += string(line[i])
				}
			}
		} else if string(line[2]) == "-" {
			minNum, err := strconv.Atoi(string(line[0]) + string(line[1]))
			s.minNum = minNum
			if err != nil {
				fmt.Println("Not successful")
			}
			if string(line[5]) == " " {
				maxNum, err := strconv.Atoi(string(line[3]) + string(line[4]))
				s.maxNum = maxNum
				if err != nil {
					fmt.Println("Not successful")
				}
				s.letter = string(line[6])
				if i >= 9 {
					s.text += string(line[i])
				}
			}
		}
	}
	return s

}

func countValidPassword(data Separations) bool {
	stringLength := utf8.RuneCountInString(data.text)
	letterCount := 0
	checkValid := false
	for i := 0; i < stringLength; i++ {

		if string(data.text[i]) == data.letter {
			letterCount += 1
		}
	}

	if letterCount >= data.minNum && letterCount <= data.maxNum {
		checkValid = true
	}

	return checkValid
}
