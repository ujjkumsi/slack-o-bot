package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Data struct {
	Category  string `json:"type"`
	Token     string `json:"token"`
	Challenge string `json:"challenge"`
}

type Challenge struct {
	Challenge string `json:"challenge"`
}

func challengeHandler(w http.ResponseWriter, r *http.Request) {
	challenge := &Data{}

	err := json.NewDecoder(r.Body).Decode(challenge)
	if err != nil {
		log.Printf("[ERROR] %s", err)
		panic(err)
	}

	log.Printf("[CHALLENGE] %s", challenge.Challenge)
	resp := &Challenge{}
	resp.Challenge = challenge.Challenge
	challengeJson, err := json.Marshal(resp)
	if err != nil {
		panic(err)
	}

	//set content-type to json
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(challengeJson)
}
