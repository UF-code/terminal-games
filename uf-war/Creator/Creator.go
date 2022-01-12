package Creator

import (
	"fmt"
	"os"
	"uf-war/Creator/CreateItem"
	"uf-war/Creator/CreateQuest"
)

func Create() {
	fmt.Println("Hey")
	CreateItem.CreateItem()
	CreateQuest.CreateQuest()
}

func CreateJson(FileName string) {
	f, err := os.Create(FileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%v Successfully Created!\n", FileName)
	defer f.Close()
}

func UpdateJson(FileName string) {

}