package Business

import (
	"math"
	"os"

	"github.com/TSF/TFSClient/Commons"
)

func PartitionFile(file *os.File, chunkSize int) ([][]byte, error) {
	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}
	fileSize := stat.Size()

	size := int(math.Ceil(float64(fileSize) / float64(chunkSize)))

	partitions := make([][]byte, chunkSize)

	buffer := make([]byte, fileSize)
	_, err = file.Read(buffer)
	if err != nil {
		return nil, err
	}

	for i := 0; i < chunkSize; i++ {
		start := i * size
		end := min(int(fileSize), (i+1)*size)
		partitions[i] = buffer[start:end]
	}

	return partitions, nil
}

func RestoreFile(filename string) error {

	var request Commons.DataNodeRequest
	request.Filename = filename
	files, err := GetMetadata(filename)

	var colum int
	var partitions [][]byte

	for _, file := range files.Chunks {
		request.BlockID = file.ChunkId
		resp := GetBlocks(request)
		for i := 0; i > len(resp.Block); i++ {
			//partitions[colum][i] := resp.Block[i]
		}
		colum++
	}

	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}

	for _, partition := range partitions {
		_, err := file.Write(partition)
		if err != nil {
			file.Close()
			return err
		}
	}

	file.Close()

	return nil
}
