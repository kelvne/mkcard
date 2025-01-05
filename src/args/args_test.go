package args_test

import (
	"os"
	"strings"
	"testing"

	"github.com/kelvne/mkcard/src/args"
)

func TestGetArgs(t *testing.T) {
	t.Run("TestGetArgs", func(t *testing.T) {
		t.Parallel()

		os.Args = []string{"cmd", "arg1", "arg2", "arg3"}

		arguments := args.GetArgs()

		if len(arguments) != 3 {
			t.Errorf("Expected 3 arguments, got %d", len(arguments))
		}

		joined := strings.Join(os.Args[1:], ",")
		if joined != "arg1,arg2,arg3" {
			t.Errorf("Expected 'arg1,arg2,arg3', got %s", joined)
		}
	})
}
