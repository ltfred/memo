package parser

import (
	"encoding/json"
	"github.com/ltfred/memo/utils"
	"os"
)

type Parser struct{}

func (pa *Parser) Add(data Memo) error {
	memos, err := pa.read()
	if err != nil {
		return err
	}

	memos = append(memos, data)
	return pa.write(memos)
}

func (pa *Parser) Show() ([]Memo, error) {
	return pa.read()
}

func (pa *Parser) read() ([]Memo, error) {
	var data []Memo
	path := utils.GetFilePath()
	if utils.FileIsExist(path) {
		file, err := os.OpenFile(path, os.O_RDWR, 0666)
		if err != nil {
			return nil, err
		}
		defer func(file *os.File) {
			_ = file.Close()
		}(file)

		decoder := json.NewDecoder(file)
		err = decoder.Decode(&data)
		if err != nil {
			return nil, err
		}
	}

	return data, nil
}

func (pa *Parser) write(data []Memo) (err error) {
	path := utils.GetFilePath()
	var file *os.File
	if utils.FileIsExist(path) {
		file, err = os.OpenFile(path, os.O_RDWR, 0666)
	} else {
		file, err = os.Create(path)
	}
	if err != nil {
		return err
	}

	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	encoder := json.NewEncoder(file)

	return encoder.Encode(data)
}
