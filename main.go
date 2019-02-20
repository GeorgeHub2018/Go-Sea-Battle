package main

import (
	"fmt"
	"time"
)

func main() {
	chTurn := make(chan Turn)

	player1 := PlayerAlgorithm{}
	player1.SelfMap = player1.MapStrategy()
	go func(chTurn chan Turn) {
		firstTurn1 := player1.AttackStrategy(Turn{Point{-1, -1}, Start})
		chTurn <- firstTurn1
		for range time.Tick(time.Second * 2) {
			turn1 := <-chTurn
			turn1 = player1.AttackStrategy(turn1)
			chTurn <- turn1
		}
	}(chTurn)

	player2 := PlayerRandom{}
	player2.SelfMap = player2.MapStrategy()
	go func(chTurn chan Turn) {
		for range time.Tick(time.Second * 2) {
			turn2 := <-chTurn
			turn2 = player2.AttackStrategy(turn2)
			chTurn <- turn2
		}
	}(chTurn)

	for range time.Tick(time.Second) {
		fmt.Println("")
		PrintMaps(player1.SelfMap, player1.EnemyMap, player2.SelfMap, player2.EnemyMap)
	}
}
