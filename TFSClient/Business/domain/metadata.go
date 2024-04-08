package domain

type Metadata struct {
	Name           string `json:"file_name"`
	ChunksQuantity int    `json:"chunks_quantity"`
}
