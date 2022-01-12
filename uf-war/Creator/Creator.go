package Creator

import (
	"fmt"
	"io"
	"os"
	"strings"
	"uf-war/Creator/CreateItem"
	"uf-war/Creator/CreateQuest"
)

func Create() {
	fmt.Println("Hey")
	CreateItem.CreateItem()
	CreateQuest.CreateQuest()
}

func CreateJson() {
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	r := strings.NewReader("Hey Whatsup!")
	io.Copy(f, r)
}
