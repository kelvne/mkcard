package mkcard_test

import (
	"os"
	"testing"

	"github.com/kelvne/mkcard/src/mkcard"
)

var (
	testCmd           = "mkcard"
	testName          = "Test Card"
	testCost          = "5"
	testPower         = "3"
	testBaseText      = "This is a test card"
	testSnakeCaseName = "test_card"
)

func TestRun(t *testing.T) {
	t.Run("TestRun", func(t *testing.T) {
		t.Parallel()

		os.Args = []string{testCmd, testName, testCost, testPower}
		code := mkcard.Run()

		if code != 1 {
			t.Errorf("Expected exit code 1, got %d", code)
		}

		os.Args = []string{testCmd, testName, testCost, testPower, testBaseText}
		code = mkcard.Run()

		if code != 0 {
			t.Errorf("Expected exit code 0, got %d", code)
		}

		cwd, err := os.Getwd()
		if err != nil {
			t.Errorf("Error getting current working directory: %s", err.Error())
		}

		_, err = os.Stat(cwd + "/" + testSnakeCaseName + ".tscn")
		if err != nil {
			t.Errorf("Error checking for tscn file: %s", err.Error())
		}
		os.Remove(cwd + "/" + testSnakeCaseName + ".tscn")

		_, err = os.Stat(cwd + "/" + testSnakeCaseName + ".gd")
		if err != nil {
			t.Errorf("Error checking for tscn file: %s", err.Error())
		}
		os.Remove(cwd + "/" + testSnakeCaseName + ".gd")
	})
}
