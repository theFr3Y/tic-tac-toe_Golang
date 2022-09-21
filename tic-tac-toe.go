package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

// clear terminal
func Clear() {
	system := runtime.GOOS
	if system == "windows" {
		console := exec.Command("cmd", "/c", "cls")
		console.Stdout = os.Stdout
		console.Run()
	} else {
		console := exec.Command("clear")
		console.Stdout = os.Stdout
		console.Run()
	}
}

func main() {
	Clear()
	lists := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	for i := 0; i < len(lists); i++ {

		// display table and turn of PLayer1(X)
		fmt.Printf("%v | %v | %v \n", lists[0], lists[1], lists[2])
		fmt.Printf("--------- \n")
		fmt.Printf("%v | %v | %v \n", lists[3], lists[4], lists[5])
		fmt.Printf("--------- \n")
		fmt.Printf("%v | %v | %v \n", lists[6], lists[7], lists[8])

		// Moving operations in the desired place and get puts
		var player1 int
		fmt.Printf("Player1(X) turn: ")
		fmt.Scanln(&player1)
		Clear()

		// If the entered place is already selected
		if lists[player1-1] == "O" || lists[player1-1] == "X" {
			fmt.Print("You cannot place it in the selected place")
			break
		} else {
			lists[player1-1] = "X"
		}

		// Player1(X) conditions to win or draw
		if lists[0] == "X" && lists[1] == "X" && lists[2] == "X" {
			fmt.Println("player1 is winner!")
			break
		} else if lists[3] == "X" && lists[4] == "X" && lists[5] == "X" {
			fmt.Println("player1 is winner!")
			break
		} else if lists[6] == "X" && lists[7] == "X" && lists[8] == "X" {
			fmt.Println("player1 is winner!")
			break
		} else if lists[0] == "X" && lists[3] == "X" && lists[6] == "X" {
			fmt.Println("player1 is winner!")
			break
		} else if lists[1] == "X" && lists[4] == "X" && lists[7] == "X" {
			fmt.Println("player1 is winner!")
			break
		} else if lists[2] == "X" && lists[5] == "X" && lists[8] == "X" {
			fmt.Println("player1 is winner!")
			break
		} else if lists[0] == "X" && lists[4] == "X" && lists[8] == "X" {
			fmt.Println("player1 is winner!")
			break
		} else if lists[2] == "X" && lists[4] == "X" && lists[6] == "X" {
			fmt.Println("player1 is winner!")
			break
		} else if i == 4 {
			fmt.Println("Draw...")
			break
		}

		// display table and turn of PLayer2(O)
		fmt.Printf("%v | %v | %v \n", lists[0], lists[1], lists[2])
		fmt.Printf("--------- \n")
		fmt.Printf("%v | %v | %v \n", lists[3], lists[4], lists[5])
		fmt.Printf("--------- \n")
		fmt.Printf("%v | %v | %v \n", lists[6], lists[7], lists[8])

		// Moving operations in the desired place and get puts
		var player2 int
		fmt.Printf("Player2(O) turn: ")
		fmt.Scanln(&player2)
		Clear()

		// If the entered location is already selected
		if lists[player2-1] == "X" || lists[player2-1] == "O" {
			fmt.Print("You cannot place it in the selected place")
			break
		} else {
			lists[player2-1] = "O"
		}

		// Player2(O) conditions to win or draw
		if lists[0] == "O" && lists[1] == "O" && lists[2] == "O" {
			fmt.Println("player2 is winner!")
			break
		} else if lists[3] == "O" && lists[4] == "O" && lists[5] == "O" {
			fmt.Println("player2 is winner!")
			break
		} else if lists[6] == "O" && lists[7] == "O" && lists[8] == "O" {
			fmt.Println("player2 is winner!")
			break
		} else if lists[0] == "O" && lists[3] == "O" && lists[6] == "O" {
			fmt.Println("player2 is winner!")
			break
		} else if lists[1] == "O" && lists[4] == "O" && lists[7] == "O" {
			fmt.Println("player2 is winner!")
			break
		} else if lists[2] == "O" && lists[5] == "O" && lists[8] == "O" {
			fmt.Println("player2 is winner!")
			break
		} else if lists[0] == "O" && lists[4] == "O" && lists[8] == "O" {
			fmt.Println("player2 is winner!")
			break
		} else if lists[2] == "O" && lists[4] == "O" && lists[6] == "O" {
			fmt.Println("player2 is winner!")
			break
		} else if i == 4 {
			fmt.Println("Draw...")
			break
		}
	}
}
