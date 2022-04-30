package basic

import (
	"fmt"
	"math/rand"
	"time"
)

func ControlStatementExample() {
	const hotelName = "Gopher Paris Inn"
	const totalRooms = 134
	const initialRoom = 110

	rand.Seed(time.Now().UTC().UnixNano())

	occupiedRooms := rand.Intn(totalRooms)
	availableRooms := totalRooms - occupiedRooms 
	occupancyRate := float64(occupiedRooms) / float64(totalRooms) * 100

	var occupancyLevel string
	switch {
		case occupancyRate < 30:
			occupancyLevel = "Low"
		case occupancyRate < 60:
			occupancyLevel = "Medium"
		case occupancyRate >= 60:
			occupancyLevel = "High"
	}

	fmt.Println("Hotel: ", hotelName)
	fmt.Println("                            Occupancy level: ", occupancyLevel)
	fmt.Printf("                            Occupancy rate: %0.2f %%\n", occupancyRate)
	fmt.Println("Number of rooms: ", totalRooms)
	fmt.Println("Rooms available: ", availableRooms)
	fmt.Printf("\n Rooms: \n")

	if availableRooms == 0 {
		fmt.Println("No rooms available for tonight")
	}

	for i := 0; i < availableRooms; i++ {
		people := rand.Intn(10 - 1) + 1
		nights := rand.Intn(10 - 1) + 1
		fmt.Printf("- %v: %v people / %v nights\n", initialRoom + i, people, nights)
	}
}