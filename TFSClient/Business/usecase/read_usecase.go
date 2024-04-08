package usecase

import (
	"context"
	"os"

	"github.com/TSF/TFSClient/Business/domain"
	"github.com/TSF/TFSClient/Business/gateway"
)

type ReadUsecase interface {
	ReadFile(context context.Context, fielname string) string
}

type readUsecase struct {
	nameNodeGateway gateway.NameNodeGateway
	dataNodeGateway gateway.DataNodeGateway
}

func NewReadUsecase(nameNodeGateway gateway.NameNodeGateway, dataNodeGateway gateway.DataNodeGateway) *readUsecase {
	return &readUsecase{
		nameNodeGateway: nameNodeGateway,
		dataNodeGateway: dataNodeGateway,
	}
}

func (r *readUsecase) ReadFile(context context.Context, fileName string) string {
	metadata, status := r.nameNodeGateway.GetMetadata(fileName)
	if status != 200 {
		return "error"
	}
	fileData := []byte{}
	for _, chunk := range metadata.Chunks {
		response := r.getChunk(fileName, context, chunk.ID, chunk.Sockets)
		if response == nil {
			return "error reading file"
		}
		fileData = append(fileData, response...)
	}
	file, err := os.Create(fileName)
	if err != nil {
		return "error creating file"
	}
	defer file.Close()

	_, err = file.Write(fileData)
	if err != nil {
		return "error creating file"
	}
	return "file downloaded"
}

func (r *readUsecase) getChunk(filename string, context context.Context, chunkId int, sockets []domain.Socket) []byte {
	for _, socket := range sockets {
		response, _ := r.dataNodeGateway.GetChunk(context, filename, chunkId, socket.ReplicaId, socket)
		if response != nil {
			return response
		}
	}
	return nil
}
