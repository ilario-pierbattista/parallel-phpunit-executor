package balancer

import (
	"fmt"
	"sort"
	"testing"

	"github.com/mpvl/unique"
)

func TestChunck(t *testing.T) {
	t.Run("Chuncks builds the expected struct", func(t *testing.T) {

		batchSize := 3
		chunks, err := MakeChunks("./test/data/sample_root", batchSize)

		if err != nil {
			t.Fail()
		}

		onlyFilesChunks := [][]string{}
		for _, c := range chunks {
			if len(c.Files) > batchSize {
				t.Fail()
			}
			onlyFilesChunks = append(onlyFilesChunks, c.Files)
		}

		chunksString := fmt.Sprint(onlyFilesChunks)
		expectedChunksString := "[[test/data/sample_root/file_1 test/data/sample_root/file_2 test/data/sample_root/node_1/file_1_1] [test/data/sample_root/node_1/node_1_1/file_1_1_1]]"

		if chunksString != expectedChunksString {
			t.Logf(`
				Received: %s
				Expected: %s
			`, chunksString, expectedChunksString)
			t.Fail()
		}

		onlyUUIDString := []string{}
		for _, c := range chunks {
			onlyUUIDString = append(onlyUUIDString, c.Uuid.String())
		}

		sort.Strings(onlyUUIDString)
		if !(unique.StringsAreUnique(onlyUUIDString)) {
			t.Logf("UUID chunks are not unique: %v", onlyUUIDString)
			t.Fail()
		}

		if len(onlyUUIDString) != len(onlyFilesChunks) {
			t.Logf("Files chunks are %d, but UUIDs are %d", len(onlyFilesChunks), len(onlyUUIDString))
			t.Fail()
		}
	})

	t.Run("Chunks gives error on not existent directory", func(t *testing.T) {
		_, err := MakeChunks("./test/data/not_existent", 1)
		if err == nil {
			t.Fail()
		}
	})

	t.Run("Chunks gives error if path is not to a directory", func(t *testing.T) {
		_, err := MakeChunks("./test/data/sample_root/file_1", 1)
		if err == nil {
			t.Fail()
		}
	})

	t.Run("Chunks gives error if batchSize is negative or zero", func(t *testing.T) {
		_, err := MakeChunks("./test/data/sample_root", 0)
		if err == nil {
			t.Fail()
		}
	})
}
