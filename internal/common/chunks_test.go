package common

import (
	"testing"

	"github.com/google/uuid"
)

func TestCunk(t *testing.T) {
	t.Run("Chunk construction", func(t *testing.T) {
		files := []string{"foo", "bar"}
		c := MakeChunk(&files)

		defaultUuid := uuid.UUID{}
		if c.Uuid == defaultUuid {
			t.Logf("UUID not initialized")
			t.Fail()
		}

		if len(c.Files) != 2 {
			t.Fail()
		}
	})
}
