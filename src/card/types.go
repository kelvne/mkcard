package card

import (
	"github.com/kelvne/mkcard/src/helpers"
)

// CardArgs represents the structured parsed arguments for making
// the card files.
type CardArgs struct {
	Name     string
	Cost     int
	Power    int
	BaseText string
}

// SnakeCaseName returns the name of the card in snake case
func (c *CardArgs) SnakeCaseName() string {
	return helpers.SnakeCase(c.Name)
}

// Card represents all information about a card
type Card struct {
	UID           string
	Name          string
	SnakeCaseName string
	Cost          int
	Power         int
	BaseText      string
}
