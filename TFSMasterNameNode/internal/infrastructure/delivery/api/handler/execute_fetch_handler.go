package handler

import (
	"context"

	"github.com/TSF/TFSMasterNameNode/internal/business/usecase"
)

type ExecuteFetchHandler interface {
	ExecuteFetchLocations()
	ExecuteFetchSockets()
}

type executeFetchHandler struct {
	fetchUsecase usecase.FetchUsecase
}

func NewExecuteFetchHandler(fetchUsecase usecase.FetchUsecase) *executeFetchHandler {
	return &executeFetchHandler{
		fetchUsecase: fetchUsecase,
	}
}

func (f *executeFetchHandler) ExecuteFetchLocations() {
	context := context.Background()
	f.fetchUsecase.GetMetadata(&context)
}

func (f *executeFetchHandler) ExecuteFetchSockets() {
	context := context.Background()
	f.fetchUsecase.GetSockets(&context)
}
