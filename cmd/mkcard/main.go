package main

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
	cardArgs  *card.CardArgs
	crd       *card.Card
	cwd       string
)

func init() {
	arguments = args.GetArgs()
	validateArgs()
	cardArgs = card.NewCardArgsFromOsArgs()
	crd = card.NewCardFromArgs(cardArgs)
}

func main() {
	var err error
	cwd, err = os.Getwd()
	if err != nil {
		fmt.Println("Unable to reach current working directory, probably thanks to symbolic links")
		panic(1)
	}

	makeTscn()
	makeGdScript()
}

func validateArgs() {
	if len(arguments) < 4 {
		fmt.Println("Usage: mkcard <card_name> <card_cost> <card_power> <card_base_text>")
		panic(1)
	}
}

func makeTscn() {
	sceneTemplate := template.New("scene")
	sceneTemplate, err := sceneTemplate.Parse(templates.TscnTemplate)
	if err != nil {
		fmt.Println("Error parsing template: ", err.Error())
		panic(1)
	}

	fpath := filepath.Join(cwd, fmt.Sprintf("%s.tscn", crd.SnakeCaseName))

	file, err := os.Create(fpath)
	if err != nil {
		fmt.Println("Error creating file: ", err.Error())
		panic(1)
	}

	sceneTemplate.Execute(file, crd)
}

func makeGdScript() {
	scriptTemplate := template.New("script")
	scriptTemplate, err := scriptTemplate.Parse(templates.ScriptTemplate)
	if err != nil {
		fmt.Println("Error parsing template: ", err.Error())
		panic(1)
	}

	fpath := filepath.Join(cwd, fmt.Sprintf("%s.gd", crd.SnakeCaseName))

	file, err := os.Create(fpath)
	if err != nil {
		fmt.Println("Error creating file: ", err.Error())
		panic(1)
	}

	scriptTemplate.Execute(file, crd)
}
