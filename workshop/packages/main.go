package main

import (
	"fmt"
	"packages/utility"
)

func main() {
	fmt.Println("hey main")

	u := utility.GetName()
	fmt.Println(u)

	utility.HeyUtil()

	utility.HeyStrUtil()
}
