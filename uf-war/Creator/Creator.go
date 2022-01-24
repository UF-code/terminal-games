package Creator

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"


func CreateTest() {
	fmt.Println("Create Test")
}


func Create() {

	fmt.Println("Getting use to it!")
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
