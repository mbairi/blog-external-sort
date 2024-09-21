package cio

import (
	"github.com/bytedance/sonic"
)

type JsonReader struct {
	*RawReader
}

func NewJsonReader(filePath string) (*JsonReader, error) {
	rawReader, err := NewRawReader(filePath)
	if err != nil {
		return nil, err
	}
	return &JsonReader{RawReader: rawReader}, nil
}

func (r *JsonReader) ReadLine() (map[string]interface{}, error) {
	rawByte, err := r.RawReader.ReadLine()
	if err != nil {
		return nil, err
	}
	data := map[string]interface{}{}
	err = sonic.Unmarshal(rawByte, &data)
	return data, err
}

func (r *JsonReader) ReadAll() ([]map[string]interface{}, error) {
	rawBytes, err := r.RawReader.ReadAll()
	rawBytes = rawBytes[0 : len(rawBytes)-1]
	if err != nil {
		return nil, err
	}
	data := make([]map[string]interface{}, len(rawBytes))
	for i := range len(rawBytes) {
		sonic.Unmarshal(rawBytes[i], &data[i])
	}
	return data, err
}
