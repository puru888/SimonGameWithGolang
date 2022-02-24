package game

import (
	"math/rand"
	"time"
)

type manager struct {
	ButtonColors []string `json:"btn_clr"`
	RandomNumber int      `json:"rnd"`
}

func GetData() manager {
	var gameData manager
	rand.Seed(time.Now().UnixNano())
	rndNum := rand.Intn(4)
	clrArr := []string{"red", "blue", "green", "yellow"}
	gameData.ButtonColors = clrArr
	gameData.RandomNumber = rndNum
	return gameData
}
