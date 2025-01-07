package helpers

import (
	"crypto/rand"
	"strings"

	"github.com/btcsuite/btcutil/base58"
)

// GenerateGodotUID generates an base-58 unique identifier similar the ones
// used by the Godot Engine.
func GenerateGodotUID() string {
	byteArray := make([]byte, 16)
	rand.Read(byteArray)

	return strings.ToLower(base58.Encode(byteArray))
}
