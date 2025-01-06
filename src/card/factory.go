package card

import (
	"strconv"

	"github.com/kelvne/mkcard/src/args"
	"github.com/kelvne/mkcard/src/helpers"
)

// NewCardArgsFromOsArgs creates a new CardArgs object from the command line arguments
func NewCardArgsFromOsArgs() *CardArgs {
	arguments := args.GetArgs()

	var (
		name     string
		cost     int
		power    int
		baseText string
	)

	for i, arg := range arguments {
		switch i {
		case 0:
			name = arg
		case 1:
			cost, _ = strconv.Atoi(arg)
		case 2:
			power, _ = strconv.Atoi(arg)
		case 3:
			baseText = arg
		}
	}

	return NewCardArgs(name, cost, power, baseText)
}

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
