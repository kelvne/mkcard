package card

import "github.com/kelvne/mkcard/src/helpers"

// SnakeCaseName returns the name of the card in snake case
func (c *CardArgs) SnakeCaseName() string {
	return helpers.SnakeCase(c.Name)
}
