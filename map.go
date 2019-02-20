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
		Field [10][10]int
	}

	//BoatList map of boats
	BoatList struct {
		List map[int]int
	}
)

//Init BoatList
func (b *BoatList) Init() {
	b.List = make(map[int]int)
	b.List[4] = 1
	b.List[3] = 2
	b.List[2] = 3
	b.List[1] = 4
}

//Clear Map
func (m *Map) Clear() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			m.Field[i][j] = Sea
		}
	}
}

//GetMarkerStr makes marker string
func GetMarkerStr(marker int) string {
	switch marker {
	case Sea:
		return " "
	case Hitted:
		return "*"
	case Boat:
		return "o"
	case HittedBoat:
		return "x"
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
