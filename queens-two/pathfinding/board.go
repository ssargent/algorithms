package pathfinding

import (
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

type Board struct {
	Size   int
	Queen  Point
	Blocks []Point
}

func NewBoard(config []string) (*Board, error) {
	board := Board{}

	// first row has our size.
	sizeRow := config[0]
	parts := strings.Split(sizeRow, " ")
	tmpSize, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, err
	}
	board.Size = tmpSize

	// second row has our queen
	queenRow := config[1]
	queenParts := strings.Split(queenRow, " ")
	tmpX, err := strconv.Atoi(queenParts[0])
	if err != nil {
		return nil, err
	}

	tmpY, err := strconv.Atoi(queenParts[1])
	if err != nil {
		return nil, err
	}

	board.Queen = Point{X: tmpX, Y: tmpY}

	for x := 2; x < len(config); x++ {
		blockRow := config[x]
		blockParts := strings.Split(blockRow, " ")
		tmpX, err := strconv.Atoi(blockParts[0])
		if err != nil {
			return nil, err
		}

		tmpY, err := strconv.Atoi(blockParts[1])
		if err != nil {
			return nil, err
		}

		board.Blocks = append(board.Blocks, Point{X: tmpX, Y: tmpY})
	}

	return &board, nil
}

func pointOnLine(start Point, end Point, blocks *[]Point) *Point {
	if isDiagonal(start, end) {
		return pointOnDiagonalLine(start, end, blocks)
	}
	return nil
}

func isDiagonal(start Point, end Point) bool {
	return false
}

func pointOnDiagonalLine(start Point, end Point, blocks *[]Point) *Point {
	return nil
}

func (b *Board) NorthCount() int {
	return 0
}

func (b *Board) SouthCount() int {
	return 0
}

func (b *Board) EastCount() int {
	return 0
}

func (b *Board) WestCount() int {
	return 0
}

func (b *Board) NorthEastCount() int {
	return 0
}

func (b *Board) NorthWestCount() int {
	return 0
}

func (b *Board) SouthEastCount() int {
	return 0
}

func (b *Board) SouthWestCount() int {
	return 0
}
