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
	FetchMetadata(context context.Context, metadata *domain.Metadata) error
	DeleteMetadataFetch(fileName string) error
	GetMetadata(context *context.Context)
	GetSockets(context *context.Context)
}

type fetchUsecase struct {
	socketGateway   gateway.SocketGateway
	metadataGateway gateway.MetadataGateway
	fetchGateway    gateway.FetchGateway
}

func NewFetchUsecase(socketGateway gateway.SocketGateway, metadataGateway gateway.MetadataGateway, fetchGateway gateway.FetchGateway) *fetchUsecase {
	return &fetchUsecase{
		socketGateway:   socketGateway,
		metadataGateway: metadataGateway,
		fetchGateway:    fetchGateway,
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

func (f *fetchUsecase) FetchMetadata(context context.Context, metadata *domain.Metadata) error {
	err := f.metadataGateway.Add(metadata)
	if err != nil {
		return err
	}
	return nil
}

func (f *fetchUsecase) DeleteMetadataFetch(fileName string) error {
	err := f.metadataGateway.Delete(fileName)
	if err != nil {
		return err
	}
	return nil
}

func (f *fetchUsecase) GetMetadata(context *context.Context) {
	metadata := f.fetchGateway.GetMetadata(context)
	for _, data := range metadata.Files {
		f.metadataGateway.Add(&domain.Metadata{
			Name:           data.FileId,
			ChunksQuantity: len(data.Chunks),
		})
		for _, chunk := range data.Chunks {
			for _, socket := range chunk.Sockets {
				f.metadataGateway.LocateChunk(&domain.ChunkLocation{
					Name:      data.FileId,
					ID:        int(chunk.ChunkId),
					ReplicaId: int(socket.ReplicaId),
					Socket: domain.Socket{
						Ip:        socket.Ip,
						Port:      socket.Port,
						ReplicaId: int(socket.ReplicaId),
					},
				})
			}
		}
	}
}

func (f *fetchUsecase) GetSockets(context *context.Context) {
	sockets := f.fetchGateway.GetSockets(context)
	for _, socket := range sockets.Sockets {
		f.socketGateway.Add(&domain.Socket{
			Ip:   socket.Ip,
			Port: socket.Port,
		})
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
