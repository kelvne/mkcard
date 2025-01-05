package card_test

import (
	"os"
	"strconv"
	"testing"

	"github.com/kelvne/mkcard/src/args"
	"github.com/kelvne/mkcard/src/card"
)

var (
	testCmd      = "mkcard"
	testName     = "Test Card"
	testCost     = "5"
	testPower    = "3"
	testBaseText = "This is a test card"
)

func TestNewCardArgs(t *testing.T) {
	t.Run("TestNewCardArgs", func(t *testing.T) {
		setOsArgs()
		args := args.GetArgs()

		if len(args) != 4 {
			t.Errorf("Expected 4 arguments, got %d", len(args))
		}

		cost := getCostArg(t, args)
		power := getPowerArg(t, args)

		cardArgs := card.NewCardArgs(args[0], cost, power, args[3])

		if cardArgs == nil {
			t.Error("Expected cardArgs to not be nil")
			return
		}

		if cardArgs.Name != testName {
			t.Errorf("Expected name to be %s, got %s", testName, cardArgs.Name)
		}

		if cardArgs.Cost != cost {
			t.Errorf("Expected cost to be %d, got %d", cost, cardArgs.Cost)
		}

		if cardArgs.Power != power {
			t.Errorf("Expected power to be %d, got %d", power, cardArgs.Power)
		}

		if cardArgs.BaseText != testBaseText {
			t.Errorf("Expected base text to be %s, got %s", testBaseText, cardArgs.BaseText)
		}
	})
}

func TestNewCardFromArgs(t *testing.T) {
	t.Run("TestNewCardArgs", func(t *testing.T) {
		setOsArgs()
		args := args.GetArgs()

		if len(args) != 4 {
			t.Errorf("Expected 4 arguments, got %d", len(args))
		}

		cost := getCostArg(t, args)
		power := getPowerArg(t, args)

		cardArgs := card.NewCardArgs(args[0], cost, power, args[3])
		card := card.NewCardFromArgs(cardArgs)

		if card == nil {
			t.Error("Expected cardArgs to not be nil")
			return
		}

		if len(card.UID) != 22 {
			t.Errorf("Expected UID to be 22 characters, got %d", len(card.UID))
		}

		if card.SnakeCaseName != cardArgs.SnakeCaseName() {
			t.Error("Expected snake case name to be right")
		}

		if card.Name != testName {
			t.Errorf("Expected name to be %s, got %s", testName, card.Name)
		}

		if card.Cost != cost {
			t.Errorf("Expected cost to be %d, got %d", cost, card.Cost)
		}

		if card.Power != power {
			t.Errorf("Expected power to be %d, got %d", power, card.Power)
		}

		if card.BaseText != testBaseText {
			t.Errorf("Expected base text to be %s, got %s", testBaseText, card.BaseText)
		}
	})
}

func setOsArgs() {
	os.Args = []string{testCmd, testName, testCost, testPower, testBaseText}
}

func getCostArg(t *testing.T, args []string) int {
	cost, err := strconv.Atoi(args[1])
	if err != nil {
		t.Errorf("Expected cost to be an integer, got %s, err: %s", args[1], err.Error())
	}

	return cost
}

func getPowerArg(t *testing.T, args []string) int {
	power, err := strconv.Atoi(args[2])
	if err != nil {
		t.Errorf("Expected power to be an integer, got %s, err: %s", args[2], err.Error())
	}

	return power
}
