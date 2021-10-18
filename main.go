package main

import (
	"fmt"
	"os"
	"solarwinds/pisigma/analyzer/analyzer"
	"solarwinds/pisigma/analyzer/downloader"
	"solarwinds/pisigma/analyzer/reporter"
	"solarwinds/pisigma/analyzer/textprocessor"
	"sync"
)

func main() {

	add := []string{
		"https://pastebin.com/raw/v0Sm2sfn",
		"https://pastebin.com/raw/fysHJ7YX",
	}

	results := downloader.MultiDownload(add)

	wg := new(sync.WaitGroup)
	wg.Add(len(*results))

	for index, value := range *results {
		go func(url string, index int, value string, wg *sync.WaitGroup) {
			defer wg.Done()
			defer fmt.Printf("Processing: %s - DONE\n", url)

			fmt.Printf("Processing: %s \n", url)
			processedText := textprocessor.NewProcessor().Process(value)
			analyzedText := analyzer.Analyze(processedText)

			workignDir, _ := os.Getwd()
			outputFile := fmt.Sprintf(workignDir+"\\out%d.json", index)

			reporter.NewJsonReporter().Report(outputFile, analyzedText)
		}(add[index], index, value, wg)
	}

	fmt.Println("Waiting for analysis to finish")
	wg.Wait()
	fmt.Println("Analysis done")
}
