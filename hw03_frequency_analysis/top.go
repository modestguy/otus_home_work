package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type wordFrequencyStruct struct {
	Word      string
	Frequency int
}

func Top10(input string) []string {
	mapOfWords := make(map[string]*wordFrequencyStruct)
	s := strings.Fields(input)
	for _, word := range s {
		if _, ok := mapOfWords[word]; ok {
			mapOfWords[word].Frequency++
		} else {
			mapOfWords[word] = &wordFrequencyStruct{Word: word, Frequency: 1}
		}
	}

	wf := getWordFrequencyMapAsSlice(mapOfWords)
	sort.Slice(wf, func(i, j int) bool {
		if wf[i].Frequency == wf[j].Frequency {
			return strings.Compare(wf[i].Word, wf[j].Word) == -1
		}

		return wf[i].Frequency > wf[j].Frequency
	})

	return getFirstNWords(wf, 10)
}

func getWordFrequencyMapAsSlice(wfMap map[string]*wordFrequencyStruct) []wordFrequencyStruct {
	wfSlice := make([]wordFrequencyStruct, 0, len(wfMap))
	for _, item := range wfMap {
		wfSlice = append(wfSlice, *item)
	}
	return wfSlice
}

func getFirstNWords(wf []wordFrequencyStruct, count int) []string {
	res := make([]string, 0, count)
	for index, value := range wf {
		if index >= count {
			break
		}
		res = append(res, value.Word)
	}
	return res
}
