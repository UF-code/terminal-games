package Creator

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"uf-war/Creator/CreateItem"
)

type Item struct {
	CreateItem.Item
}

func (i Item) conv() []byte {
	bs, err := json.Marshal(i.Name)
	fmt.Println(err)
	return bs
}
func Create() {
	// fmt.Println("Hey")
	// CreateItem.CreateItem()
	// CreateQuest.CreateQuest()
	i1 := Item{
		CreateItem.Item{
			Name: "Hey Dude",
		},
	}
	i1.conv()
}

func CheckFileExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func CreateJson(path string) {
	f, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%v Successfully Created!\n", path)
	defer f.Close()
}

type data interface {
	conv()
}

func WriteToJson(path string, d data) {
	// message := []byte("Hello, Gophers!")

	i1 := Item{
		CreateItem.Item{
			Name: "Hey Dude",
		},
	}
	i1.conv()

	switch d.(type) {
	case Item:
		err := ioutil.WriteFile(path, i1.conv(), 0644)
		if err != nil {
			log.Fatal(err)
		}
	}

}

func ReadJson(path string) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(content))
}

func UpdateJson() {

}

func DeleteJson() {

}
