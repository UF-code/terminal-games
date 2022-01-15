package CreateItem

import "fmt"

type Item struct {
	Name   string
	Power  int
	Health int
}

func CreateItem() {
	fmt.Println("CreateItem")
}
