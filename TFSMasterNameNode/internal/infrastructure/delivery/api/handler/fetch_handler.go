package handler

import (
	"context"
	"net/http"

	"github.com/TSF/TFSMasterNameNode/internal/business/constant"
	"github.com/TSF/TFSMasterNameNode/internal/business/domain"
	"github.com/TSF/TFSMasterNameNode/internal/business/usecase"
	"github.com/TSF/TFSMasterNameNode/internal/infrastructure/config"
)

type fetchHandler struct {
	fetchUsecase usecase.FetchUsecase
	config.UnimplementedFetchServer
}

func NewFetchHandler(fetchUsecase usecase.FetchUsecase) *fetchHandler {
	return &fetchHandler{
		fetchUsecase: fetchUsecase,
	}
}

func (f *fetchHandler) FetchSocket(context context.Context, request *config.VoidRequest) (*config.Sockets, error) {
	sockets := f.fetchUsecase.FetchSocket(context)
	return sockets, nil
}

func (f *fetchHandler) FetchLocations(context context.Context, request *config.VoidRequest) (*config.Metadata, error) {
	locations := f.fetchUsecase.FetchLocations(context)
	return locations, nil
}

func (f *fetchHandler) FetchMetadata(context context.Context, metadata *config.FileMetadataFetch) (*config.Response, error) {
	err := f.fetchUsecase.FetchMetadata(context, &domain.Metadata{
		Name:           metadata.FileId,
		ChunksQuantity: int(metadata.ChunksQuantity),
	})
	if err != nil {
		return nil, err
	}
	return &config.Response{
		Code:    http.StatusOK,
		Message: constant.SavedMetadata,
	}, nil
}

func (f *fetchHandler) DeleteMetadataFetch(context context.Context, metadata *config.DeleteMetadata) (*config.Response, error) {
	err := f.fetchUsecase.DeleteMetadataFetch(metadata.FileId)
	if err != nil {
		return nil, err
	}
	return &config.Response{
		Code:    http.StatusOK,
		Message: constant.SavedMetadata,
	}, nil
}
