package basic

import (
	"fmt"
	"math/rand"
	"time"
)

func FunctionsExample() {
	const hotelName = "Gopher Paris Inn"
	const totalRooms = 134
	const initialRoom = 110

	rand.Seed(time.Now().UTC().UnixNano())

	occupiedRooms := rand.Intn(totalRooms)
	availableRooms := totalRooms - occupiedRooms 
	occupancyRate := occupancyRate(totalRooms, occupiedRooms)

	fmt.Println("Hotel: ", hotelName)
	fmt.Println("                            Occupancy level: ", occupancyLevel(occupancyRate))
	fmt.Printf("                            Occupancy rate: %0.2f %%\n", occupancyRate)
	fmt.Println("Number of rooms: ", totalRooms)
	fmt.Println("Rooms available: ", availableRooms)
	fmt.Printf("\n Rooms: \n")

	if availableRooms == 0 {
		fmt.Println("No rooms available for tonight")
	}

	for i := 0; i < availableRooms; i++ {
		printRoomDetails(initialRoom + i)
	}
}

func occupancyLevel(occupancyRate float64) (occupancyLevel string) {
	switch {
	case occupancyRate < 30:
		occupancyLevel = "Low"
	case occupancyRate < 60:
		occupancyLevel = "Medium"
	case occupancyRate >= 60:
		occupancyLevel = "High"
	}
	return occupancyLevel
}

func occupancyRate(totalRooms int, occupiedRooms int) (occupancyRate float64) {
	occupancyRate = float64(occupiedRooms) / float64(totalRooms) * 100
	return
}

func printRoomDetails(roomNumber int) {
	people := rand.Intn(10 - 1) + 1
	nights := rand.Intn(10 - 1) + 1
	nightText := "nights"
	if nights == 1 {
		nightText = "night"
	}
	fmt.Printf("- %v: %v people / %v %v\n", roomNumber, people, nights, nightText)
}

