package main

import (
	"errors"
	"fmt"
	"math/rand"
	"testing"
)

//you can run all you tests by typing
//go test -v
//in the terminal window

// this is a utility function for testing
// it will return a random square on the grid
// it does not keep track of any previously returned grids
func getRandomGridSquare() []int {

	row := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	column := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	return []int{rand.Intn(len(row)) + 1, rand.Intn(len(column)) + 1}

}

// these are the two tests we have for our functions in main
// the purpose of tests is to mimic interaction with our code
// there is no "user input" - the test is the calling code
func TestPlaceAShip(t *testing.T) {
	//Arrange
	grid := CreateGrid()
	//Act
	desiredCol := 3
	desiredRow := 5
	updatedGrid, _ := PlaceShip(grid, desiredCol, desiredRow)
	//Assert
	actual := isShipAt(updatedGrid, 3, 5)
	want := true
	if actual != want {
		t.Error("Ship was not placed at [3,5]")
	}
}
func TestCanPlaceNineShips(t *testing.T) {
	//Arrange
	grid := CreateGrid()
	grid, _ = PlaceShip(grid, 1, 0)
	grid, _ = PlaceShip(grid, 1, 1)
	grid, _ = PlaceShip(grid, 1, 2)
	grid, _ = PlaceShip(grid, 2, 0)
	grid, _ = PlaceShip(grid, 2, 1)
	grid, _ = PlaceShip(grid, 2, 2)
	grid, _ = PlaceShip(grid, 3, 0)
	grid, _ = PlaceShip(grid, 3, 1)
	grid, _ = PlaceShip(grid, 3, 2)

	//Act
	shipCount := countShips(grid)
	//Assert
	if shipCount != 9 {
		t.Errorf("Expected 9 ships, not %d", shipCount)
	}
}

