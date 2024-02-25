package common

import "github.com/google/uuid"

type Chunks []Chunk

type Chunk struct {
	Files []string
	Uuid  uuid.UUID
}

func MakeChunk(files *[]string) Chunk {
	return Chunk{
		Files: *files,
		Uuid:  uuid.New(),
	}
}
