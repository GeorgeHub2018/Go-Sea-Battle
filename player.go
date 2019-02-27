package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"
)

//Answers
const (
	Start = iota
	Miss
	Goal
	Kill
	End
	Move
)

type (
	//Point struct
	Point struct {
		X int
		Y int
	}
	//Turn struct
	Turn struct {
		Point  Point
		Answer int
	}
)

//Strategy interface
type Strategy interface {
	MapStrategy() Map
	AttackStrategy(t Turn) Turn
	GetName() string
	GetSelfMap() Map
	SetSelfMap(m Map)
	GetEnemyMap() Map
	SetEnemyMap(m Map)
	GetID() int
	SetID(i int)
	GetLastMove() Point
	Debug(s string)
	Warning(s string)
	Error(s string)
	Log(s string)
}

//Player struct
type Player struct {
	SelfMap  Map
	EnemyMap Map
	LastMove Point
	ID       int
}

//ItoB converts int to bool
func ItoB(i int) bool {
	if i == 1 {
		return true
	}
	return false
}

//ItoA converts int to string
func ItoA(i int) string {
	return strconv.Itoa(i)
}

//AnswerString returns answer at string format
func AnswerString(i int) string {
	switch i {
	case Start:
		return "Start"
	case Miss:
		return "Miss"
	case Goal:
		return "Goal"
	case Kill:
		return "Kill"
	case End:
		return "End"
	case Move:
		return "Move"
	default:
		return "no answer"
	}
}

//String Turn
func (t *Turn) String() string {
	return fmt.Sprintf("X: %d Y: %d Answer: %s", t.Point.X, t.Point.Y, AnswerString(t.Answer))
}

//RandomByMap Point
func (point *Point) RandomByMap(m Map) {
	point.Random()
	for {
		if m.Field[point.X][point.Y] != Sea {
			point.Random()
		} else {
			break
		}
	}
}

//Random Point
func (point *Point) Random() {
	point.X = rand.Intn(10)
	point.Y = rand.Intn(10)
}

//Clear Point
func (point *Point) Clear() {
	point.X = -1
	point.Y = -1
}

//Debug Player
func (p *Player) Debug(s string) {
	if debugs {
		log.Println("DEBUG: " + s)
	}
}

//Warning Player
func (p *Player) Warning(s string) {
	if warnings {
		log.Println("WARNING: " + s)
	}
}

//Error Player
func (p *Player) Error(s string) {
	if errors {
		log.Println("ERROR: " + s)
	}
}

//Log Player
func (p *Player) Log(s string) {
	if logs {
		log.Println("LOG: " + s)
	}
}

//GetName Player
func (p *Player) GetName() string {
	return "player"
}

//GetSelfMap Player
func (p *Player) GetSelfMap() Map {
	return p.SelfMap
}

//SetSelfMap Player
func (p *Player) SetSelfMap(m Map) {
	p.SelfMap = m
}

//GetEnemyMap Player
func (p *Player) GetEnemyMap() Map {
	return p.EnemyMap
}

//SetEnemyMap Player
func (p *Player) SetEnemyMap(m Map) {
	p.EnemyMap = m
}

//GetID Player
func (p *Player) GetID() int {
	return p.ID
}

//SetID Player
func (p *Player) SetID(i int) {
	p.ID = i
}

//GetLastMove Player
func (p *Player) GetLastMove() Point {
	return p.LastMove
}

//MapStrategy Player
func (p *Player) MapStrategy() Map {
	p.Warning("MapStrategy not implemented")
	return Map{}
}

//AttackStrategy Player
func (p *Player) AttackStrategy(t Turn) Turn {
	p.Warning("AttackStrategy not implemented")
	return Turn{}
}

//StartStrategy for player
func StartStrategy(s Strategy, chTurn chan Turn, chMain chan int) {
	// init maps
	s.SetSelfMap(s.MapStrategy())
	selfMap := s.GetSelfMap()
	selfMap.Name = s.GetName() + "Self"

	enemyMap := s.GetEnemyMap()
	enemyMap.Clear()
	enemyMap.Name = s.GetName() + "Enemy"

	go func(chTurn chan Turn, chMain chan int) {
		// first player do first turn
		if s.GetID() == 1 {
			chTurn <- s.AttackStrategy(Turn{Point{-1, -1}, Start})
		}

		// process other turns
		for range time.Tick(time.Second) {
			// get other player turn
			turn := <-chTurn
			if turn.Answer == End {
				log.Println(s.GetName() + " WIN")
				close(chMain)
				return
			}

			// make answer
			turn = s.AttackStrategy(turn)
			// send answer
			chTurn <- turn
			if turn.Answer == End {
				return
			}
			// send to main channel
			chMain <- 0
		}
	}(chTurn, chMain)
}
