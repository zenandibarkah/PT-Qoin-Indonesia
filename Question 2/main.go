package main

import (
	function "dicegame/functions"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var totalPlayer, totalDice int

	fmt.Print("Masukkan jumlah pemain: ")
	fmt.Scan(&totalPlayer)
	fmt.Print("Masukkan jumlah dadu: ")
	fmt.Scan(&totalDice)

	dicePlayer := make([][]int, totalPlayer)
	pointPlayer := make([]int, totalPlayer)
	for i := 0; i < totalPlayer; i++ {
		dicePlayer[i] = function.ScrollDice(totalDice)
	}

	turn := 1

	for {
		fmt.Printf("\n=======================\nGiliran %d lempar dadu:\n", turn)
		for i := 0; i < totalPlayer; i++ {
			function.PrintResult(i, pointPlayer, dicePlayer)
		}

		moveDice := make([][]int, totalPlayer)

		// Evaluation result
		for j := 0; j < totalPlayer; j++ {
			dicePlayer[j], moveDice[j] = function.EvaluationAndDistribution(j, dicePlayer[j], pointPlayer)
		}

		// Process of distributing dice with a value of 1 to other players
		for k := 0; k < totalPlayer; k++ {
			sidePlayer := (k + 1) % totalPlayer
			dicePlayer[sidePlayer] = append(dicePlayer[sidePlayer], moveDice[k]...)
		}

		fmt.Println("Setelah evaluasi:")

		for l := 0; l < totalPlayer; l++ {
			function.PrintResult(l, pointPlayer, dicePlayer)
		}

		// Process check if there is only 1 player left with the dice
		activePlayer := []int{}
		for i, dice := range dicePlayer {
			if len(dice) > 0 {
				activePlayer = append(activePlayer, i)
			}
		}

		if len(activePlayer) == 1 {
			fmt.Println("\n=======================")
			fmt.Printf("Game berakhir karena hanya pemain %d yang memiliki dadu.\n", activePlayer[0]+1)
			break
		}

		// Process re-roll the remaining dice for each player
		for n := 0; n < totalPlayer; n++ {
			if len(dicePlayer[n]) > 0 {
				dicePlayer[n] = function.ScrollDice(len(dicePlayer[n]))
			}
		}

		turn++
	}

	winner := 0
	for i := 1; i < totalPlayer; i++ {
		if pointPlayer[i] > pointPlayer[winner] {
			winner = i
		}
	}

	fmt.Printf("Game dimenangkan oleh pemain #%d dengan %d point.\n", winner+1, pointPlayer[winner])

}
