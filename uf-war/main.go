package main

import (
	"uf-war/Character"
	"uf-war/Character/PlayableCharacter"
	"uf-war/Inventory"
	"uf-war/Inventory/Tired"
)

func main() {
	Character.Character()
	PlayableCharacter.PC()
	Inventory.Test()
	Tired.Hey()
}
