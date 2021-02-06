package main

import (
	"fmt"
	"log"

	"github.com/ssargent/queens-two/pathfinding"
)

func main() {
	fmt.Println("Queens Two")
	board, err := pathfinding.NewBoard([]string{})
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

	fmt.Printf("Total Vulnerable Squares: %d\n", totalVulnerable)
}
