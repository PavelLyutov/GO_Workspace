package fileReader

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func TxtReader() []string {
	var f string
	_, err := fmt.Scanln(&f)
	if err != nil {
		log.Fatalf("wrong input")
	}

	//Paths should start after the project directory
	file, err := os.Open(f)
	if err != nil {
		log.Fatalf("wrong input, please enter the path after the project directory")
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	var usernames []string
	for scanner.Scan() {
		// fmt.Println(scanner.Text())  // print the token in unicode-char
		usernames = append(usernames, scanner.Text())
	}
	return usernames
}
