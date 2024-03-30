package main

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error loading dotenv")
		fmt.Println(err)
	}

	vendorName := os.Getenv("VENDOR_NAME")
	streamUrl := os.Getenv("STREAM_URL")

	stream := NewStream(vendorName, streamUrl)

	stream.RunCheck()

	marshalledStream, _ := json.Marshal(stream)
	response := string(marshalledStream)

	fmt.Println(response)
}
