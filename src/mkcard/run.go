package mkcard

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/kelvne/mkcard/src/args"
	"github.com/kelvne/mkcard/src/card"
	"github.com/kelvne/mkcard/templates"
)

var (
	arguments []string
	crd       *card.Card
	cwd       string
	err       error
)

// Run is the entry point for the mkcard package it returns the exit code of the program
func Run() int {
	arguments = args.GetArgs()
	if code := validateArgs(); code != 0 {
		return code
	}

	cardArgs := card.NewCardArgsFromOsArgs()
	crd = card.NewCardFromArgs(cardArgs)

	cwd, err = os.Getwd()
	if err != nil {
		fmt.Println("Unable to reach current working directory, probably thanks to symbolic links")
		fmt.Printf("Error: %s\n", err.Error())
		return 1
	}

	if code := makeTscn(); code != 0 {
		return code
	}

	if code := makeGdScript(); code != 0 {
		return code
	}

	return 0
}

func validateArgs() int {
	if len(arguments) < 4 {
		fmt.Println("Usage: mkcard <card_name> <card_cost> <card_power> <card_base_text>")
		return 1
	}

	return 0
}

func makeTscn() int {
	sceneTemplate := template.New("scene")
	sceneTemplate, err := sceneTemplate.Parse(templates.TscnTemplate)
	if err != nil {
		fmt.Println("Error parsing template: ", err.Error())

		return 1
	}

	fpath := filepath.Join(cwd, fmt.Sprintf("%s.tscn", crd.SnakeCaseName))

	file, err := os.Create(fpath)
	if err != nil {
		fmt.Println("Error creating file: ", err.Error())

		return 1
	}

	sceneTemplate.Execute(file, crd)

	return 0
}

func makeGdScript() int {
	scriptTemplate := template.New("script")
	scriptTemplate, err := scriptTemplate.Parse(templates.ScriptTemplate)
	if err != nil {
		fmt.Println("Error parsing template: ", err.Error())

		return 1
	}

	fpath := filepath.Join(cwd, fmt.Sprintf("%s.gd", crd.SnakeCaseName))

	file, err := os.Create(fpath)
	if err != nil {
		fmt.Println("Error creating file: ", err.Error())

		return 1
	}

	scriptTemplate.Execute(file, crd)

	return 0
}
