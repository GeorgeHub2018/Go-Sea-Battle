package main

import (
	"fmt"
	"strconv"
)

//Map markers
const (
	Sea = iota
	Hitted
	Boat
	HittedBoat
)

type (
	//Map struct
	Map struct {
		Name  string
		Field [10][10]int
		Boats BoatList
	}
)

//Clear Map
func (m *Map) Clear() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			m.Field[i][j] = Sea
		}
	}
}

//IsBoatDead Map
func (m *Map) IsBoatDead(boat BoatStruct) bool {
	if boat.IsHorizontal == true {
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
					return false
				}
			}
		}
		return true
	}

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
				return false
			}
		}
	}
	return true
}

//MakePadding Map
func (m *Map) MakePadding(boat BoatStruct) {
	if boat.IsHorizontal == true {
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
				if (i >= boat.LeftTopPoint.X) && (i <= boat.LeftTopPoint.X+boat.Size-1) &&
					(j == boat.LeftTopPoint.Y) {
					continue
				}

				m.Field[i][j] = Hitted
			}
		}
	}

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
			if (i >= boat.LeftTopPoint.X) && (i <= boat.LeftTopPoint.X+boat.Size-1) &&
				(j == boat.LeftTopPoint.Y) {
				continue
			}

			m.Field[i][j] = Hitted
		}
	}
}

//IsBoatZone Map
func (m *Map) IsBoatZone(boat BoatStruct, x int, y int) bool {
	// Is horizontal
	if boat.IsHorizontal {
		if (x >= boat.LeftTopPoint.X) && (x <= boat.LeftTopPoint.Y+boat.Size) &&
			(y == boat.LeftTopPoint.Y) {
			return true
		}
		return false
	}

	// Is vertical
	if (y >= boat.LeftTopPoint.Y) && (y <= boat.LeftTopPoint.Y+boat.Size) &&
		(x == boat.LeftTopPoint.X) {
		return true
	}
	return false
}

//FindBoat Map
func (m *Map) FindBoat(x, y int) (b BoatStruct) {
	for _, boat := range m.Boats.List {
		if !boat.IsExist() {
			continue
		}

		if m.IsBoatZone(boat, x, y) {
			return boat
		}
	}
	return
}

//MakeBoat Map
func (m *Map) MakeBoat(x, y int) (b BoatStruct) {
	b.Clear()
	b.LeftTopPoint.X = x
	b.LeftTopPoint.Y = y
	// find left top
	for {
		if b.LeftTopPoint.X > 0 {
			if m.Field[b.LeftTopPoint.X-1][b.LeftTopPoint.Y] == Boat {
				b.IsHorizontal = true
				b.LeftTopPoint.X--
				continue
			}
		}
		break
	}

	for {
		if b.LeftTopPoint.Y > 0 {
			if m.Field[b.LeftTopPoint.X][b.LeftTopPoint.Y-1] == Boat {
				b.IsHorizontal = false
				b.LeftTopPoint.Y--
				continue
			}
		}
		break
	}
	// calc size
	b.Size = 0
	if b.IsHorizontal {
		for i := b.LeftTopPoint.X; i < 10; i++ {
			if (m.Field[i][b.LeftTopPoint.Y] == Boat) || (m.Field[i][b.LeftTopPoint.Y] == HittedBoat) {
				b.Size++
				continue
			}
			break
		}
	}

	for j := b.LeftTopPoint.Y; j < 10; j++ {
		if (m.Field[b.LeftTopPoint.X][j] == Boat) || (m.Field[b.LeftTopPoint.X][j] == HittedBoat) {
			b.Size++
			continue
		}
		break
	}
	return
}

//FindOrMakeBoat Map
func (m *Map) FindOrMakeBoat(x, y int) (b BoatStruct) {
	b = m.FindBoat(x, y)
	if !b.IsExist() {
		b = m.MakeBoat(x, y)
	}
	return
}

//MakeAutoPadding Map
func (m *Map) MakeAutoPadding() {
	for _, boat := range m.Boats.List {
		if boat.IsDead {
			continue
		}

		if m.IsBoatDead(boat) {
			boat.IsDead = true
			m.MakePadding(boat)
		}
	}
}

//BoatKilled Map
func (m *Map) BoatKilled(x, y int) bool {
	boat := m.FindBoat(x, y)
	if boat.IsExist() {
		if boat.IsHorizontal == true {
			for i := boat.LeftTopPoint.X; i < boat.LeftTopPoint.X+boat.Size; i++ {
				if i < 0 {
					continue
				}

				if m.Field[i][boat.LeftTopPoint.Y] != HittedBoat {
					fmt.Println("BoatKilled false", boat, x, y)
					return false
				}
			}
		}

		for j := boat.LeftTopPoint.Y - 1; j < boat.LeftTopPoint.Y+1+boat.Size; j++ {
			if j < 0 {
				continue
			}

			if m.Field[boat.LeftTopPoint.X][j] != HittedBoat {
				fmt.Println("BoatKilled false", boat, x, y)
				return false
			}
		}
	}

	fmt.Println("BoatKilled true")
	return true
}

//GetMarkerStr makes marker string
func GetMarkerStr(marker int) string {
	switch marker {
	case Sea:
		return " "
	case Hitted:
		return "*"
	case Boat:
		return "O"
	case HittedBoat:
		return "X"
	default:
		return " "
	}
}

//PrintMaps print maps
func PrintMaps(m1Own Map, m1Enemy Map, m2Own Map, m2Enemy Map) {
	fmt.Println("Player1Own    Player1Enemy    Player2Own    Player2Enemy")
	fmt.Println(" |ABCDEFGHIJ   |ABCDEFGHIJ     |ABCDEFGHIJ   |ABCDEFGHIJ")
	fmt.Println("------------  ------------    ------------  ------------")
	for j := 0; j < 10; j++ {
		s1 := strconv.Itoa(j) + "|"
		s2 := strconv.Itoa(j) + "|"
		s3 := strconv.Itoa(j) + "|"
		s4 := strconv.Itoa(j) + "|"
		for i := 0; i < 10; i++ {
			s1 = s1 + GetMarkerStr(m1Own.Field[i][j])
			s2 = s2 + GetMarkerStr(m1Enemy.Field[i][j])
			s3 = s3 + GetMarkerStr(m2Own.Field[i][j])
			s4 = s4 + GetMarkerStr(m2Enemy.Field[i][j])
		}
		fmt.Println(s1 + "  " + s2 + "    " + s3 + "  " + s4)
	}
}
