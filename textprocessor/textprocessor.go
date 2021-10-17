package textprocessor

import "strings"

type Processor interface {
	Process(text string) TextStatistics
}

type textProcessor struct {
}

type TextStatistics struct {
	OriginalText string
	Stats        map[string]int
}

func NewProcessor() *textProcessor {
	txtProcessor := textProcessor{}
	return &txtProcessor
}

func (processor *textProcessor) Process(text string) TextStatistics {
	textStatistics := TextStatistics{text, make(map[string]int)}
	for _, word := range strings.Fields(text) {
		textStatistics.Stats[word]++
	}
	return textStatistics
}
