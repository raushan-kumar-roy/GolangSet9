package main

import (
	"encoding/csv"
	"os"
)

func main() {
	file, err := os.Create("friends.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write([]string{"Name", "Mobile Number", "Alternate Mobile Number", "Address", "City/Town"})
	if err != nil {
		panic(err)
	}

	friends := [][]string{
		{"Alice", "123-456-7890", "", "123 Main Street", "New York"},
		{"Bob", "234-567-8901", "345-678-9012", "456 Market Street", "Chicago"},
		{"Charlie", "345-678-9012", "", "789 Park Avenue", "Los Angeles"},
		{"David", "456-789-0123", "567-890-1234", "321 Maple Street", "Boston"},
		{"Eve", "567-890-1234", "", "654 Pine Street", "San Francisco"},
		{"Frank", "678-901-2345", "789-012-3456", "159 Oak Street", "Houston"},
		{"Greta", "789-012-3456", "", "753 Cedar Street", "Philadelphia"},
		{"Hannah", "890-123-4567", "901-234-5678", "246 Birch Street", "Seattle"},
		{"Igor", "901-234-5678", "", "369 Willow Street", "Miami"},
		{"Jenny", "012-345-6789", "123-456-7890", "159 Maple Street", "Dallas"},
	}
	for _, friend := range friends {
		err = writer.Write(friend)
		if err != nil {
			panic(err)
		}
	}
}