func TestCannotPlaceMoreThanNineShips(t *testing.T) {
	// Arrange
	grid := CreateGrid()
	grid, _ = PlaceShip(grid, 0, 0)
	grid, _ = PlaceShip(grid, 1, 1)
	grid, _ = PlaceShip(grid, 2, 2)
	grid, _ = PlaceShip(grid, 3, 3)
	grid, _ = PlaceShip(grid, 4, 4)
	grid, _ = PlaceShip(grid, 5, 5)
	grid, _ = PlaceShip(grid, 6, 6)
	grid, _ = PlaceShip(grid, 0, 1)
	grid, _ = PlaceShip(grid, 1, 2)

	// Act
	_, got := PlaceShip(grid, 3, 4)
	// Assert
	want := errors.New("Cannot place more than 9 ships.")
	if got.Error() != want.Error() {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestAttemptToPlaceTenthShipDoesntChangeGrid(t *testing.T) {
	//Arrange
	grid := CreateGrid()
	grid, _ = PlaceShip(grid, 0, 0)
	grid, _ = PlaceShip(grid, 1, 1)
	grid, _ = PlaceShip(grid, 2, 2)
	grid, _ = PlaceShip(grid, 3, 3)
	grid, _ = PlaceShip(grid, 4, 4)
	grid, _ = PlaceShip(grid, 5, 5)
	grid, _ = PlaceShip(grid, 6, 6)
	grid, _ = PlaceShip(grid, 0, 1)
	grid, _ = PlaceShip(grid, 1, 2)
	//Act
	newGrid, _ := PlaceShip(grid, 3, 5)
	//Assert
	if newGrid != grid {
		t.Error("Tenth ship was placed.")
	}

}
func TestCannotStackShips(t *testing.T) {
	// Arrange
	grid := CreateGrid()
	grid, _ = PlaceShip(grid, 3, 5)

	// Act
	_, got := PlaceShip(grid, 3, 5)

	// Assert
	want := errors.New("Two ships cannot be placed on the same place.")
	if got.Error() != want.Error() {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestCanPlaceShipAtx0(t *testing.T) {
	//Arrange
	grid := CreateGrid()
	//Act
	_, err := PlaceShip(grid, 0, 1)
	//Assert
	if err != nil {
		t.Error("Cannot place a ship at x=0.")
	}
}
func TestCanPlaceShipAtY0(t *testing.T) {
	//Arrange
	grid := CreateGrid()
	//Act
	_, err := PlaceShip(grid, 1, 0)
	//Assert
	if err != nil {
		t.Error("Cannot place a ship at y=0.")
	}
}
func TestCanPlaceShipAtX6(t *testing.T) {
	//Arrange
	grid := CreateGrid()
	//Act
	_, err := PlaceShip(grid, 6, 1)
	//Assert
	if err != nil {
		t.Error("Cannot place a ship at x=6.")
	}

}
func TestCanPlaceShipAtY6(t *testing.T) {
	//Arrange
	grid := CreateGrid()
	//Act
	_, err := PlaceShip(grid, 1, 6)
	//Assert
	if err != nil {
		t.Error("Cannot place a ship at y=6.")
	}
}

func TestCannotPlaceShipAtX7(t *testing.T) {
	//Arrange
	grid := CreateGrid()
	//Act
	_, got := PlaceShip(grid, 7, 2)
	//Assert
	want := errors.New("coordinates out of bounds")
	if got.Error() != want.Error() {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestCannotPlaceShipAtY7(t *testing.T) {
	//Arrange
	grid := CreateGrid()
	//Act
	_, got := PlaceShip(grid, 2, 7)
	//Assert
	want := errors.New("coordinates out of bounds")
	if got.Error() != want.Error() {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestCannotPlaceShipAtNegativeX(t *testing.T) {
	//Act
	grid := CreateGrid()
	//Act
	_, got := PlaceShip(grid, -1, 5)
	//Assert
	want := errors.New("coordinates out of bounds")
	if got.Error() != want.Error() {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestCannotPlaceShipAtNegativeY(t *testing.T) {
	//Act
	grid := CreateGrid()
	//Act
	_, got := PlaceShip(grid, 1, -1)
	//Assert
	want := errors.New("coordinates out of bounds")
	if got.Error() != want.Error() {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestReportShotTaken(t *testing.T) {
	// Arrange
	grid := CreateGrid()

	// Act
	_, result, _ := takeShot(grid, 3, 5)

	// Assert
	want := "MISS"
	if result != want {
		t.Errorf("Shot was not recorded.")
	}
}

func TestReportsShipIsHit(t *testing.T) {
	// Arrange
	grid := CreateGrid()
	grid, _ = PlaceShip(grid, 3, 5)

	// Act
	_, result, _ := takeShot(grid, 3, 5)

	// Assert
	want := "HIT"
	if result != want {
		t.Error("Ship was not hit!")
	}
}

func TestShotAtSunkShipReportsMiss(t *testing.T) {
	// Arrange
	grid := CreateGrid()
	grid, _ = PlaceShip(grid, 3, 5)
	grid, _, _ = takeShot(grid, 3, 5)

	// Act
	_, result, _ := takeShot(grid, 3, 5)

	// Assert
	want := "MISS"
	if result != want {
		t.Error("Shooting at a sunk ship still shows as a hit.")
	}
}
func TestCanShootAtX6(t *testing.T) {
	//Arrange
	grid := CreateGrid()
	//Act
	_, _, got := takeShot(grid, 6, 1)
	//Assert

	if got != nil {
		t.Error("Cannot shoot at x=6.")
	}
}

func TestCanShootAtY6(t *testing.T) {
	//Arrange
	grid := CreateGrid()
	//Act
	_, _, got := takeShot(grid, 1, 6)
	//Assert

	if got != nil {
		t.Error("Cannot shoot at y=6.")
	}
}

func TestCanShootAtX0(t *testing.T) {
	//Arrange
	grid := CreateGrid()
	//Act
	_, _, got := takeShot(grid, 0, 1)
	//Assert

	if got != nil {
		t.Error("Cannot shoot at x=0.")
	}
}

func TestCanShootAtY0(t *testing.T) {
	//Arrange
	grid := CreateGrid()
	//Act
	_, _, got := takeShot(grid, 1, 0)
	//Assert

	if got != nil {
		t.Error("Cannot shoot at y=0.")
	}
}

func TestCannotShootAtXCoordinatePast6(t *testing.T) {
	//Arrange
	grid := CreateGrid()
	//Act
	_, _, got := takeShot(grid, 7, 6)
	//Assert
	want := errors.New("coordinates out of bounds")
	if got.Error() != want.Error() {
		t.Error("Shot was taken outside of the grid.")
	}
}
func TestCannotShootAtYCoordinatePast6(t *testing.T) {
	//Arrange
	grid := CreateGrid()
	//Act
	_, _, got := takeShot(grid, 6, 7)
	//Assert
	want := errors.New("coordinates out of bounds")
	if got.Error() != want.Error() {
		t.Error("Shot was taken outside of the grid.")
	}
}

func TestCannotShootAtNegativeXCoordinate(t *testing.T) {
	//Arrange
	grid := CreateGrid()
	//Act
	_, _, got := takeShot(grid, -1, 5)
	//Assert
	want := errors.New("coordinates out of bounds")
	if got.Error() != want.Error() {
		t.Error("Shot was taken outside of the grid.")
	}
}

func TestCannotShootAtNegativeYCoordinate(t *testing.T) {
	//Arrange
	grid := CreateGrid()
	//Act
	_, _, got := takeShot(grid, 1, -1)
	//Assert
	want := errors.New("coordinates out of bounds")
	if got.Error() != want.Error() {
		t.Error("Shot was taken outside of the grid.")
	}
}
func TestGameOverAt0Ships(t *testing.T) {
	//Arrange
	grid := CreateGrid()
	grid, _ = PlaceShip(grid, 0, 0)
	grid, _ = PlaceShip(grid, 1, 1)
	grid, _ = PlaceShip(grid, 2, 2)
	grid, _ = PlaceShip(grid, 3, 3)
	grid, _ = PlaceShip(grid, 4, 4)
	grid, _ = PlaceShip(grid, 5, 5)
	grid, _ = PlaceShip(grid, 6, 6)
	grid, _ = PlaceShip(grid, 0, 1)
	grid, _ = PlaceShip(grid, 1, 2)
	shipCount := countShips(grid)
	//Act
	grid, _, _ = takeShot(grid, 0, 0)
	grid, _, _ = takeShot(grid, 1, 1)
	grid, _, _ = takeShot(grid, 2, 2)
	grid, _, _ = takeShot(grid, 3, 3)
	grid, _, _ = takeShot(grid, 4, 4)
	grid, _, _ = takeShot(grid, 5, 5)
	grid, _, _ = takeShot(grid, 6, 6)
	grid, _, _ = takeShot(grid, 0, 1)
	grid, _, _ = takeShot(grid, 1, 2)
	//Assert
	got := isGameOver(grid, shipCount)
	want := true
	if got != want {
		t.Error("Game should have ended.")
	}
}
func TestGameNotOverWith1ShipLeft(t *testing.T) {
	//Arrange
	grid := CreateGrid()
	grid, _ = PlaceShip(grid, 0, 0)
	grid, _ = PlaceShip(grid, 1, 1)
	grid, _ = PlaceShip(grid, 2, 2)
	grid, _ = PlaceShip(grid, 3, 3)
	grid, _ = PlaceShip(grid, 4, 4)
	grid, _ = PlaceShip(grid, 5, 5)
	grid, _ = PlaceShip(grid, 6, 6)
	grid, _ = PlaceShip(grid, 0, 1)
	grid, _ = PlaceShip(grid, 1, 2)
	shipCount := countShips(grid)
	//Act
	grid, _, _ = takeShot(grid, 0, 0)
	grid, _, _ = takeShot(grid, 1, 1)
	grid, _, _ = takeShot(grid, 2, 2)
	grid, _, _ = takeShot(grid, 3, 3)
	grid, _, _ = takeShot(grid, 4, 4)
	grid, _, _ = takeShot(grid, 5, 5)
	grid, _, _ = takeShot(grid, 6, 6)
	grid, _, _ = takeShot(grid, 0, 1)
	//Assert
	got := isGameOver(grid, shipCount)
	want := false
	if got != want {
		t.Error("Game should not have ended.")
	}
}

func TestAllShipsBeingShotWithUnder9PlacedEndsGame(t *testing.T) {
	//Arrange
	grid := CreateGrid()
	grid, _ = PlaceShip(grid, 0, 0)
	grid, _ = PlaceShip(grid, 1, 1)
	shipCount := countShips(grid)
	//Act
	grid, _, _ = takeShot(grid, 0, 0)
	grid, _, _ = takeShot(grid, 1, 1)
	//Assert
	got := isGameOver(grid, shipCount)
	want := true
	if got != want {
		t.Error("Game should have ended.")
	}
}
func TestTurnChangesFromPlayer1ToPlayer2(t *testing.T) {
	//Arrange
	currentPlayer := 1
	//Act
	currentPlayer = changeTurnsBetweenPlayers1And2(currentPlayer)
	//Assert
	if currentPlayer != 2 {
		t.Error("Expected current player to be 2, got ", currentPlayer)
	}
}
func TestTurnChangesFromPlayer2ToPlayer1(t *testing.T) {
	//Arrange
	currentPlayer := 2
	//Act
	currentPlayer = changeTurnsBetweenPlayers1And2(currentPlayer)
	//Assert
	if currentPlayer != 1 {
		t.Error("Expected current player to be 1, got ", currentPlayer)
	}
}
func TestTurnChangesIfShotResultIsHit(t *testing.T) {
	//Arrange
	currentPlayer := 1
	grid := CreateGrid()
	grid, _ = PlaceShip(grid, 3, 5)
	//Act
	grid, result, _ := takeShot(grid, 3, 5)
	got := turnChangesDependingOnShotTakingResult(currentPlayer, result)
	//Assert
	want := 2
	if got != want {
		t.Error("Expected current player to change to 2 after a hit, got ", got)
	}
}

func TestTurnChangesIfShotResultIsMiss(t *testing.T) {
	//Arrange
	currentPlayer := 2
	grid := CreateGrid()
	grid, _ = PlaceShip(grid, 3, 5)
	//Act
	grid, result, _ := takeShot(grid, 2, 5)
	got := turnChangesDependingOnShotTakingResult(currentPlayer, result)
	//Assert
	want := 1
	if got != want {
		t.Error("Expected current player to change to 2 after a miss, got ", got)
	}
}
func TestTurnDoesntChangesIfShotResultIsInvalid(t *testing.T) {
	//Arrange
	currentPlayer := 2
	grid := CreateGrid()
	grid, _ = PlaceShip(grid, 3, 5)
	//Act
	grid, result, _ := takeShot(grid, 2, 7)
	got := turnChangesDependingOnShotTakingResult(currentPlayer, result)
	//Assert
	want := 2
	if got != want {
		t.Error("Expected current player to change to 2 after a miss, got ", got)
	}
}

//other tests here that fail

// sometimes we write tests to test our own functions.
func TestGetRandomGridSquare(t *testing.T) {
	gridSquare := getRandomGridSquare()

	//literally only exists here to show you the output
	//should not exist in a real test
	fmt.Println(gridSquare)

	//poor test making use of magic numbers
	//you should probably re-write it
	if gridSquare[0] <= 0 || gridSquare[0] >= 10 {
		t.Error("Grid square row should be >0 and <10, but got: ", gridSquare[0])
	}

	if gridSquare[1] <= 0 || gridSquare[1] >= 10 {
		t.Error("Grid square column should be >0 and <10, but got: ", gridSquare[1])
	}
}
