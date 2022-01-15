package Creator

import (
	"fmt"
	"os"
)

func Create() {
	// fmt.Println("Hey")
	// CreateItem.CreateItem()
	// CreateQuest.CreateQuest()

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

func UpdateJson(path string) {

}
