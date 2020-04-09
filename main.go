package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"propellerads-test/handler"
)

func main() {
	fromFile()
}

func fromFile() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	file.Close()

	h := handler.New()
	fmt.Println("Result:")
	for _, group := range h.GetGroups(txtlines) {
		fmt.Printf("%v\n", group)
	}
}
