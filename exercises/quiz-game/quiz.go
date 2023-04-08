package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func readCsvFile(filepath string) {
	var questions []string
	var answers []string
	var user_answers string
	var correct_answers int
	csvfile, err := os.Open(filepath)
	if err != nil {
		log.Fatalln("Could not open csv file because of ", err)
	}
	defer func(csvfile *os.File) {
		err := csvfile.Close()
		if err != nil {
			log.Fatalln("Error closing the csv file ", err)
		}
	}(csvfile)
	csvReader := csv.NewReader(csvfile)
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln("Not able to read from csv file", err)
		}
		questions := append(questions, record[0])
		answers := append(answers, record[1])
		fmt.Println("The question is ", questions)
		fmt.Println("Enter your answer: ")
		inputReader := bufio.NewReader(os.Stdin)
		user_answers, _ = inputReader.ReadString('\n')
		user_answers = strings.TrimSuffix(user_answers, "\n")
		if user_answers == answers[0] {
			correct_answers++
		}
	}
	fmt.Println("The number of correct answer is ", correct_answers)
}

func main() {
	readCsvFile("problems.csv")
}
