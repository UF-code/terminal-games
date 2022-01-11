package PlayableCharacter

import (
	"fmt"
	"uf-war/Character/NonPlayableCharacter"
)

func PC() {
	fmt.Println("Hey PC")
	NonPlayableCharacter.NPC()
	fmt.Println("Hey PC")

	// s := Character.Character_Template{
	// 	Health: 100,
	// 	Power:  200,
	// }

	// fmt.Println(s)
}
