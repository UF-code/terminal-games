package CreateItem

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	// "uf-war/Creator"
)

func CreateItem() {
	// Creator.CreateTest()
	// CreateAssassinItems()
	// CreateWarriorItems()
	// CreateMageItems()

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Which Class Item You Want To Create")
	fmt.Println("Assassin(1) Warrior(2) Mage(3)")
	scanner.Scan()
	choice, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	// fmt.Printf("%v %T \n", choice, choice)

	switch {
	case choice == 1:
		CreateAssassinItem()
	case choice == 2:
		CreateWarriorItem()
	case choice == 3:
		CreateMageItem()
	default:
		fmt.Println("There Are No Class As You Mention...")
		fmt.Println("Exiting...")
	}

}

func CreateAssassinItem() {
	fmt.Println("Create Assassins Items")
}

func CreateWarriorItem() {
	fmt.Println("Create Warrior Items")
}

func CreateMageItem() {
	fmt.Println("Create Mage Items")
}
