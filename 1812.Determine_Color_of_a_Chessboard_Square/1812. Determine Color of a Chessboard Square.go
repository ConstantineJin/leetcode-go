package main

func squareIsWhite(coordinates string) bool {
	return coordinates[0]%2 != coordinates[1]%2
}
