package hw03frequencyanalysis

import (
	"math"
	"regexp"
	"sort"
	"strings"
)

type WordAndFreq struct {
	Word      string
	Frequency int
}

var wordValidator = regexp.MustCompile("(([А-Яа-яA-Za-z-]+)([-.]*)([А-Яа-яA-Za-z]+))|([А-Яа-яA-Za-z-]+)")

func Top10(text string) []string {
	words := strings.Fields(text)
	frequencyMap := make(map[string]int)

	for _, word := range words {
		trueWord := wordValidator.Find([]byte(word))

		if trueWord != nil && (len(trueWord) == 1 && trueWord[0] != '-' || len(trueWord) != 1) {
			frequencyMap[strings.ToLower(string(trueWord))]++
		}
	}

	wordsAndFreqs := []WordAndFreq{}

	for word, frequency := range frequencyMap {
		wordsAndFreqs = append(wordsAndFreqs, WordAndFreq{word, frequency})
	}

	sort.Slice(wordsAndFreqs, func(i, j int) bool {
		if wordsAndFreqs[i].Frequency == wordsAndFreqs[j].Frequency {
			return wordsAndFreqs[i].Word < wordsAndFreqs[j].Word
		}

		return wordsAndFreqs[i].Frequency > wordsAndFreqs[j].Frequency
	})

	totalWordCount := int(math.Min(float64(len(wordsAndFreqs)), 10))
	result := make([]string, totalWordCount)

	for i := 0; i < totalWordCount; i++ {
		result[i] = wordsAndFreqs[i].Word
	}

	return result
}
