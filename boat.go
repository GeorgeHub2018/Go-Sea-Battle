package main

import "fmt"

type (
	//BoatStruct struct
	BoatStruct struct {
		IsHorizontal bool
		LeftTopPoint Point
		Size         int
		IsDead       bool
	}
	//BoatList map of boats
	BoatList struct {
		List [10]BoatStruct
	}
)

//IsExist BoatStruct
func (b *BoatStruct) IsExist() bool {
	if (b.LeftTopPoint.X == -1) || (b.LeftTopPoint.Y == -1) {
		return false
	}
	return true
}

//String BoatStruct
func (b *BoatStruct) String() string {
	return fmt.Sprintf("LeftTopX: %d LeftTopY: %d Size: %d, IsHor: %t IsDead: %t",
		b.LeftTopPoint.X, b.LeftTopPoint.Y, b.Size, b.IsHorizontal, b.IsDead)
}

//Clear Boat
func (b *BoatStruct) Clear() {
	b.IsHorizontal = false
	b.LeftTopPoint.X = -1
	b.LeftTopPoint.Y = -1
	b.IsDead = false
}

//Init Boat
func (b *BoatStruct) Init(size int) {
	b.Clear()
	b.Size = size
}

//Init BoatList
func (bl *BoatList) Init() {
	bl.List[0].Init(4)
	bl.List[1].Init(3)
	bl.List[2].Init(3)
	bl.List[3].Init(2)
	bl.List[4].Init(2)
	bl.List[5].Init(2)
	bl.List[6].Init(1)
	bl.List[7].Init(1)
	bl.List[8].Init(1)
	bl.List[9].Init(1)
}
