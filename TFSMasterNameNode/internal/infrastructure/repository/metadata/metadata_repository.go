package metadata

import (
	"fmt"
	"sync"

	"github.com/TSF/TFSMasterNameNode/internal/business/constant"
	"github.com/TSF/TFSMasterNameNode/internal/business/domain"
)

type metadataRepository struct {
	metadata map[string]map[int]map[string]string
	mu       sync.Mutex
}

func NewMetadataRepository() *metadataRepository {
	return &metadataRepository{
		metadata: map[string]map[int]map[string]string{},
	}
}

func (m *metadataRepository) Add(md *domain.Metadata) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	_, exist := m.metadata[md.Name]
	if exist {
		return fmt.Errorf(constant.FileWasFound)
	}

	m.metadata[md.Name] = map[int]map[string]string{}

	for i := 0; i < md.ChunksQuantity; i++ {
		m.metadata[md.Name][i] = map[string]string{}
	}
	return nil
}

func (m *metadataRepository) LocateChunk(location *domain.ChunkLocation) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	_, exist := m.metadata[location.Name]
	if !exist {
		return fmt.Errorf(constant.FileNotFound)
	}
	_, exist = m.metadata[location.Name][location.ID]
	if !exist {
		return fmt.Errorf(constant.ChunkNotFound)
	}

	m.metadata[location.Name][location.ID][location.Socket.Ip] = location.Socket.Port

	return nil
}

func (m *metadataRepository) Delete(fileName string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	_, exist := m.metadata[fileName]
	if !exist {
		return fmt.Errorf(constant.FileNotFound)
	}

	delete(m.metadata, fileName)

	return nil
}

func (m *metadataRepository) Get(fileName string) (*domain.MetadataResponse, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	_, exist := m.metadata[fileName]
	if !exist {
		return nil, fmt.Errorf(constant.FileNotFound)
	}

	for _, value := range m.metadata[fileName] {
		if len(value) == 0 {
			return nil, fmt.Errorf(constant.FileNotAvailable)
		}
	}

	return m.buildMetadata(fileName), nil
}

func (m *metadataRepository) GetAll() []*domain.MetadataResponse {
	metadata := []*domain.MetadataResponse{}
	for key := range m.metadata {
		metadata = append(metadata, m.buildMetadata(key))
	}
	return metadata
}

func (m *metadataRepository) buildMetadata(fileName string) *domain.MetadataResponse {
	metadata := &domain.MetadataResponse{
		Name:   fileName,
		Chunks: m.buildChunks(fileName),
	}
	return metadata
}

func (m *metadataRepository) buildChunks(fileName string) []domain.Chunk {
	chunks := []domain.Chunk{}
	for key, value := range m.metadata[fileName] {
		chunk := domain.Chunk{
			ID:      key,
			Sockets: m.buildSocktes(value),
		}
		chunks = append(chunks, chunk)
	}
	return chunks
}

func (m *metadataRepository) buildSocktes(machines map[string]string) []domain.Socket {
	sockets := []domain.Socket{}
	for key, value := range machines {
		socket := domain.Socket{
			Ip:   key,
			Port: value,
		}
		sockets = append(sockets, socket)
	}
	return sockets
}
