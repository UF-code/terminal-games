package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Player struct {
	Player   string
	Position int
}

func (p Player) DecidePosition(player string, position int) (string, int) {
	p.Player = player
	p.Position = position

	return p.Player, p.Position
}

func Display(display_list [9]string) {

	fmt.Println("Tic Tac Toe v.0.0.1")
	fmt.Println()
	fmt.Println()

	fmt.Printf(" %v | %v | %v \n", display_list[0], display_list[1], display_list[2])
	fmt.Printf("-----------\n")
	fmt.Printf(" %v | %v | %v \n", display_list[3], display_list[4], display_list[5])
	fmt.Printf("-----------\n")
	fmt.Printf(" %v | %v | %v \n", display_list[6], display_list[7], display_list[8])

	fmt.Println()
	fmt.Println()
	fmt.Println("Tic Tac Toe v.0.0.1")

}

func Input() int {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Position You Want to Fill: ")
	scanner.Scan()
	position, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		fmt.Println(err)
		panic("You have caused a trouble kid!")
	} else {
		return int(position) - 1
	}

}

func CheckWinner(display_list [9]string) (string, bool) {
	Winner := ""
	GameOver := false
	if display_list[0] == display_list[1] && display_list[1] == display_list[2] && display_list[2] == "X" ||
		display_list[3] == display_list[4] && display_list[4] == display_list[5] && display_list[5] == "X" ||
		display_list[6] == display_list[7] && display_list[7] == display_list[8] && display_list[8] == "X" ||
		display_list[0] == display_list[3] && display_list[3] == display_list[6] && display_list[6] == "X" ||
		display_list[1] == display_list[4] && display_list[4] == display_list[7] && display_list[7] == "X" ||
		display_list[2] == display_list[5] && display_list[5] == display_list[8] && display_list[8] == "X" ||
		display_list[0] == display_list[4] && display_list[4] == display_list[8] && display_list[8] == "X" ||
		display_list[2] == display_list[4] && display_list[4] == display_list[6] && display_list[6] == "X" {
		Winner = "X"
		GameOver = true
	} else if display_list[0] == display_list[1] && display_list[1] == display_list[2] && display_list[2] == "O" ||
		display_list[3] == display_list[4] && display_list[4] == display_list[5] && display_list[5] == "O" ||
		display_list[6] == display_list[7] && display_list[7] == display_list[8] && display_list[8] == "O" ||
		display_list[0] == display_list[3] && display_list[3] == display_list[6] && display_list[6] == "O" ||
		display_list[1] == display_list[4] && display_list[4] == display_list[7] && display_list[7] == "O" ||
		display_list[2] == display_list[5] && display_list[5] == display_list[8] && display_list[8] == "O" ||
		display_list[0] == display_list[4] && display_list[4] == display_list[8] && display_list[8] == "O" ||
		display_list[2] == display_list[4] && display_list[4] == display_list[6] && display_list[6] == "O" {
		Winner = "O"
		GameOver = true
	}
	return Winner, GameOver
}

func (p Player) Game() {
	turns := []string{"X", "O", "X", "O", "X", "O", "X", "O", "X"}
	positions := [9]string{" ", " ", " ", " ", " ", " ", " ", " ", " "}

	for i, turn := range turns {
		player, position := p.DecidePosition(turn, Input())

		if positions[position] != " " {
			panic("You have caused a trouble kid!")
		}
		positions[position] = player

		Display(positions)
		Winner, GameOver := CheckWinner(positions)

		if GameOver {
			fmt.Println("Winner is", Winner)
			break
		} else if !GameOver && i == 8 {
			fmt.Println("Tie ")
			break
		}

	}

}

func main() {
	p := Player{}

	p.Game()

}
