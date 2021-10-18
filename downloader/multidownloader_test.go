package downloader

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type downloaderMock struct {
}

func (d downloaderMock) Download(url string, c chan string) error {
	c <- url
	return nil
}

func newMockDownloader() Downloader {
	return downloaderMock{}
}

func TestMultiDownload(t *testing.T) {
	origNewDownloader := newDownloader
	newDownloader = newMockDownloader
	defer func() { newDownloader = origNewDownloader }()

	urls := []string{"a", "b"}
	results := MultiDownload(urls)

	assert.Equal(t, urls, *results)
}
