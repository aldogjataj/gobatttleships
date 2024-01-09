package main

import (
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

//these are the two tests we have for our functions in main
//the purpose of tests is to mimic interaction with our code
//there is no "user input" - the test is the calling code

// here is an example of a failing test - what do we need to do to fix it?
func TestCreateGrid(t *testing.T) {
	grid := CreateGrid()

	gridSizeCols := len(grid)
	if gridSizeCols != 7 {
		t.Error("Grid has the wrong Number of collums! Expected size of 7, got: ", gridSizeCols)
	}
	gridSizeRows := len(grid[0])
	if gridSizeRows != 7 {
		t.Error("Grid has the wrong number of rows. Expected 7, got:", gridSizeRows)
	}
}

//one good place to start here is by using our utility function
//to target a random grid square rather than 1,1 co-ordinates every time

func TestPlayerOneTakingShot(t *testing.T) {
	grid := CreateGrid()
	shotResult := PlayerOneTurn(grid, []int{1, 1})
	if shotResult != false {
		t.Error("Shot should be false!")
	}
}

func TestPlayerTwoTakingShot(t *testing.T) {
	grid := CreateGrid()
	shotResult := PlayerTwoTurn(grid, []int{1, 1})
	if shotResult != true {
		t.Error("Shot should be true!")
	}
}

func TestPlaceAShip(t *testing.T) {
	//Arrange
	grid := CreateGrid()
	//Act
	desiredCol := 3
	desiredRow := 5
	updatedGrid, _ := PlaceShip(grid, desiredCol, desiredRow)
	//Assert
	actual := updatedGrid[3][5]
	want := "S"
	if actual != want {
		t.Error("Ship was not placed at [3,5]")
	}
}
func TestCanPlaceNineShips(t *testing.T) {
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
	shipCount := 0
	for _, row := range grid {
		for _, cell := range row {
			if cell == "S" {
				shipCount++
			}
		}
	}
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

	_, err := PlaceShip(grid, 3, 4)
	// Assert
	if err == nil {
		t.Errorf("Expected an error for placing a 10th ship, didn't get one.")
	}
	if grid[3][4] == "S" {
		t.Errorf("A tenth ship was placed.")
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
	_, err := PlaceShip(grid, 3, 5)
	//Assert
	if err == nil {
		t.Errorf("Expected an error, didn't get one.")
	}

}
func TestCannotStackShips(t *testing.T) {
	// Arrange
	grid := CreateGrid()
	grid, _ = PlaceShip(grid, 3, 5)

	// Act
	_, err := PlaceShip(grid, 3, 5)

	// Assert
	got := err
	want := "Two ships cannot be placed on the same place"
	if got.Error() != want {
		t.Error("ships are on the same place")
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
func TestCanPlaceShipAty0(t *testing.T) {
	//Arrange
	grid := CreateGrid()
	//Act
	_, err := PlaceShip(grid, 1, 0)
	//Assert
	if err != nil {
		t.Error("Cannot place a ship at y=0.")
	}
}
func TestCanPlaceAtx6(t *testing.T) {
	//Arrange
	grid := CreateGrid()
	//Act
	_, err := PlaceShip(grid, 6, 1)
	//Assert
	if err != nil {
		t.Error("Cannot place a ship at x=6.")
	}
}
func TestCanPlaceAty6(t *testing.T) {
	//Arrange
	grid := CreateGrid()
	//Act
	_, err := PlaceShip(grid, 1, 6)
	//Assert
	if err != nil {
		t.Error("Cannot place a ship at y=6.")
	}
}

func TestCannotPlaceatX7(t *testing.T) {
	//Arrange
	grid := CreateGrid()
	//Act
	_, err := PlaceShip(grid, 7, 2)
	//Assert
	if err == nil {
		t.Error("Error did not return from attempting to place a ship outside the grid. ")
	}
}

func TestCannotPlaceatY7(t *testing.T) {
	//Arrange
	grid := CreateGrid()
	//Act
	_, err := PlaceShip(grid, 2, 7)
	//Assert
	if err == nil {
		t.Error("Error did not return from attempting to place a ship outside the grid. ")
	}
}
func TestCannotPlaceShipAtNegativeX(t *testing.T) {
	//Act
	grid := CreateGrid()
	//Act
	_, err := PlaceShip(grid, -1, 5)
	//Assert
	if err == nil {
		t.Error("Error did not return from attempting to place a ship outside the grid. ")
	}
}
func TestCannotPlaceShipAtNegativeY(t *testing.T) {
	//Act
	grid := CreateGrid()
	//Act
	_, err := PlaceShip(grid, 1, -1)
	//Assert
	if err == nil {
		t.Error("Error did not return from attempting to place a ship outside the grid. ")
	}
}
func TestShotIsRecorded(t *testing.T) {
	// Arrange
	grid := CreateGrid()

	// Act
	_, result, err := takeShot(grid, 3, 5)

	// Assert
	if err != nil {
		t.Error("Unexpected error:", err)
	}
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
	_, result, err := takeShot(grid, 3, 5)

	// Assert
	if err != nil {
		t.Error("Unexpected error:", err)
	}
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
	_, result, err := takeShot(grid, 3, 5)

	// Assert
	if err != nil {
		t.Error("Unexpected error:", err)
	}
	want := "MISS"
	if result != want {
		t.Error("Shooting at a sunk ship still shows as a hit.")
	}
}
func TestCanShootAtX6(t *testing.T) {
	//Arrange
	grid := CreateGrid()
	//Act
	_, result, err := takeShot(grid, 6, 1)
	//Assert
	if err != nil {
		t.Error("Unexpected error:", err)
	}
	want := "MISS"
	if result != want {
		t.Error("Cannot shoot at x=6.")
	}
}

func TestCanShootAtY6(t *testing.T) {
	//Arrange
	grid := CreateGrid()
	//Act
	_, result, err := takeShot(grid, 1, 6)
	//Assert
	if err != nil {
		t.Error("Unexpected error:", err)
	}
	want := "MISS"
	if result != want {
		t.Error("Cannot shoot at y=6.")
	}
}
func TestCanShootAtX0(t *testing.T) {
	//Arrange
	grid := CreateGrid()
	//Act
	_, result, err := takeShot(grid, 0, 1)
	//Assert
	if err != nil {
		t.Error("Unexpected error:", err)
	}
	want := "MISS"
	if result != want {
		t.Error("Cannot shoot at x=0.")
	}
}

func TestCanShootAtY0(t *testing.T) {
	//Arrange
	grid := CreateGrid()
	//Act
	_, result, err := takeShot(grid, 1, 0)
	//Assert
	if err != nil {
		t.Error("Unexpected error:", err)
	}
	want := "MISS"
	if result != want {
		t.Error("Cannot shoot at y=0.")
	}
}
func TestCannotShootAtXCoordinatePast6(t *testing.T) {
	//Arrange
	grid := CreateGrid()
	//Act
	_, result, err := takeShot(grid, 7, 6)
	//Assert
	if err == nil || result == "MISS" {
		t.Error("Shot was taken outside of the grid.")
	}
}
func TestCannotShootAtYCoordinatePast6(t *testing.T) {
	//Arrange
	grid := CreateGrid()
	//Act
	_, result, err := takeShot(grid, 6, 7)
	//Assert
	if err == nil || result == "MISS" {
		t.Error("Shot was taken outside of the grid.")
	}
}

func TestCannotShootAtNegativeXCoordinate(t *testing.T) {
	//Arrange
	grid := CreateGrid()
	//Act
	_, result, err := takeShot(grid, -1, 5)
	//Assert
	if err == nil || result == "MISS" {
		t.Error("Shot was taken outside of the grid.")
	}
}

func TestCannotShootAtNegativeYCoordinate(t *testing.T) {
	//Arrange
	grid := CreateGrid()
	//Act
	_, result, err := takeShot(grid, 1, -1)
	//Assert
	if err == nil || result == "MISS" {
		t.Error("Shot was taken outside of the grid.")
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
