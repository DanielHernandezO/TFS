package usecase

import (
	"context"

	"github.com/TSF/TFSMasterNameNode/internal/business/domain"
	"github.com/TSF/TFSMasterNameNode/internal/business/gateway"
	"github.com/TSF/TFSMasterNameNode/internal/infrastructure/config"
)

type FetchUsecase interface {
	FetchSocket(context context.Context) *config.Sockets
	FetchLocations(context context.Context) *config.Metadata
}

type fetchUsecase struct {
	socketGateway   gateway.SocketGateway
	metadataGateway gateway.MetadataGateway
}

func NewFetchUsecase(socketGateway gateway.SocketGateway, metadataGateway gateway.MetadataGateway) *fetchUsecase {
	return &fetchUsecase{
		socketGateway:   socketGateway,
		metadataGateway: metadataGateway,
	}
}

func (f *fetchUsecase) FetchSocket(context context.Context) *config.Sockets {
	socketsData := f.socketGateway.Get()

	sockets := []*config.Socket{}
	for _, value := range socketsData {
		sockets = append(sockets, &config.Socket{Ip: value.Ip, Port: value.Port})
	}

	return &config.Sockets{
		Sockets: sockets,
	}
}

func (f *fetchUsecase) FetchLocations(context context.Context) *config.Metadata {
	metadata := f.metadataGateway.GetAll()
	files := []*config.FileMetadata{}
	for i := 0; i < len(metadata); i++ {
		files = append(files, &config.FileMetadata{
			FileId: metadata[i].Name,
			Chunks: f.buildFileChunks(metadata[i].Chunks),
		})
	}
	return &config.Metadata{
		Files: files,
	}
}

func (f *fetchUsecase) buildFileChunks(socketsData []domain.Chunk) []*config.Chunk {
	chunks := []*config.Chunk{}
	for _, value := range socketsData {
		chunks = append(chunks, &config.Chunk{
			ChunkId: int64(value.ID),
			Sockets: f.buildSockets(value.Sockets),
		})
	}
	return chunks
}

func (f *fetchUsecase) buildSockets(socketsData []domain.Socket) []*config.Socket {
	chunks := []*config.Socket{}
	for _, value := range socketsData {
		chunks = append(chunks, &config.Socket{
			Ip:        value.Ip,
			Port:      value.Port,
			ReplicaId: int32(value.ReplicaId),
		})
	}
	return chunks
}
