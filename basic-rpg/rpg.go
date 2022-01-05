package main

import "fmt"

type Character struct {
	Name       string
	Class      string
	Level      int
	Experience int
	Health     int
	Power      int
	Inventory
}

type NPC_Shop struct {
	Character
}

type NPC_Quest struct {
	Character
}

type Enemy struct {
	Character
}

type Player struct {
	Character
}

type Item struct {
	Name  string
	Price int
}

type Potion struct {
	Item
	Health int
}

type Weapon struct {
	Item
	Power int
}

type Wallet struct {
	Coin int
}

type Bag struct {
	Weapon
	Potion
}

type Inventory struct {
	Wallet
	Bag
}

func (player Player) Purchase(item Item) (int, bool) {
	Purchased := false
	if player.Coin >= item.Price {
		player.Coin -= item.Price
		Purchased = true
	} else {
		fmt.Printf("You don't have enough money to buy %v that item Price: %v\n", item.Name, item.Price)
		fmt.Println("Your Balance: ", player.Coin)
	}
	return player.Coin, Purchased
}

func (player Player) AddToInventory() {

}

func main() {

}
