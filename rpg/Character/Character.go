package Character

import (
	"fmt"
	NonPlayableCharacter "rpg/Character/NPC"
	PlayableCharacter "rpg/Character/PC"
)

func Character() {
	fmt.Println("Hey")
	NonPlayableCharacter.NPC()
	PlayableCharacter.PC()

}
