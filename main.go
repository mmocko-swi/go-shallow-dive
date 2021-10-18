package main

import (
	"fmt"
	"solarwinds/pisigma/analyzer/downloader"
)

func main() {

	add := []string{
		"https://pastebin.com/raw/v0Sm2sfn",
		"https://pastebin.com/raw/fysHJ7YX",
	}

	// address1 := "https://pastebin.com/raw/v0Sm2sfn"
	// address2 := "https://pastebin.com/raw/fysHJ7YX"

	results := downloader.MultiDownload(add)

	for index, value := range *results {
		fmt.Printf("%d : %d", index, len(value))
	}

	//textDownloader := downloader.NewDownloader()
	//channel1 := make(chan string)
	// channel2 := make(chan string)
	// go textDownloader.Download(address2, channel1)
	// go textDownloader.Download(address1, channel2)

	// text1, text2 := <-channel1, <-channel2

	// processor := textprocessor.NewProcessor()
	// var stats1 = processor.Process(text1)
	// var stats2 = processor.Process(text2)

	// var results1 = analyzer.Analyze(stats1)
	// var results2 = analyzer.Analyze(stats2)

	// fmt.Println(results1)
	// fmt.Println(results2)

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
