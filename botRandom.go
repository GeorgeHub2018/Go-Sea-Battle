package main

import (
	"math/rand"
)

//BotRandom struct
type BotRandom struct {
	Player
}

//GetName BotRandom
func (p *BotRandom) GetName() string {
	return "BotRandom" + ItoA(p.ID)
}

//MapStrategy BotRandom
func (p *BotRandom) MapStrategy() Map {
	p.Debug("MapStrategy " + p.GetName() + " begin")

	m := Map{}
	m.Clear()
	m.Boats = BoatList{}
	m.Boats.Init()
	for i, boat := range m.Boats.List {
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
		m.Boats.List[i] = boat
	}

	p.Debug("MapStrategy " + p.GetName() + " end")

	return m
}

//AttackStrategy BotRandom
func (p *BotRandom) AttackStrategy(t Turn) Turn {
	p.Debug("AttackStrategy " + p.GetName() + " begin " + t.String())

	turn := Turn{Point{-1, -1}, Start}
	switch t.Answer {
	//Start
	case Start:
		turn.Point.RandomByMap(p.EnemyMap)
		turn.Answer = Move
	//Miss
	case Miss:
		if p.SelfMap.Field[t.Point.X][t.Point.Y] == Boat {
			p.SelfMap.Field[t.Point.X][t.Point.Y] = HittedBoat
			if p.SelfMap.IsBoatKilled(t.Point.X, t.Point.Y) {
				if p.SelfMap.IsAllBoatsKilled() {
					turn.Answer = End
				} else {
					turn.Answer = Kill
				}
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
	//Goal
	case Goal:
		if (p.LastMove.X >= 0) && (p.LastMove.Y >= 0) {
			p.EnemyMap.Field[p.LastMove.X][p.LastMove.Y] = HittedBoat
		}
		turn.Point.RandomByMap(p.EnemyMap)
		turn.Answer = Move
	//Kill
	case Kill:
		if (p.LastMove.X >= 0) && (p.LastMove.Y >= 0) {
			p.EnemyMap.Field[p.LastMove.X][p.LastMove.Y] = HittedBoat
		}
		boat := p.EnemyMap.FindOrMakeBoat(p.LastMove.X, p.LastMove.Y)
		p.EnemyMap.MakePadding(boat)
		turn.Point.RandomByMap(p.EnemyMap)
		turn.Answer = Move
	//End
	case End:
		turn.Answer = End
	//Move
	case Move:
		if p.SelfMap.Field[t.Point.X][t.Point.Y] == Boat {
			p.SelfMap.Field[t.Point.X][t.Point.Y] = HittedBoat
			if p.SelfMap.IsBoatKilled(t.Point.X, t.Point.Y) {
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
	p.Debug("AttackStrategy " + p.GetName() + " end " + turn.String())
	return turn
}
