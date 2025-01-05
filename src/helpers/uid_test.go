package helpers_test

import (
	"testing"

	"github.com/kelvne/mkcard/src/helpers"
)

func TestGenerateGodotUID(t *testing.T) {
	t.Run("TestGenerateGodotUID", func(t *testing.T) {
		t.Parallel()

		uid := helpers.GenerateGodotUID()

		if l := len(uid); l != 22 {
			t.Errorf("Expected string with 22 characters, got %d characters", l)
		}
	})
}
