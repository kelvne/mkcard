package card

import "github.com/kelvne/mkcard/src/helpers"

// NewCardArgs creates a new structured card arguments object
func NewCardArgs(name string, cost, power int, baseText string) *CardArgs {
	return &CardArgs{
		Name:     name,
		Cost:     cost,
		Power:    power,
		BaseText: baseText,
	}
}

// NewCardFromArgs creates a new card object from a CardArgs object
func NewCardFromArgs(cardArgs *CardArgs) *Card {
	return &Card{
		UID:           helpers.GenerateGodotUID(),
		Name:          cardArgs.Name,
		SnakeCaseName: cardArgs.SnakeCaseName(),
		Cost:          cardArgs.Cost,
		Power:         cardArgs.Power,
		BaseText:      cardArgs.BaseText,
	}
}
