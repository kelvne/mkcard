package card

// CardArgs represents the structured parsed arguments for making
// the card files.
type CardArgs struct {
	Name     string
	Cost     int
	Power    int
	BaseText string
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
