package main

import (
	"encoding/json"
	"html/template"
	"log"
	"myGame/game"
	"net/http"
)

func homePage(writer http.ResponseWriter, req *http.Request) {
	renderTemplate(writer, "index.html")
}

func playGame(writer http.ResponseWriter, req *http.Request) {
	gameData := game.GetData()
	out, err := json.MarshalIndent(gameData, "", "    ")
	if err != nil {
		log.Println(err)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(out)
}

func renderTemplate(writer http.ResponseWriter, page string) {
	re, err := template.ParseFiles(page)
	if err != nil {
		log.Println("Error")
		return
	}
	err = re.Execute(writer, nil)
	if err != nil {
		log.Println(err)
		return
	}
}

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/play", playGame)
	log.Println("Server is starting on port 8080")
	http.ListenAndServe(":8080", nil)
}
