<<<<<<< HEAD
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
=======
package main

import (
	"fmt"
	"math/rand"
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
}

//Player struct
type Player struct {
	Name     string
	SelfMap  Map
	EnemyMap Map
	LastMove Point
}

//PlayerRandom struct
type PlayerRandom struct {
	Player
	Name string "Random"
}

//PlayerAlgorithm struct
type PlayerAlgorithm struct {
	Player
	Name string "Algorithm"
}

//ItoB converts int to bool
func ItoB(i int) bool {
	if i == 1 {
		return true
	}
	return false
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
	if debug {
		fmt.Println("DEBUG " + s)
	}
}

//MapStrategy PlayerRandom
func (p *PlayerRandom) MapStrategy() Map {
	p.Debug("MapStrategy PlayerRandom begin")
	defer p.Debug("MapStrategy PlayerRandom end")

	m := Map{}
	m.Clear()
	m.Boats = BoatList{}
	m.Boats.Init()
	for _, boat := range m.Boats.List {
		boat.IsHorizontal = ItoB(rand.Intn(2))
		for {
			if boat.IsHorizontal == true {
				boat.LeftTopPoint.Random()
				if boat.LeftTopPoint.X+boat.Size > 9 {
					continue
				}

				isCanPlace := true
				for i := boat.LeftTopPoint.X - 1; i < boat.LeftTopPoint.X+1+boat.Size; i++ {
					for j := boat.LeftTopPoint.Y - 1; j <= boat.LeftTopPoint.Y+1; j++ {
						if i < 0 {
							continue
						}
						if i > 9 {
							continue
						}
						if j < 0 {
							continue
						}
						if j > 9 {
							continue
						}

						if m.Field[i][j] == Boat {
							isCanPlace = false
							break
						}
					}
					if isCanPlace == false {
						break
					}
				}
				if isCanPlace == true {
					for i := boat.LeftTopPoint.X; i < boat.LeftTopPoint.X+boat.Size; i++ {
						m.Field[i][boat.LeftTopPoint.Y] = Boat
					}
					break
				}
			} else {
				boat.LeftTopPoint.Random()
				if boat.LeftTopPoint.Y+boat.Size > 9 {
					continue
				}

				isCanPlace := true
				for j := boat.LeftTopPoint.Y - 1; j < boat.LeftTopPoint.Y+1+boat.Size; j++ {
					for i := boat.LeftTopPoint.X - 1; i <= boat.LeftTopPoint.X+1; i++ {
						if i < 0 {
							continue
						}
						if i > 9 {
							continue
						}
						if j < 0 {
							continue
						}
						if j > 9 {
							continue
						}
						if m.Field[i][j] == Boat {
							isCanPlace = false
							break
						}
					}
					if isCanPlace == false {
						break
					}
				}
				if isCanPlace == true {
					for j := boat.LeftTopPoint.Y; j < boat.LeftTopPoint.Y+boat.Size; j++ {
						m.Field[boat.LeftTopPoint.X][j] = Boat
					}
					break
				}
			}
		}
	}

	return m
}

//AttackStrategy PlayerRandom
func (p *PlayerRandom) AttackStrategy(t Turn) Turn {
	p.Debug("AttackStrategy PlayerRandom begin")
	defer p.Debug("AttackStrategy PlayerRandom end")
	turn := Turn{Point{-1, -1}, Start}
	switch t.Answer {
	case Start:
		turn.Point.RandomByMap(p.EnemyMap)
		turn.Answer = Move
	case Miss:
		if p.SelfMap.Field[t.Point.X][t.Point.Y] == Boat {
			if p.SelfMap.BoatKilled(t.Point.X, t.Point.Y) {
				turn.Answer = Kill
				fmt.Println(p.Name + "kill")
			} else {
				turn.Answer = Goal
			}
		} else {
			turn.Point.RandomByMap(p.EnemyMap)
			turn.Answer = Miss
		}
		if (p.LastMove.X >= 0) && (p.LastMove.Y >= 0) {
			p.EnemyMap.Field[p.LastMove.X][p.LastMove.Y] = Hitted
		}
	case Goal:
		turn.Point.RandomByMap(p.EnemyMap)
		turn.Answer = Move
		if (p.LastMove.X >= 0) && (p.LastMove.Y >= 0) {
			p.EnemyMap.Field[p.LastMove.X][p.LastMove.Y] = HittedBoat
		}
	case Kill:
		boat := p.EnemyMap.FindOrMakeBoat(p.LastMove.X, p.LastMove.Y)
		p.EnemyMap.MakePadding(boat)
		turn.Point.RandomByMap(p.EnemyMap)
		turn.Answer = Move
		if (p.LastMove.X >= 0) && (p.LastMove.Y >= 0) {
			p.EnemyMap.Field[p.LastMove.X][p.LastMove.Y] = HittedBoat
		}
	case End:
		turn.Answer = End
	case Move:
		if p.SelfMap.Field[t.Point.X][t.Point.Y] == Boat {
			if p.SelfMap.BoatKilled(t.Point.X, t.Point.Y) {
				turn.Answer = Kill
			} else {
				turn.Answer = Goal
			}
		} else {
			turn.Point.RandomByMap(p.EnemyMap)
			turn.Answer = Miss
		}
	}

	p.LastMove = turn.Point
	return turn
}

