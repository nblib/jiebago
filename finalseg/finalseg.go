// Package finalseg is the Golang implementation of Jieba's finalseg module.
package finalseg

import (
	"regexp"
)

var (
	reHan  = regexp.MustCompile(`\p{Han}+`)
	reSkip = regexp.MustCompile(`(\d+\.\d+|[a-zA-Z0-9]+)`)
)

func cutHanSync(sentence string) []string {
	result := make([]string, 0)

	runes := []rune(sentence)
	_, posList := viterbiNew(runes)
	begin, next := 0, 0
	for i, char := range runes {
		pos := posList[i]
		switch pos {
		case State_B:
			begin = i
		case State_E:
			result = append(result, string(runes[begin:i+1]))
			next = i + 1
		case State_S:
			result = append(result, string(char))
			next = i + 1
		}
	}
	if next < len(runes) {
		result = append(result, string(runes[next:]))
	}

	return result
}

func CutSync(sentence string) []string {
	result := make([]string, 0)
	s := sentence
	var hans string
	var hanLoc []int
	var nonhanLoc []int

	for {
		hanLoc = reHan.FindStringIndex(s)
		if hanLoc == nil {
			if len(s) == 0 {
				break
			}
		} else if hanLoc[0] == 0 {
			hans = s[hanLoc[0]:hanLoc[1]]
			s = s[hanLoc[1]:]
			//for han := range cutHan(hans) {
			//	result = append(result, han)
			//}
			result = append(result, cutHanSync(hans)...)
			continue
		}
		nonhanLoc = reSkip.FindStringIndex(s)
		if nonhanLoc == nil {
			if len(s) == 0 {
				break
			}
		} else if nonhanLoc[0] == 0 {
			nonhans := s[nonhanLoc[0]:nonhanLoc[1]]
			s = s[nonhanLoc[1]:]
			if nonhans != "" {
				result = append(result, nonhans)
				continue
			}
		}
		var loc []int
		if hanLoc == nil && nonhanLoc == nil {
			if len(s) > 0 {
				result = append(result, s)
				break
			}
		} else if hanLoc == nil {
			loc = nonhanLoc
		} else if nonhanLoc == nil {
			loc = hanLoc
		} else if hanLoc[0] < nonhanLoc[0] {
			loc = hanLoc
		} else {
			loc = nonhanLoc
		}
		result = append(result, s[:loc[0]])
		s = s[loc[0]:]
	}

	return result
}
