package balancer

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

type Chunks []Chunk

type Chunk struct {
	files []string
}

// TODO a lot can be done to improve error handling. Must learn how.

func MakeChunks(rootPath string, batchSize int) (Chunks, error) {
	var files []string

	if batchSize < 1 {
		return nil, fmt.Errorf("batchSize must be non-zero positive integer, %d given", batchSize)
	}

	info, err := os.Stat(rootPath)
	if err != nil {
		return nil, err
	}

	if !info.IsDir() {
		return nil, fmt.Errorf("%s is not a folder", rootPath)
	}

	err = filepath.WalkDir(
		rootPath,
		func(filePath string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}

			info, err = os.Stat(filePath)
			if err != nil {
				return err
			}

			if info.IsDir() {
				return nil
			}

			files = append(files, filePath)

			return nil
		},
	)

	if err != nil {
		return nil, err
	}

	return divide(files, batchSize), nil
}

type chunkWithCounter struct {
	Chunk
	size int
}

func divide(files []string, batchSize int) Chunks {
	chunks := Chunks{}
	lastChunk := initChunkWithCounter()

	for _, f := range files {
		if lastChunk.size >= batchSize {
			swapChunks(&chunks, &lastChunk)
		}

		lastChunk.files = append(lastChunk.files, f)
		lastChunk.size++
	}

	if lastChunk.size > 0 {
		swapChunks(&chunks, &lastChunk)
	}

	return chunks
}

func swapChunks(chunks *Chunks, currentChunk *chunkWithCounter) {
	*chunks = append(*chunks, Chunk{files: currentChunk.files})
	*currentChunk = initChunkWithCounter()
}

func initChunkWithCounter() chunkWithCounter {
	return chunkWithCounter{size: 0}
}
