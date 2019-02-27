<<<<<<< HEAD
package main

import (
	"log"
	"os"
	"time"
)

const (
	debugs      = false
	warnings    = false
	errors      = false
	logs        = false
	isLogInFile = false
)

const (
	logFileName = "app.log"
)

// IsFileExists returns whether the given file or directory exists or not
func IsFileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

//DeleteFile delete file
func DeleteFile(path string) {
	// delete file
	os.Remove(path)
}

var chMessage = make(chan string)

func main() {
	if isLogInFile {
		if exist, err := IsFileExists(logFileName); (err == nil) && exist {
			DeleteFile(logFileName)
		}

		f, error := os.OpenFile(logFileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if error != nil {
			log.Fatal(error)
		}
		defer f.Close()
		log.SetOutput(f)
	}

	chTurn := make(chan Turn)
	chMain := make(chan int)

	// Bot 1
	var player1 Strategy = &BotRandom{}
	player1.SetID(1)
	StartStrategy(player1, chTurn, chMain)
	// Bot 2
	var player2 Strategy = &BotRandom{}
	player2.SetID(2)
	StartStrategy(player2, chTurn, chMain)

	// Print
	for range chMain {
		log.Println("")
		PrintMaps(player1.GetSelfMap(), player1.GetEnemyMap(), player2.GetSelfMap(), player2.GetEnemyMap())
		time.Sleep(100 * time.Millisecond)
	}

	os.Exit(0)
}
=======
package main

import (
	"fmt"
	"time"
)

const (
	debug = false
)

func main() {
	chTurn := make(chan Turn)

	player1 := PlayerAlgorithm{}
	player1.SelfMap = player1.MapStrategy()
	player1.SelfMap.Name = player1.Name + "Self"
	player1.EnemyMap.Name = player1.Name + "Enemy"
	go func(chTurn chan Turn) {
		firstTurn1 := player1.AttackStrategy(Turn{Point{-1, -1}, Start})
		chTurn <- firstTurn1
		for range time.Tick(time.Second) {
			turn1 := <-chTurn
			turn1 = player1.AttackStrategy(turn1)
			chTurn <- turn1
		}
	}(chTurn)

	player2 := PlayerRandom{}
	player2.SelfMap = player2.MapStrategy()
	player2.SelfMap.Name = player2.Name + "Self"
	player2.EnemyMap.Name = player2.Name + "Enemy"
	go func(chTurn chan Turn) {
		for range time.Tick(time.Second) {
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
>>>>>>> 09d48927fc4507cb93a42304ec539761626aa40b
