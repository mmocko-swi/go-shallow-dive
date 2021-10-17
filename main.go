package main

import (
	"solarwinds/pisigma/analyzer/downloader"
	"solarwinds/pisigma/analyzer/textprocessor"
)

func main() {
	address1 := "https://pastebin.com/raw/v0Sm2sfn"
	address2 := "https://pastebin.com/raw/fysHJ7YX"

	textDownloader := downloader.TextDownloader{}
	channel1 := make(chan string)
	channel2 := make(chan string)
	go textDownloader.Download(address2, channel1)
	go textDownloader.Download(address1, channel2)

	text1, text2 := <-channel1, <-channel2

	processor := textprocessor.NewProcessor()
	processor.Process(text1)
	processor.Process(text2)

	// counts := make(map[string]int)
	// input := bufio.NewScanner(os.Stdin)
	// for input.Scan() {
	// 	counts[input.Text()]++
	// }

	// for line, count := range counts {
	// 	if count > 1 {
	// 		fmt.Printf("%d\t%s\n", count, line)
	// 	}
	// }

}
