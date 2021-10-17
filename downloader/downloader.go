package downloader

import (
	"io/ioutil"
	"net/http"
)

type Downloader interface {
	Download(address string, c chan string) error
}

type TextDownloader struct {
}

func (downloader *TextDownloader) Download(address string, c chan string) error {
	response, err := http.Get(address)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	c <- string(responseData)
	return nil
}
