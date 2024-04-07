package handler

import (
	"net/http"

	"github.com/TSF/TFSMasterNameNode/internal/business/constant"
	"github.com/TSF/TFSMasterNameNode/internal/business/domain"
	"github.com/TSF/TFSMasterNameNode/internal/business/usecase"
	"github.com/gin-gonic/gin"
)

type MetadataHandler interface {
	Add(c *gin.Context)
	Delete(c *gin.Context)
	Get(c *gin.Context)
}

type metadataHandler struct {
	metadataUsecase usecase.Metadatausecase
}

func NewMetadataHanlder(metadataUsecase usecase.Metadatausecase) *metadataHandler {
	return &metadataHandler{
		metadataUsecase: metadataUsecase,
	}
}

func (m *metadataHandler) Add(c *gin.Context) {
	var metadata domain.Metadata

	if err := c.BindJSON(&metadata); err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{
			Code:    http.StatusBadRequest,
			Message: constant.InvalidBody,
		})
		return
	}
	err := m.metadataUsecase.Add(&metadata)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, domain.Response{
		Code:    http.StatusOK,
		Message: constant.SavedMetadata,
	})
}

func (m *metadataHandler) Delete(c *gin.Context) {
	fileName := c.Param(constant.FileName)

	err := m.metadataUsecase.Delete(fileName)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, domain.Response{
		Code:    http.StatusOK,
		Message: constant.DeletedMetadata,
	})
}

func (m *metadataHandler) Get(c *gin.Context) {
	fileName := c.Param(constant.FileName)

	metadata, err := m.metadataUsecase.Get(fileName)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, metadata)
}
