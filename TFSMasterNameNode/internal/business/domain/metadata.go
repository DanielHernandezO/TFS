package domain

type Metadata struct {
	Name           string `json:"file_name"`
	ChunksQuantity int    `json:"chunks_quantity"`
}

type MetadataResponse struct {
	Name   string  `json:"file_name"`
	Chunks []Chunk `json:"chunks"`
}

type Chunk struct {
	ID      int      `json:"chunk_id"`
	Sockets []Socket `json:"sockets"`
}
