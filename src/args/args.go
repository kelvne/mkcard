package args

import "os"

// GetArgs returns the command line arguments without the call to the CLI
func GetArgs() []string {
	return os.Args[1:]
}
