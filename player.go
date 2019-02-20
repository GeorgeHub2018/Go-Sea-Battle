package main

import "math/rand"

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
	SelfMap  Map
	EnemyMap Map
	LastMove Point
}

//PlayerRandom struct
type PlayerRandom struct {
	Player
}

//PlayerAlgorithm struct
type PlayerAlgorithm struct {
	Player
}

//MapStrategy PlayerRandom
func (p *PlayerRandom) MapStrategy() Map {
	m := Map{}
	m.Clear()
	boatList := BoatList{}
	boatList.Init()
	for boatSize, boatCount := range boatList.List {
		for count := 0; count < boatCount; count++ {
			isHorizontal := rand.Intn(2)
			for {
				if isHorizontal == 1 {
					x := rand.Intn(10)
					y := rand.Intn(10)
					if x+boatSize > 9 {
						continue
					}

					isCanPlace := true
					for i := x - 1; i < x+1+boatSize; i++ {
						for j := y - 1; j <= y+1; j++ {
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
						for i := x; i < x+boatSize; i++ {
							m.Field[i][y] = Boat
						}
						break
					}
				} else {
					x := rand.Intn(10)
					y := rand.Intn(10)
					if y+boatSize > 9 {
						continue
					}

					isCanPlace := true
					for j := y - 1; j < y+1+boatSize; j++ {
						for i := x - 1; i <= x+1; i++ {
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
						for j := y; j < y+boatSize; j++ {
							m.Field[x][j] = Boat
						}
						break
					}
				}
			}
		}
	}

	return m
}

//AttackStrategy PlayerRandom
func (p *PlayerRandom) AttackStrategy(t Turn) Turn {
	//TODO:
	return t
}

//MapStrategy PlayerAlgorithm
func (p *PlayerAlgorithm) MapStrategy() Map {
	m := Map{}
	m.Clear()
	boatList := BoatList{}
	boatList.Init()
	for boatSize, boatCount := range boatList.List {
		for count := 0; count < boatCount; count++ {
			isHorizontal := rand.Intn(2)
			for {
				if isHorizontal == 1 {
					x := rand.Intn(10)
					y := rand.Intn(10)
					if x+boatSize > 9 {
						continue
					}

					isCanPlace := true
					for i := x - 1; i < x+1+boatSize; i++ {
						for j := y - 1; j <= y+1; j++ {
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
						for i := x; i < x+boatSize; i++ {
							m.Field[i][y] = Boat
						}
						break
					}
				} else {
					x := rand.Intn(10)
					y := rand.Intn(10)
					if y+boatSize > 9 {
						continue
					}

					isCanPlace := true
					for j := y - 1; j < y+1+boatSize; j++ {
						for i := x - 1; i <= x+1; i++ {
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
						for j := y; j < y+boatSize; j++ {
							m.Field[x][j] = Boat
						}
						break
					}
				}
			}
		}
	}

	return m
}

//AttackStrategy PlayerAlgorithm
func (p *PlayerAlgorithm) AttackStrategy(t Turn) Turn {
	turn := Turn{Point{-1, -1}, Start}
	switch t.Answer {
	case Start:
		turn.Point.X = rand.Intn(10)
		turn.Point.Y = rand.Intn(10)
		turn.Answer = Move
	case Miss:
		if p.SelfMap.Field[t.Point.X][t.Point.Y] == Boat {
			turn.Answer = Goal
			p.LastMove = turn.Point
		} else {
			turn.Point.X = rand.Intn(10)
			turn.Point.Y = rand.Intn(10)
			turn.Answer = Miss
		}
	case Goal:
		p.EnemyMap.Field[p.LastMove.X][p.LastMove.Y] = HittedBoat
		turn.Point.X = rand.Intn(10)
		turn.Point.Y = rand.Intn(10)
		turn.Answer = Move
	case Kill:
		p.EnemyMap.Field[p.LastMove.X][p.LastMove.Y] = HittedBoat
		//TODO: add padding
		turn.Point.X = rand.Intn(10)
		turn.Point.Y = rand.Intn(10)
		turn.Answer = Move
	case End:
		turn.Answer = End
	case Move:
		if p.SelfMap.Field[t.Point.X][t.Point.Y] == Boat {
			turn.Answer = Goal
		} else {
			turn.Point.X = rand.Intn(10)
			turn.Point.Y = rand.Intn(10)
			turn.Answer = Miss
		}
	}

	p.LastMove = turn.Point
	return turn
}
