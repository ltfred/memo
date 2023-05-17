package parser

import (
	"encoding/json"
	"errors"
	"github.com/ltfred/memo/utils"
	"os"
	"strconv"
)

type Parser struct{}

func (pa *Parser) Add(data Memo) error {
	memosMap, err := pa.read()
	if err != nil {
		return err
	}

	memosMap[strconv.FormatInt(data.CreateAt, 10)] = data
	return pa.write(memosMap)
}

func (pa *Parser) Show() (map[string]Memo, error) {
	return pa.read()
}

func (pa *Parser) Delete(uuid string) error {
	memosMap, err := pa.read()
	if err != nil {
		return err
	}
	delete(memosMap, uuid)

	return pa.write(memosMap)
}

func (pa *Parser) Modify(uuid string, data Memo) error {
	memosMap, err := pa.read()
	if err != nil {
		return err
	}
	memosMap[uuid] = data

	return pa.write(memosMap)
}

func (pa *Parser) GetRecord(uuid string) (Memo, error) {
	memosMap, err := pa.read()
	if err != nil {
		return Memo{}, err
	}
	if _, ok := memosMap[uuid]; !ok {
		return Memo{}, errors.New("record not exist")
	}

	return memosMap[uuid], nil
}

func (pa *Parser) read() (map[string]Memo, error) {
	data := make(map[string]Memo, 0)
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

func (pa *Parser) write(data map[string]Memo) (err error) {
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
