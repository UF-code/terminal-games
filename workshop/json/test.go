package main

import (
	"fmt"
	"os"
)

func main() {

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

func WriteJson(FileName string) {

}

func ReadJson(FileName string) {

}

func UpdateJson(FileName string) {

}

func DeleteJson(FileName string) {

}
