package api

import "github.com/gin-gonic/gin"

func mapUrls(router *gin.Engine, dependencies *Dependencies) {
	groupPath := router.Group("/namenode")
	groupPath.GET("/ping", dependencies.pingHandler.Ping)
	groupPath.GET("/get/metadata/:file_name", dependencies.metadataHandler.Get)
	groupPath.POST("/add/metadata", dependencies.metadataHandler.Add)
	groupPath.DELETE("/delete/:file_name", dependencies.metadataHandler.Delete)
	groupPath.GET("/get/replicas", dependencies.socketHandler.GetReplicas)
}
