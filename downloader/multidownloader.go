package downloader

import "sync"

var newDownloader = createDownloader

func createDownloader() Downloader {
	return NewDownloader()
}

func MultiDownload(urls []string) *[]string {

	wg := new(sync.WaitGroup)
	wg.Add(len(urls))

	//var results [len(urls)]string
	results := make([]string, len(urls))

	for index, url := range urls {
		go func(wg *sync.WaitGroup, index int, url string, res []string) {
			defer wg.Done()
			c := make(chan string)

			var downloader = newDownloader()
			go downloader.Download(url, c)

			result := <-c
			results[index] = result
		}(wg, index, url, results)
	}

	wg.Wait()

	return &results
}
