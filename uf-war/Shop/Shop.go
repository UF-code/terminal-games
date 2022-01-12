package Shop

import (
	"fmt"
	"uf-war/Character"
)

type Market struct {
	Character.Character_Template
}

func Market_test() {
	m := Market{}

	fmt.Println(m)
}
