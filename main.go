package main

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error loading dotenv")
		fmt.Println(err)
	}

	stream := NewStream()
	stream.RunCheck()

	marshalledStream, _ := json.Marshal(stream)
	response := string(marshalledStream)

	fmt.Println(response)
}
