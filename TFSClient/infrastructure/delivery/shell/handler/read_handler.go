package handler

import (
	"context"

	"github.com/TSF/TFSClient/Business/usecase"
)

type ReadHandler interface {
	ReadFile(filename string) string
}

type readHandler struct {
	readUsecase usecase.ReadUsecase
}

func NewReadHandler(readUsecase usecase.ReadUsecase) *readHandler {
	return &readHandler{
		readUsecase: readUsecase,
	}
}

func (r *readHandler) ReadFile(filename string) string {
	context := context.Background()
	response := r.readUsecase.ReadFile(context, filename)
	return response
}
