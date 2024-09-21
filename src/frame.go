package src

import (
	"blog-external-sort/src/cio"
	"blog-external-sort/src/utils"
	"fmt"
	"io"
	"os"
	"slices"
)

type ChunkyFrame struct {
	id            string
	filePath      string
	processFolder string
	chunkNames    []string
	chunkSize     int
}

func NewChunkyFrame(
	filePath string,
	chunkSize int,
	processFolder string,
) (*ChunkyFrame, error) {
	id := "partition"

	frame := &ChunkyFrame{
		filePath:      filePath,
		chunkSize:     chunkSize,
		id:            id,
		processFolder: processFolder,
		chunkNames:    make([]string, 0),
	}

	err := frame.Chunker()
	if err != nil {
		return nil, err
	}
	return frame, nil
}

func (f *ChunkyFrame) Sort(sorter func(a, b map[string]interface{}) int) {
	f.ChunkSort(sorter)
	f.MergeChunks(sorter)
}

func (f *ChunkyFrame) ChunkSort(sorter func(a, b map[string]interface{}) int) {
	for _, chunkName := range f.chunkNames {
		reader, err := cio.NewJsonReader(chunkName)
		if err != nil {
			panic(err)
		}
		data, err := reader.ReadAll()
		slices.SortFunc(data, sorter)
		writer, err := cio.NewJsonWriter(chunkName + ".temp")
		writer.WriteAll(data)
	}
	f.RePackageNames()
}

func (f *ChunkyFrame) MergeChunks(sorter func(a, b map[string]interface{}) int) {
	readers := make([]*cio.JsonReader, len(f.chunkNames))
	pq := utils.NewPriorityQueue(sorter)
	for i, chunkName := range f.chunkNames {
		readers[i], _ = cio.NewJsonReader(chunkName)
		data, _ := readers[i].ReadLine()
		pq.PushItem(&utils.Item{Index: i, Value: data})
	}

	writer, _ := cio.NewJsonWriter(f.processFolder + "/id.jsonl")
	defer writer.Close()

	for pq.Len() > 0 {
		item := pq.PopItem()
		writer.WriteLine(item.Value)
		newData, err := readers[item.Index].ReadLine()
		if err != nil {
			continue
		}
		pq.Push(&utils.Item{Index: item.Index, Value: newData})
	}
}

func (f *ChunkyFrame) Chunker() error {
	reader, err := cio.NewRawReader(f.filePath)
	if err != nil {
		return err
	}

	var (
		chunk      [][]byte
		chunkNames []string
		i          = 0
		chunkCount = 0
	)

	for {
		rawByte, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		chunk = append(chunk, rawByte)
		i++

		if i == f.chunkSize {
			writeFilePath := fmt.Sprintf("%s/id_%d.jsonl", f.processFolder, chunkCount)
			if err := writeChunk(writeFilePath, chunk); err != nil {
				return err
			}
			chunkNames = append(chunkNames, writeFilePath)
			chunk = chunk[:0] // Reset chunk slice
			i = 0
			chunkCount++
		}
	}

	// Handle any remaining chunk
	if len(chunk) > 0 {
		writeFilePath := fmt.Sprintf("%s/id_%d.jsonl", f.processFolder, chunkCount)
		if err := writeChunk(writeFilePath, chunk); err != nil {
			return err
		}
		chunkNames = append(chunkNames, writeFilePath)
	}

	f.chunkNames = chunkNames
	return nil
}

func writeChunk(filePath string, chunk [][]byte) error {
	writer, err := cio.NewRawWriter(filePath)
	if err != nil {
		return err
	}
	return writer.WriteAll(chunk)
}

func (f *ChunkyFrame) RePackageNames() error {
	for _, chunkName := range f.chunkNames {
		os.Remove(chunkName)
		os.Rename(chunkName+".temp", chunkName)
	}
	return nil
}
