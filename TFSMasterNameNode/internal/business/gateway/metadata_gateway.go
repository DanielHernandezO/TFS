package gateway

import "github.com/TSF/TFSMasterNameNode/internal/business/domain"

type MetadataGateway interface {
	Add(md *domain.Metadata) error
	Delete(fileName string) error
	Get(fileName string) (*domain.MetadataResponse, error)
	GetAll() []*domain.MetadataResponse
	LocateChunk(location *domain.ChunkLocation) error
}
