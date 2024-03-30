package main

import (
	"log"
	"net/http"
	"os"
)

type Stream struct {
	Title       string `json:"title"`
	Url         string `json:"url"`
	StatusCode  int    `json:"status_code"`
	Isavailable bool   `json:"is_available"`
}

// RunCheck runs the relavent checks
// the stream url and updates the
// Isavailable value
func (s *Stream) RunCheck() {
	resp, err := http.Get(s.Url)
	if err != nil {
		log.Fatal(err)
	}

	s.StatusCode = resp.StatusCode

	if s.StatusCode == 200 {
		s.Isavailable = true
	}
}

// NewStream sets up a new Stream Struct
func NewStream() Stream {
	title := os.Getenv("VENDOR_NAME")
	url := os.Getenv("STREAM_URL")

	stream := Stream{
		Title:       title,
		Url:         url,
		Isavailable: false,
	}

	return stream
}
