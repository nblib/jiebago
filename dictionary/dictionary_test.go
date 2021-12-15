package dictionary

import (
	"sync"
	"testing"
)

type Dict struct {
	freqMap map[string]float64
	posMap  map[string]string
	sync.RWMutex
}

func (d *Dict) Load(reader *TokenReader) error {
	d.Lock()
	defer d.Unlock()

	for reader.HasNext() {
		token := reader.Next()
		d.freqMap[token.Text()] = token.Frequency()
		if len(token.Pos()) > 0 {
			d.posMap[token.Text()] = token.Pos()
		}
	}
	return reader.Err()
}

func (d *Dict) AddToken(token *Token) {
	d.Lock()
	defer d.Unlock()
	d.freqMap[token.Text()] = token.Frequency()
	if len(token.Pos()) > 0 {
		d.posMap[token.Text()] = token.Pos()
	}
}

func TestLoadDictionary(t *testing.T) {
	d := &Dict{freqMap: make(map[string]float64), posMap: make(map[string]string)}
	err := LoadDictionary(d, "../dict/jieba.dict.utf8")
	if err != nil {
		t.Fatalf(err.Error())
	}
	if len(d.freqMap) != 348982 {
		t.Fatalf("Failed to load userdict.txt, got %d tokens with frequency, expected 7",
			len(d.freqMap))
	}
	if len(d.posMap) != 348982 {
		t.Fatalf("Failed to load userdict.txt, got %d tokens with pos, expected 6", len(d.posMap))
	}
}

/**
old benchmark:
	goos: windows
	goarch: amd64
	pkg: jiebago/dictionary
	cpu: Intel(R) Core(TM) i7-8650U CPU @ 1.90GHz
	BenchmarkLoadDictionary
	BenchmarkLoadDictionary-8   	       2	 530282250 ns/op
	PASS
new benchmark:
	goos: windows
	goarch: amd64
	pkg: jiebago/dictionary
	cpu: Intel(R) Core(TM) i7-8650U CPU @ 1.90GHz
	BenchmarkLoadDictionary
	BenchmarkLoadDictionary-8   	       4	 271067675 ns/op
	PASS

promote 48.88%
*/
func BenchmarkLoadDictionary(b *testing.B) {
	for i := 0; i < b.N; i++ {
		d := &Dict{freqMap: make(map[string]float64), posMap: make(map[string]string)}
		err := LoadDictionary(d, "../dict/jieba.dict.utf8")
		if err != nil {
			b.Fatalf(err.Error())
		}
	}
}

func TestAddToken(t *testing.T) {
	d := &Dict{freqMap: make(map[string]float64), posMap: make(map[string]string)}
	LoadDictionary(d, "../dict/jieba.dict.utf8")
	d.AddToken(&Token{"好用", 99, "a"})
	if d.freqMap["好用"] != 99 {
		t.Fatalf("Failed to add token, got frequency %f, expected 99", d.freqMap["好用"])
	}
	if d.posMap["好用"] != "a" {
		t.Fatalf("Failed to add token, got pos %s, expected \"a\"", d.posMap["好用"])
	}
}
