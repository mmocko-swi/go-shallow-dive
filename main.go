package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"solarwinds/pisigma/analyzer/analyzer"
	"solarwinds/pisigma/analyzer/downloader"
	"solarwinds/pisigma/analyzer/merger"
	"solarwinds/pisigma/analyzer/reporter"
	"solarwinds/pisigma/analyzer/textprocessor"
	"sync"
)

func main() {

	add := []string{
		"https://pastebin.com/raw/v0Sm2sfn",
		"https://pastebin.com/raw/fysHJ7YX",
	}

	outputFiles := make([]string, 2, len(add))

	results := downloader.MultiDownload(add)

	wg := new(sync.WaitGroup)
	wg.Add(len(*results))

	for index, value := range *results {
		go func(url string, index int, value string, wg *sync.WaitGroup, outputFiles []string) {
			defer wg.Done()
			defer fmt.Printf("Processing: %s - DONE\n", url)

			fmt.Printf("Processing: %s \n", url)
			processedText := textprocessor.NewProcessor().Process(value)
			analyzedText := analyzer.Analyze(processedText)

			workignDir, _ := os.Getwd()
			outputFile := fmt.Sprintf(workignDir+"\\out%d.json", index)

			reporter.NewJsonReporter().Report(outputFile, analyzedText)
			outputFiles[index] = outputFile
		}(add[index], index, value, wg, outputFiles)
	}

	fmt.Println("Waiting for analysis to finish")
	wg.Wait()
	fmt.Println("Analysis done")

	fmt.Println("Starting to merge")

	var finalResults []analyzer.ResultText
	finalResults = make([]analyzer.ResultText, 0)
	for _, outputFile := range outputFiles {
		rightFile, err := os.Open(outputFile)

		if err != nil {
			fmt.Println("Error")
			break
		}

		defer rightFile.Close()

		rightBytes, _ := ioutil.ReadAll(rightFile)
		rightResults := make([]analyzer.ResultText, 0)
		json.Unmarshal([]byte(rightBytes), &rightResults)
		finalResults = merger.Merge(&finalResults, &rightResults)
	}

	fmt.Println("Merging - done")

	workignDir, _ := os.Getwd()
	finalJson := workignDir + "\\final.json"
	reporter.NewJsonReporter().Report(finalJson, &finalResults)

	fmt.Println("All done!")
}
