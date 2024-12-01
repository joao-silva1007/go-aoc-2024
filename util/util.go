package util

import (
	"bufio"
	"log"
	"os"
)

func ReadInput(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	var strArray []string

	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		strArray = append(strArray, fileScanner.Text())
	}

	if err := fileScanner.Err(); err != nil {
		log.Fatal(err)
	}
	return strArray
}
