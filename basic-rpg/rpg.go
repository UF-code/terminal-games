package main

import "fmt"

type Character struct {
	Class
	Name       string
	Level      int
	Experience int
	Inventory
}

type Class struct {
	Job    string
	Health int
	Power  int
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
	Items []Item
}

type Inventory struct {
	Wallet
	Bag
}

func (player *Player) Purchase(item Item) {
	// Purchased := false
	if player.Coin >= item.Price {
		player.Coin -= item.Price
		player.Items = append(player.Items, item)
	} else {
		fmt.Printf("You don't have enough money to buy %v Item's Price: %v\n", item.Name, item.Price)
		fmt.Println("Your Balance: ", player.Coin)
	}
	// return Purchased
}

// func (player *Player) AddToBag(item Item) {
// 	player.Items = append(player.Items, item)
// }

// func (player *Player) addName() {
// 	player.Name = "uf-war"
// }

func CreateClass() (Class, Class) {
	WarriorKing := Class{
		Job:    "Warior King",
		Health: 2000,
		Power:  1000,
	}

	WitchKing := Class{
		Job:    "Witch King",
		Health: 1000,
		Power:  2000,
	}

	return WarriorKing, WitchKing
}

func CreatePlayer() Player {
	player := new(Player)

	return *player
}

func main() {
	player := CreatePlayer()
	fmt.Println(player)
	// player.addName()
	fmt.Println(player)

	// var x Player // x == nil
	// var x *Player
	// x = new(Player) // x == &Player{}

	// p := Player{
	// 	Character: Character{
	// 		Name:       "uf-war",
	// 		Class:      "Warrior King",
	// 		Level:      100,
	// 		Experience: 0,
	// 		Health:     10000,
	// 		Power:      1000,
	// 		Inventory: Inventory{
	// 			Wallet: Wallet{
	// 				Coin: 1000000,
	// 			},
	// 		},
	// 	},
	// }

	// fmt.Println(p)
	// fmt.Println(p.Name)

	// w := Item{
	// 	Name:  "God King's Platinium Sword",
	// 	Price: 199,
	// 	Weapon: Weapon{
	// 		Power: 1949,
	// 	},
	// }
	// w2 := Item{
	// 	Name:  "God King's Gold Sword",
	// 	Price: 99,
	// 	Weapon: Weapon{
	// 		Power: 949,
	// 	},
	// }

	// fmt.Println(w.Name)
	// fmt.Println(w.Price)
	// fmt.Println(w.Power)

	// p.Purchase(w)
	// // fmt.Println(*p.Purchase(w))

	// fmt.Println(p.Items)

	// p.Items = append(p.Items, w)

	// fmt.Println(p.Items)

	// p.Items = append(p.Items, w2)
	// fmt.Println(p.Items)

	// // p.AddToBag(w)

	// fmt.Println(p.Items)
}
