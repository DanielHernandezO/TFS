package Commons

type DataNodeResponse struct {
	Code          int32
	Message       string
	ChunkMetadata ChunkMetadata
	ChunkBytes    byte
}

type ChunkWriteRequest struct {
	ChunkMetadata ChunkMetadata
	Socket        Socket
	ChunkBytes    byte
}

type ChunkWriteResponse struct {
	Code          int32
	Message       string
	ChunkMetadata ChunkMetadata
}

type ChunkMetadata struct {
	Name      string
	ChunkId   int32
	ReplicaId int32
}

type DataNodeReadRequest struct {
	Request ChunkMetadata
}
