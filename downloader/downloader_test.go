package downloader

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDownload_WhenValidUrlIsGiven_ThenDownloadIsPerformed(t *testing.T) {
	tests := []string{
		"https://pastebin.com/raw/v0Sm2sfn",
		"https://pastebin.com/raw/fysHJ7YX",
	}

	for _, address := range tests {

		sut := TextDownloader{}

		c := make(chan string)

		go sut.Download(address, c)

		result := <-c

		assert.True(t, result != "", "Text should be downloaded")
	}
}
