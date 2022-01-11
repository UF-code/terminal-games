package main

import (
	"rpg/Character"
	"rpg/Character/PlayableCharacter"
	"rpg/Inventory"
	"rpg/Inventory/Tired"
)

func main() {
	Character.Character()
	PlayableCharacter.PC()
	Inventory.Test()
	Tired.Hey()
}
