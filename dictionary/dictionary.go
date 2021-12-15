// Package dictionary contains a interface and wraps all io related work.
// It is used by jiebago module to read/write files.
package dictionary

import (
	"os"
	"path/filepath"
)

// DictLoader is the interface that could add one token or load
// tokens from channel.
type DictLoader interface {
	Load(*TokenReader) error
	AddToken(*Token)
}

// LoadDictionary reads the given file and passes all tokens to a DictLoader.
func LoadDictionary(dl DictLoader, fileName string) error {
	filePath, err := dictPath(fileName)
	if err != nil {
		return err
	}
	dictFile, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer dictFile.Close()
	reader := NewTokenReader(dictFile)

	return dl.Load(reader)
}

func dictPath(dictFileName string) (string, error) {
	if filepath.IsAbs(dictFileName) {
		return dictFileName, nil
	}
	var dictFilePath string
	cwd, err := os.Getwd()
	if err != nil {
		return dictFilePath, err
	}
	dictFilePath = filepath.Clean(filepath.Join(cwd, dictFileName))
	return dictFilePath, nil
}
