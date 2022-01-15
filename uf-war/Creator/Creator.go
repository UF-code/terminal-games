package Creator

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"uf-war/Creator/CreateItem"
	"uf-war/Creator/CreateQuest"
)

type Item struct {
	CreateItem.Item
}

type Quest struct {
	CreateQuest.Quest
}

func (i Item) conv() {
	bs, err := json.Marshal(i.Name)
	fmt.Println(bs)
	fmt.Println(err)
	fmt.Println("Item")
	// return bs
	// return bs
}
func (q Quest) conv() []uint8 {
	bs, err := json.Marshal(q.Name)
	fmt.Println(bs)
	fmt.Printf("%T \n", bs)
	fmt.Println(err)
	fmt.Println("Quest")
	return bs
	// return bs
	// return bs
}
func Create() {
	// fmt.Println("Hey")
	// CreateItem.CreateItem()
	// CreateQuest.CreateQuest()
	// i1 := Item{
	// 	CreateItem.Item{
	// 		Name: "Hey Dude",
	// 	},
	// }
	// i1.conv()
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
	// err := ioutil.WriteFile(path, data, 0644)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// i1 := Item{
	// 	CreateItem.Item{
	// 		Name: "Hey Dude",
	// 	},
	// }
	// i1.conv()

	switch d.(type) {
	case Item:
		d.conv()
	case Quest:
		d.conv()
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
