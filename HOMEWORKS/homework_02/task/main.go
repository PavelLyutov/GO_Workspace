package main

import (
	"encoding/json"
	"fmt"
	"homework_02/fileReader"
	"homework_02/getInfoGithub"
	"log"
)

func main() {

	fmt.Println("Please enter the file name below:")
	var usernames []string
	usernames = fileReader.TxtReader()
	for i := 0; i < len(usernames); i++ {
		userInfo := getInfoGithub.GetFromGithub(usernames[i])

		userData, err := json.MarshalIndent(userInfo, "", "     ")
		if err != nil {
			log.Fatalf("JSON marshaling failed: %s", err)
		}
		fmt.Println(string(userData))
	}


}
