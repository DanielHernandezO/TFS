package handler

import (
	"net/http"

	"github.com/TSF/TFSMasterNameNode/internal/business/constant"
	"github.com/TSF/TFSMasterNameNode/internal/business/domain"
	"github.com/TSF/TFSMasterNameNode/internal/business/usecase"
	"github.com/gin-gonic/gin"
)

type SocketHandler interface {
	GetReplicas(c *gin.Context)
	Update()
}

type socketHandler struct {
	socketUsecase usecase.SokcetUsecase
}

func NewSocketHandler(socketUsecase usecase.SokcetUsecase) *socketHandler {
	return &socketHandler{
		socketUsecase: socketUsecase,
	}
}

func (s *socketHandler) GetReplicas(c *gin.Context) {
	replicas := s.socketUsecase.GetReplicas()
	if len(*replicas) == 0 {
		c.JSON(http.StatusNotFound, &domain.Response{
			Code:    http.StatusNotFound,
			Message: constant.ReplicaNotFound,
		})
		return
	}
	c.JSON(http.StatusOK, replicas)
}

func (s *socketHandler) Update() {
	s.socketUsecase.Update()
}
