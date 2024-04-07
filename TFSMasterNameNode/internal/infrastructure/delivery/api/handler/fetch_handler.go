package handler

import (
	"context"

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

func (f *fetchHandler) ExecuteFetchLocations() {

}
