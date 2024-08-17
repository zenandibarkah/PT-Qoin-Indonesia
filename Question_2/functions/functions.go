package functions

import (
	"fmt"
	"math/rand"
)

func ScrollDice(totalDice int) []int {
	dice := make([]int, totalDice)
	for i := 0; i < totalDice; i++ {
		dice[i] = rand.Intn(6) + 1
	}

	return dice
}

func EvaluationAndDistribution(player int, dicePlayer []int, point []int) (result []int, moveValue []int) {
	for _, value := range dicePlayer {
		if value == 6 {
			point[player]++
		} else if value == 1 {
			moveValue = append(moveValue, 1)
		} else {
			result = append(result, value)
		}
	}

	return
}

func PrintResult(player int, point []int, dicePlayer [][]int) {
	fmt.Printf("	Pemain #%d (%d): ", player+1, point[player])
	if len(dicePlayer[player]) == 0 {
		fmt.Println("_ (Berhenti bermain karena tidak memiliki dadu)")
	} else {
		for i, dice := range dicePlayer[player] {
			if i > 0 {
				fmt.Print(", ")
			}
			fmt.Print(dice)
		}
		fmt.Println()
	}
}
