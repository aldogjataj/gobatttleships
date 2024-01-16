package main

import (
	"errors"
)

/*
This game of battleships is very simple to start:
There are 2 players
Each player has a grid which is 7*7
Each player has 9 Battleships, each of which can occupy only one square on their grid
Each player can place their battleships anywhere on this grid
Players take it in turns to pick any grid square reference
If the player hits a battleship, then it is sunk, and the turn passes to the opponent
If the player misses a battleship then it is called a miss, and the turn passes to the opponent
The player to first sink all their opponent's battleships is the winner
*/

// All code in here is example code, you do not have to keep any of it.

func PlayerOneTurn(playerTwoGrid [7][7]string, shotCoordinates []int) (shotStatus bool) {
	return false // shot missed
}

func PlayerTwoTurn(playerOneGrid [7][7]string, shotCoordinates []int) (shotStatus bool) {
	return true // shot hit
}
func checkCoordinatesAreInTheGrid(col, row int) error {
	if col < 0 || col > 6 || row < 0 || row > 6 {
		return errors.New("coordinates out of bounds")
	}
	return nil // Coordinates are valid
}

func CreateGrid() (grid [7][7]string) {
	// this is a fixed array, not a slice
	return [7][7]string{}
}

func countShips(grid [7][7]string) int {
	shipCount := 0
	for _, row := range grid {
		for _, cell := range row {
			if cell == "S" {
				shipCount++
			}
		}
	}
	return shipCount
}

func PlaceShip(grid [7][7]string, col, row int) ([7][7]string, error) {
	err := checkCoordinatesAreInTheGrid(col, row)
	if err != nil {
		return grid, err
	}

	shipCount := countShips(grid)
	if shipCount >= 9 {
		return grid, errors.New("Cannot place more than 9 ships.")
	}

	if grid[col][row] == "S" {
		return grid, errors.New("Two ships cannot be placed on the same place.")
	}

	grid[col][row] = "S"
	return grid, nil
}
func takeShot(grid [7][7]string, col, row int) ([7][7]string, string, error) {
	var result string
	err := checkCoordinatesAreInTheGrid(col, row)

	if err != nil {
		return grid, "", err
	}

	if grid[col][row] == "S" {
		grid[col][row] = "HIT"
		result = "HIT"
	} else {
		grid[col][row] = "X"
		result = "X"
	}
	return grid, result, nil
}
func isGameOver(grid [7][7]string, numberOfShips int) bool {
	numberOfShotShips := 0
	for _, row := range grid {
		for _, cell := range row {
			if cell == "HIT" {
				numberOfShotShips++
			}
		}
	}
	if numberOfShotShips == numberOfShips {
		return true
	}
	return false
}

func isShipAt(grid [7][7]string, col, row int) bool {
	return grid[col][row] == "S"
}
