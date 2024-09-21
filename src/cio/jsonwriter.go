package cio

import (
	"github.com/bytedance/sonic"
)

type JsonWriter struct {
	*RawWriter
}

func NewJsonWriter(filepath string) (*JsonWriter, error) {
	rawWriter, err := NewRawWriter(filepath)
	if err != nil {
		return nil, err
	}
	return &JsonWriter{RawWriter: rawWriter}, nil
}

func (w *JsonWriter) WriteLine(data interface{}) error {
	asByte, err := sonic.Marshal(data)
	if err != nil {
		return err
	}
	return w.RawWriter.WriteLine(asByte)
}

func (w *JsonWriter) WriteAll(data []map[string]interface{}) error {
	rawBytes := make([][]byte, len(data))
	for i := range len(data) {
		rawBytes[i], _ = sonic.Marshal(data[i])
		rawBytes[i] = append(rawBytes[i], byte('\n'))
	}
	return w.RawWriter.WriteAll(rawBytes)
}
