package balancer

import (
	"fmt"
	"testing"
)

func TestChunck(t *testing.T) {
	t.Run("Chuncks builds the expected struct", func(t *testing.T) {
		c, err := MakeChunks("./test/data/sample_root", 3)

		if err != nil {
			t.Fail()
		}

		chunksString := fmt.Sprint(c)
		expectedChunksString := "[{[test/data/sample_root/file_1 test/data/sample_root/file_2 test/data/sample_root/node_1/file_1_1]} {[test/data/sample_root/node_1/node_1_1/file_1_1_1]}]"

		if chunksString != expectedChunksString {
			fmt.Println("Received:", chunksString)
			fmt.Println("Expected:", expectedChunksString)
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
