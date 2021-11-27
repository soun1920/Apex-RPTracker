package main

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)


func get_Request(playerId string) **Apextrackergg {
	godotenv.Load()

	token := os.Getenv("TRACKER_TOKEN")
	URL := "https://public-api.tracker.gg/v2/apex/standard/profile/origin/"
	param := "?TRN-Api-Key=" + token
	var profile *Apextrackergg

	client := &http.Client{}
	req, err := http.NewRequest("GET", URL+playerId+param, nil)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(body, &profile); err != nil {
		log.Fatal(err)
	}
	return &profile
}

func main() {
	res := get_Request("soun1920")
	fmt.Printf("%v", (*res).Data.Segments[0].Stats.RankScore.Value)
}
