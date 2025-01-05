package card_test

import (
	"testing"

	"github.com/kelvne/mkcard/src/card"
)

func TestSnakeCaseName(t *testing.T) {
	t.Run("TestSnakeCaseName", func(t *testing.T) {
		t.Parallel()

		cases := []*card.CardArgs{
			{Name: "Test Card"},
			{Name: "MyTest"},
			{Name: "testCardin"},
			{Name: "Card TesT"},
			{Name: "AnOtHeRCaRd"},
			{Name: "DiFF CaRD"},
		}

		expected := []string{
			"test_card",
			"my_test",
			"test_cardin",
			"card_test",
			"an_ot_he_r_ca_rd",
			"diff_card",
		}

		for i, c := range cases {
			if snaked := c.SnakeCaseName(); snaked != expected[i] {
				t.Errorf("Expected %s, got %s", expected[i], snaked)
			}
		}
	})
}
