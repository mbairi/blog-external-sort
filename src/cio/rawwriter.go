package cio

import (
	"bufio"
	"bytes"
	"os"
)

type RawWriter struct {
	filePath string
	file     *os.File
	cursor   *bufio.Writer
}

func NewRawWriter(filePath string) (*RawWriter, error) {
	file, err := os.Create(filePath)
	if err != nil {
		return nil, err
	}

	cursor := bufio.NewWriter(file)
	return &RawWriter{filePath: filePath, file: file, cursor: cursor}, nil
}

func (r *RawWriter) WriteLine(rawBytes []byte) error {
	rawBytes = append(rawBytes, '\n')
	_, err := r.cursor.Write(rawBytes)
	return err
}

func (r *RawWriter) WriteAll(lines [][]byte) error {
	var buffer bytes.Buffer
	for _, line := range lines {
		buffer.Write(line)
	}
	return os.WriteFile(r.filePath, buffer.Bytes(), os.ModeAppend)
}

func (r *RawWriter) Close() {
	r.file.Sync()
	r.cursor.Flush()
}
