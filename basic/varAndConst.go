package basic

import "fmt"

func VarAndConstExample() {
	const hotelName string = "Gopher Hotel"
	const longitude = 24.806078
	const latitude = -78.243027

	var occupancy int = 12
	fmt.Printf("%v\n %v\n %v\n %v\n", hotelName, longitude, latitude, occupancy)
}