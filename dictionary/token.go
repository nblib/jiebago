package dictionary

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// Token represents a Chinese word with (optional) frequency and POS.
type Token struct {
	text      string
	frequency float64
	pos       string
}

//Text returns token's text.
func (t Token) Text() string {
	return t.text
}

// Frequency returns token's frequency.
func (t Token) Frequency() float64 {
	return t.frequency
}

// Pos returns token's POS.
func (t Token) Pos() string {
	return t.pos
}

// NewToken creates a new token.
func NewToken(text string, frequency float64, pos string) *Token {
	return &Token{text: text, frequency: frequency, pos: pos}
}

type TokenReader struct {
	scanner *bufio.Scanner
	token   *Token
	err     error
}

func (o *TokenReader) HasNext() bool {
	scanner := o.scanner
	token := &Token{}
	var err error
	if scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, " ")
		token.text = strings.TrimSpace(strings.Replace(fields[0], "\ufeff", "", 1))
		if length := len(fields); length > 1 {
			token.frequency, err = strconv.ParseFloat(fields[1], 64)
			if err != nil {
				o.err = err
				return false
			}
			if length > 2 {
				token.pos = strings.TrimSpace(fields[2])
			}
		}
		o.token = token
		return true
	}
	return false
}
func (o *TokenReader) Next() *Token {
	return o.token
}
func (o *TokenReader) Err() error {
	return o.err
}
func NewTokenReader(file *os.File) *TokenReader {
	scanner := bufio.NewScanner(file)
	return &TokenReader{scanner: scanner}
}
