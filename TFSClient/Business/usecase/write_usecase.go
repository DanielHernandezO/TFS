package usecase

import (
	"context"
	"fmt"
	"os"
	"path"

	"github.com/TSF/TFSClient/Business/gateway"
)

type WriteUsecase interface {
	WriteFile(context context.Context, location string) string
}

type writeUsecase struct {
	nameNodeGateway gateway.NameNodeGateway
	dataNodeGateway gateway.DataNodeGateway
}

func NewWriteUsecase(nameNodeGateway gateway.NameNodeGateway, dataNodeGateway gateway.DataNodeGateway) *writeUsecase {
	return &writeUsecase{
		nameNodeGateway: nameNodeGateway,
		dataNodeGateway: dataNodeGateway,
	}
}

func (w *writeUsecase) WriteFile(context context.Context, location string) string {
	fileData, err := os.Stat(location)

	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("El archivo no existe")
		} else {
			fmt.Println("Error al obtener información del archivo:", err)
		}
		return "Error getting file"
	}
	if fileData.IsDir() {
		fmt.Println("La ruta es un directorio")
		return "Error getting file"
	}

	file, err := os.Open(location)
	if err != nil {
		return "Error getting file"
	}

	bufferSize := 128 * 1024
	partitions, errPartition := Partition(file, bufferSize)
	if partitions == nil {
		return errPartition
	}
	filename := path.Base(location)
	status := w.nameNodeGateway.SendMetadata(location, len(partitions))
	if status != 200 {
		return "Error saving metadata"
	}

	for i, partition := range partitions {
		retry := 3
		for j := 0; j <= retry; j++ {
			replicas, status := w.nameNodeGateway.GetReplicas()
			if status != 200 {
				w.nameNodeGateway.DeleteMetadata(filename)
				return "Error saving metadata"
			}
			err := w.dataNodeGateway.SendChunk(context, filename, i, replicas, partition)
			if err != nil && j >= retry {
				w.nameNodeGateway.DeleteMetadata(filename)
				return "Error saving metadata"
			}
			if err == nil {
				break
			}
		}
	}
	return "ok"
}
