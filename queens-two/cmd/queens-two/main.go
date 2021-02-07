package main

import (
	"fmt"
	"log"
	"math"

	"github.com/ssargent/queens-two/pathfinding"
)

func main() {
	fmt.Println("Queens Two")

	boardData := make([]string, 0)
	boardData = append(boardData, "20 3")
	boardData = append(boardData, "5 5")
	boardData = append(boardData, "10 10")
	boardData = append(boardData, "15 15")
	boardData = append(boardData, "16 13")

	board, err := pathfinding.NewBoard(boardData)
	if err != nil {
		log.Fatalf("Error creating board %v", err)
	}

	northSquares := board.NorthCount()
	southSquares := board.SouthCount()
	eastSquares := board.EastCount()
	westSquares := board.WestCount()
	northEastSquares := board.NorthEastCount()
	southEastSquares := board.SouthEastCount()
	southWestSquares := board.SouthWestCount()
	northWestSquares := board.NorthWestCount()

	totalVulnerable := northSquares + southSquares + eastSquares + westSquares + northEastSquares + southEastSquares + southWestSquares + northWestSquares
	/**/
	start := pathfinding.Point{X: 3, Y: 2}
	end := pathfinding.Point{X: 2, Y: 5}

	angle := math.Atan2(start.Y-end.Y, start.X-end.X)

	for _, p := range board.Blocks {
		angle := math.Atan2(board.Queen.Y-p.Y, board.Queen.X-p.X)
		resolvedAngle := (angle*(180/math.Pi) + 180)

		if math.Mod(resolvedAngle, 45) == 0 {
			fmt.Printf("Queen is blocked by point %f, %f\n", p.X, p.Y)
		}

	}

	fmt.Printf("Angle is %f\n", (angle*(180/math.Pi) + 180))
	fmt.Printf("Total Vulnerable Squares: %d\n", totalVulnerable)
}
