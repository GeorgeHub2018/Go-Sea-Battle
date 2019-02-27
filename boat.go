<<<<<<< HEAD
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
=======
package main

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
		List [10]BoatStruct //map[int]BoatStruct
	}
)

//IsExist BoatStruct
func (b *BoatStruct) IsExist() bool {
	if (b.LeftTopPoint.X == -1) || (b.LeftTopPoint.Y == -1) {
		return false
	}
	return true
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
	//b.List = []BoatStruct //make(map[int]BoatStruct)
	bl.List[0].Init(4) // = BoatStruct{false, Point{-1, -1}, 4, false}
	bl.List[1].Init(3) // = BoatStruct{false, Point{-1, -1}, 3, false}
	bl.List[2].Init(3) // = BoatStruct{false, Point{-1, -1}, 3, false}
	bl.List[3].Init(2) // = BoatStruct{false, Point{-1, -1}, 2, false}
	bl.List[4].Init(2) // = BoatStruct{false, Point{-1, -1}, 2, false}
	bl.List[5].Init(2) // = BoatStruct{false, Point{-1, -1}, 2, false}
	bl.List[6].Init(1) // = BoatStruct{false, Point{-1, -1}, 1, false}
	bl.List[7].Init(1) // = BoatStruct{false, Point{-1, -1}, 1, false}
	bl.List[8].Init(1) // = BoatStruct{false, Point{-1, -1}, 1, false}
	bl.List[9].Init(1) // = BoatStruct{false, Point{-1, -1}, 1, false}
}
>>>>>>> 09d48927fc4507cb93a42304ec539761626aa40b
