package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Player struct {
	Name          string `json:"name"`
	MatchesPlayed int    `json:"matches_played"`
	RunScored     int    `json:"runs"`
}

func main() {
	// Read players.json file
	data, err := ioutil.ReadFile("players.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Unmarshal JSON data into a slice of Player structs
	var players []Player
	err = json.Unmarshal(data, &players)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Modify each player's matches played and runs
	for i, player := range players {
		players[i].MatchesPlayed += 7
		players[i].RunScored += 100
		fmt.Println(player)
	}

	// Marshal modified players data into JSON
	modifiedData, err := json.MarshalIndent(players, "", " ")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Write modified data to modified_players.json file
	err = ioutil.WriteFile("modified_players.json", modifiedData, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
}
