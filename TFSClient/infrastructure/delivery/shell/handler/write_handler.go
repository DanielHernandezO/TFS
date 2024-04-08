package handler

import (
	"context"

	"github.com/TSF/TFSClient/Business/usecase"
)

type WriteHandler interface {
	WriteFile(location string) string
}

type writeHandler struct {
	writeUsecase usecase.WriteUsecase
}

func NewWriteHandler(writeUsecase usecase.WriteUsecase) *writeHandler {
	return &writeHandler{
		writeUsecase: writeUsecase,
	}
}

func (w *writeHandler) WriteFile(location string) string {
	context := context.Background()
	response := w.writeUsecase.WriteFile(context, location)
	return response
}
