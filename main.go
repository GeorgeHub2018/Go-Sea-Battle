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
