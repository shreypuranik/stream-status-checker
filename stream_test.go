package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestNewStreamReturnsExpectedStreamStruct(t *testing.T) {
	a := assert.New(t)

	os.Setenv("VENDOR_NAME", "Example Vendor")
	os.Setenv("STREAM_URL", "blah.mp3")

	s := NewStream()

	a.Equal(s.Title, os.Getenv("VENDOR_NAME"), "title of new stream struct is the same as VENDOR_NAME env variable")
	a.Equal(s.Url, os.Getenv("STREAM_URL"), "url of new stream struct is the same as STREAM_URL env variable")

}
