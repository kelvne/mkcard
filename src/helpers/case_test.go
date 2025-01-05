package helpers_test

import (
	"testing"

	"github.com/kelvne/mkcard/src/helpers"
)

func TestSnakeCase(t *testing.T) {
	t.Run("TestSnakeCase", func(t *testing.T) {
		t.Parallel()

		cases := []string{
			"Test Card",
			"MyTest",
			"testCardin",
			"Card TesT",
			"AnOtHeRCaRd",
			"DiFF CaRD",
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
			if snaked := helpers.SnakeCase(c); snaked != expected[i] {
				t.Errorf("Expected %s, got %s", expected[i], snaked)
			}
		}
	})
}