//MapStrategy PlayerAlgorithm
func (p *PlayerAlgorithm) MapStrategy() Map {
	p.Debug("MapStrategy PlayerAlgorithm begin")
	defer p.Debug("MapStrategy PlayerAlgorithm end")
	m := Map{}
	m.Clear()
	m.Boats = BoatList{}
	m.Boats.Init()
	for _, boat := range m.Boats.List {
		//for count := 0; count < boatCount; count++ {
		boat.IsHorizontal = ItoB(rand.Intn(2))
		for {
			//p.Debug("cycle")
			if boat.IsHorizontal == true {
				boat.LeftTopPoint.Random()
				if boat.LeftTopPoint.X+boat.Size > 9 {
					continue
				}

				isCanPlace := true
				for i := boat.LeftTopPoint.X - 1; i < boat.LeftTopPoint.X+1+boat.Size; i++ {
					for j := boat.LeftTopPoint.Y - 1; j <= boat.LeftTopPoint.Y+1; j++ {
						if i < 0 {
							continue
						}
						if i > 9 {
							continue
						}
						if j < 0 {
							continue
						}
						if j > 9 {
							continue
						}

						if m.Field[i][j] == Boat {
							isCanPlace = false
							break
						}
					}
					if isCanPlace == false {
						break
					}
				}
				if isCanPlace == true {
					for i := boat.LeftTopPoint.X; i < boat.LeftTopPoint.X+boat.Size; i++ {
						m.Field[i][boat.LeftTopPoint.Y] = Boat
					}
					break
				}
			} else {
				boat.LeftTopPoint.Random()
				if boat.LeftTopPoint.Y+boat.Size > 9 {
					continue
				}

				isCanPlace := true
				for j := boat.LeftTopPoint.Y - 1; j < boat.LeftTopPoint.Y+1+boat.Size; j++ {
					for i := boat.LeftTopPoint.X - 1; i <= boat.LeftTopPoint.X+1; i++ {
						if i < 0 {
							continue
						}
						if i > 9 {
							continue
						}
						if j < 0 {
							continue
						}
						if j > 9 {
							continue
						}
						if m.Field[i][j] == Boat {
							isCanPlace = false
							break
						}
					}
					if isCanPlace == false {
						break
					}
				}
				if isCanPlace == true {
					for j := boat.LeftTopPoint.Y; j < boat.LeftTopPoint.Y+boat.Size; j++ {
						m.Field[boat.LeftTopPoint.X][j] = Boat
					}
					break
				}
			}
		}
		//}
	}

	return m
}

//AttackStrategy PlayerAlgorithm
func (p *Player) AttackStrategy(t Turn) Turn {
	p.Debug("AttackStrategy PlayerRandom begin")
	defer p.Debug("AttackStrategy PlayerRandom end")
	turn := Turn{Point{-1, -1}, Start}
	switch t.Answer {
	case Start:
		turn.Point.RandomByMap(p.EnemyMap)
		turn.Answer = Move
	case Miss:
		if p.SelfMap.Field[t.Point.X][t.Point.Y] == Boat {
			if p.SelfMap.BoatKilled(t.Point.X, t.Point.Y) {
				turn.Answer = Kill
				fmt.Println(p.Name + "kill")
			} else {
				turn.Answer = Goal
			}
		} else {
			turn.Point.RandomByMap(p.EnemyMap)
			turn.Answer = Miss
		}
		if (p.LastMove.X >= 0) && (p.LastMove.Y >= 0) {
			p.EnemyMap.Field[p.LastMove.X][p.LastMove.Y] = Hitted
		}
	case Goal:
		turn.Point.RandomByMap(p.EnemyMap)
		turn.Answer = Move
		if (p.LastMove.X >= 0) && (p.LastMove.Y >= 0) {
			p.EnemyMap.Field[p.LastMove.X][p.LastMove.Y] = HittedBoat
		}
	case Kill:
		boat := p.EnemyMap.FindOrMakeBoat(p.LastMove.X, p.LastMove.Y)
		p.EnemyMap.MakePadding(boat)
		turn.Point.RandomByMap(p.EnemyMap)
		turn.Answer = Move
		if (p.LastMove.X >= 0) && (p.LastMove.Y >= 0) {
			p.EnemyMap.Field[p.LastMove.X][p.LastMove.Y] = HittedBoat
		}
	case End:
		turn.Answer = End
	case Move:
		if p.SelfMap.Field[t.Point.X][t.Point.Y] == Boat {
			if p.SelfMap.BoatKilled(t.Point.X, t.Point.Y) {
				turn.Answer = Kill
			} else {
				turn.Answer = Goal
			}
		} else {
			turn.Point.RandomByMap(p.EnemyMap)
			turn.Answer = Miss
		}
	}

	p.LastMove = turn.Point
	return turn
}
>>>>>>> 09d48927fc4507cb93a42304ec539761626aa40b
