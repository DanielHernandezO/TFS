package Business

import (
	"math"
	"os"

	"github.com/TSF/TFSClient/Commons"
)

func ParticionFile(file *os.File, size int) ([][]byte, error) {

	var request Commons.NameNodeRequest

	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}
	fileSize := stat.Size()

	numPartitions := int(math.Ceil(float64(fileSize) / float64(size)))

	partitions := make([][]byte, numPartitions)

	request.File_name = file.Name()
	request.ChunksQuantity = int(len(partitions))
	resp, err := WritteAlert(request)

	if resp.StatusCode != 200 {
		return nil, err
	}
	if err != nil {
		return nil, err
	}

	buffer := make([]byte, fileSize)
	_, err = file.Read(buffer)
	if err != nil {
		return nil, err
	}

	for i := 0; i < numPartitions; i++ {
		start := i * size
		end := min(int(fileSize), (i+1)*size)
		partitions[i] = buffer[start:end]
	}

	return partitions, nil
}

func RestoreFile(filename string) error {

	var request Commons.DataNodeReadRequest
	request.Request.Name = filename
	files, err := GetMetadata(filename)

	var colum int
	var partitions [][]byte

	for _, file := range files.Chunks {
		request.Request.ChunkId = file.ChunkId
		request.Request.ReplicaId = file.Sockets.Replica_id
		//resp := GetBlocks(request)

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
