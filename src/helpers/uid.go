package helpers

import (
	"crypto/rand"

	"github.com/btcsuite/btcutil/base58"
)

// GenerateGodotUID generates an base-58 unique identifier similar the ones
// used by the Godot Engine.
func GenerateGodotUID() string {
	byteArray := make([]byte, 16)
	rand.Read(byteArray)

	return base58.Encode(byteArray)
}
