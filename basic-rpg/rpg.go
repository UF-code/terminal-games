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
	Potion
	Weapon
}

type Potion struct {
	Health int
}

type Weapon struct {
	Power int
}

type Wallet struct {
	Coin int
}

type Bag struct {
	Item
}

type Inventory struct {
	Wallet
	Bag
}

func (player Player) Purchase(item Item) (int, bool, Item) {
	Purchased := false
	if player.Coin >= item.Price {
		player.Coin -= item.Price
		Purchased = true
	} else {
		fmt.Printf("You don't have enough money to buy %v Item's Price: %v\n", item.Name, item.Price)
		fmt.Println("Your Balance: ", player.Coin)
	}
	return player.Coin, Purchased, item
}

func (player Player) AddToInventory(item Item) {

}

func main() {
	p := Player{
		Character: Character{
			Name:       "uf-war",
			Class:      "Warrior King",
			Level:      100,
			Experience: 0,
			Health:     10000,
			Power:      1000,
			Inventory: Inventory{
				Wallet: Wallet{
					Coin: 1000000,
				},
				Bag: Bag{},
			},
		},
	}

	fmt.Println(p)
	fmt.Println(p.Name)

	w := Item{
		Name:  "God King's Platinium Sword",
		Price: 199,
		Weapon: Weapon{
			Power: 1949,
		},
	}

	fmt.Println(w.Name)
	fmt.Println(w.Price)
	fmt.Println(w.Power)

	fmt.Println(p.Purchase(w))

	b := Bag{
		w,
	}

	fmt.Println(b)

	fmt.Println(p.Bag)
	p.Bag = b
	fmt.Println(p.Bag)

}
