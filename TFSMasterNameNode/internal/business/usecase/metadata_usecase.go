package usecase

import (
	"github.com/TSF/TFSMasterNameNode/internal/business/domain"
	"github.com/TSF/TFSMasterNameNode/internal/business/gateway"
)

type Metadatausecase interface {
	Add(metadata *domain.Metadata) error
	Delete(fileName string) error
	Get(fileName string) (*domain.MetadataResponse, error)
}

type metadataUsecase struct {
	metadataGateway gateway.MetadataGateway
}

func NewMetadataUsecase(metadataGateway gateway.MetadataGateway) *metadataUsecase {
	return &metadataUsecase{
		metadataGateway: metadataGateway,
	}
}

func (m *metadataUsecase) Add(metadata *domain.Metadata) error {
	err := m.metadataGateway.Add(metadata)
	if err != nil {
		return err
	}
	return nil
}

func (m *metadataUsecase) Delete(fileName string) error {
	err := m.metadataGateway.Delete(fileName)
	if err != nil {
		return err
	}
	return nil
}

func (m *metadataUsecase) Get(fileName string) (*domain.MetadataResponse, error) {
	response, err := m.metadataGateway.Get(fileName)
	if err != nil {
		return nil, err
	}
	return response, nil
}
