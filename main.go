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

func CreateGrid() (grid [7][7]string) {
	// this is a fixed array, not a slice
	return [7][7]string{}
}

func PlaceShip(grid [7][7]string, col, row int) ([7][7]string, error) {
	shipCount := 0
	for _, row := range grid {
		for _, cell := range row {
			if cell == "S" {
				shipCount++
			}
		}
	}

	if shipCount >= 9 {
		return grid, errors.New("Cannot place more than 9 ships")
	}
	if col >= 7 || col < 0 || row >= 7 || row < 0 {
		return grid, errors.New("Out of bounds.")
	}
	if grid[col][row] == "S" {
		return grid, errors.New("Two ships cannot be placed on the same place")
	}
	grid[col][row] = "S"
	return grid, nil
}
func takeShot(grid [7][7]string, col, row int) ([7][7]string, string, error) {
	var result string
	if col >= 7 || col < 0 || row >= 7 || row < 0 {
		return grid, "", errors.New("Out of bounds.")
	}

	if grid[col][row] == "S" {
		grid[col][row] = "HIT"
		result = "HIT"
	} else {
		grid[col][row] = "MISS"
		result = "MISS"
	}
	return grid, result, nil
}
