package usecase

import (
	"io"
	"os"
)

func Partition(file *os.File, blockSize int) ([][]byte, string) {
	var blocks [][]byte
	buffer := make([]byte, blockSize)
	for {
		bytesRead, err := file.Read(buffer)
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, ""
		}

		block := make([]byte, bytesRead)
		copy(block, buffer[:bytesRead])
		blocks = append(blocks, block)
	}

	return blocks, ""
}
