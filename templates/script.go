package templates

const ScriptTemplate = `extends Node

const card_name: String = "{{.Name}}"
const base_cost: int = {{.Cost}}
const base_power: int = {{.Power}}
const base_text: String = "{{.BaseText}}"

var alterations: Dictionary = {
	"cost": 0,
	"power": 0,
	"text_disabled": false,
}

func _ready() -> void:
	pass
`
