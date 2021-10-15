package textprocessor

import "strings"

type Processor interface {
	Process(text string) TextStatistics
}

type TextProcessor struct {
}

type TextStatistics struct {
	OriginalText string
	Stats        map[string]int
}

func (processor *TextProcessor) Process(text string) TextStatistics {
	textStatistics := TextStatistics{text, make(map[string]int)}
	for _, word := range strings.Fields(text) {
		textStatistics.Stats[word]++
		println(word)
	}
	return textStatistics
}
