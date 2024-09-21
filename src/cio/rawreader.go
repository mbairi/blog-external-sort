package cio

import (
	"bufio"
	"bytes"
	"os"
)

type RawReader struct {
	filePath string
	file     *os.File
	cursor   *bufio.Reader
}

func NewRawReader(filePath string) (*RawReader, error) {
	file, err := os.OpenFile(filePath, os.O_RDONLY, os.ModeTemporary)
	if err != nil {
		return nil, err
	}

	cursor := bufio.NewReader(file)
	return &RawReader{filePath: filePath, file: file, cursor: cursor}, nil
}

func (r *RawReader) ReadLine() ([]byte, error) {
	return r.cursor.ReadBytes(byte('\n'))
}

func (r *RawReader) ReadAll() ([][]byte, error) {
	rawBytes, err := os.ReadFile(r.filePath)
	if err != nil {
		return nil, err
	}
	rawBytesSplit := bytes.Split(rawBytes, []byte("\n"))
	return rawBytesSplit, nil
}
