package Commons

type Chunk struct {
	ChunkId int32  `json:"chunk_id"`
	Sockets Socket `json:"sockets"`
}
