package helpers_test

import (
	"fmt"

	"github.com/kelvne/mkcard/src/helpers"
)

func ExampleSnakeCase() {
	fmt.Println(helpers.SnakeCase("Test Card"))
	fmt.Println(helpers.SnakeCase("MyTest"))
	fmt.Println(helpers.SnakeCase("AnOtHeRCaRd"))
	fmt.Println(helpers.SnakeCase("DiFF CaRD"))
	// Output: test_card
	// my_test
	// an_ot_he_r_ca_rd
	// diff_card
}
