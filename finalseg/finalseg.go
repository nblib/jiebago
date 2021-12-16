// Package finalseg is the Golang implementation of Jieba's finalseg module.
package finalseg

import (
	"github.com/wangbin/jiebago/util"
	"regexp"
)

var (
	reHan  = regexp.MustCompile(`\p{Han}+`)
	reSkip = regexp.MustCompile(`(\d+\.\d+|[a-zA-Z0-9]+)`)
)

func cutHanSync(sentence string, result *util.StrArrBuffer) {

	runes := []rune(sentence)
	_, posList := viterbiNew(runes)
	begin, next := 0, 0
	for i, char := range runes {
		pos := posList[i]
		switch pos {
		case State_B:
			begin = i
		case State_E:
			result.Write(string(runes[begin : i+1]))
			next = i + 1
		case State_S:
			result.Write(string(char))
			next = i + 1
		}
	}
	if next < len(runes) {
		result.Write(string(runes[next:]))
	}
}

// CutSync 分割字符串
// 分3中情况:
// 1. 字符串汉字开头,使用hmm分割开头的汉字部分,剩余部分重新判断开头
// 2. 字符串(\d+\.\d+|[a-zA-Z0-9]+)开头, 直接提取出来,剩余部分重新判断开头
// 3. 字符串开头不满足上述两个条件,将不满足的部分直接提取出来,剩余部分重新判断开头
func CutSync(sentence string, result * util.StrArrBuffer) {
	s := sentence
	var hans string
	var hanLoc []int
	var nonhanLoc []int

	for {
		//情况1
		hanLoc = reHan.FindStringIndex(s)
		if hanLoc == nil {
			if len(s) == 0 {
				break
			}
		} else if hanLoc[0] == 0 {
			hans = s[hanLoc[0]:hanLoc[1]]
			s = s[hanLoc[1]:]
			cutHanSync(hans, result)
			continue
		}
		//情况2
		nonhanLoc = reSkip.FindStringIndex(s)
		if nonhanLoc == nil {
			if len(s) == 0 {
				break
			}
		} else if nonhanLoc[0] == 0 {
			nonhans := s[nonhanLoc[0]:nonhanLoc[1]]
			s = s[nonhanLoc[1]:]
			if nonhans != "" {
				result.Write(nonhans)
				continue
			}
		}
		//情况3
		var loc []int
		if hanLoc == nil && nonhanLoc == nil {
			if len(s) > 0 {
				result.Write(s)
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
		result.Write(s[:loc[0]])
		s = s[loc[0]:]
	}
}
