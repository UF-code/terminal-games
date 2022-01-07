package main

import (
	"fmt"
	"gopackages/utility"
)

func main() {
	fmt.Println("Hey")
	fmt.Println(utility.GetMyName())
	fmt.Println(utility.GetPackageName())
	fmt.Println(utility.GetName())

	s := utility.Test{
		Speak: "HEY",
	}

	fmt.Println(s)

	n := main.Name()

}
