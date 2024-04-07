package usecase

import (
	"context"

	"github.com/TSF/TFSMasterNameNode/internal/business/domain"
	"github.com/TSF/TFSMasterNameNode/internal/business/gateway"
)

type Metadatausecase interface {
	Add(context context.Context, metadata *domain.Metadata) error
	Delete(context context.Context, fileName string) error
	Get(fileName string) (*domain.MetadataResponse, error)
}

type metadataUsecase struct {
	metadataGateway gateway.MetadataGateway
	fetchGatewat    gateway.FetchGateway
}

func NewMetadataUsecase(metadataGateway gateway.MetadataGateway, fetchGatewat gateway.FetchGateway) *metadataUsecase {
	return &metadataUsecase{
		metadataGateway: metadataGateway,
		fetchGatewat:    fetchGatewat,
	}
}

func (m *metadataUsecase) Add(context context.Context, metadata *domain.Metadata) error {
	err := m.metadataGateway.Add(metadata)
	if err != nil {
		return err
	}
	_, err = m.fetchGatewat.FetchMetadata(&context, metadata)
	if err != nil {
		return err
	}
	return nil
}

func (m *metadataUsecase) Delete(context context.Context, fileName string) error {
	err := m.metadataGateway.Delete(fileName)
	if err != nil {
		return err
	}
	_, err = m.fetchGatewat.DeleteMetadataFetch(&context, fileName)
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
