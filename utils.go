package main

import (
	"log"
	"os"
)

const Know_Your_Speaker_List_ID = "5d1302381c555023f6581e4e"

func validate_and_fetch_keys_tokens() (string, string) {
	api_key := os.Getenv("TRELLO_API_KEY")
	if api_key == "" {
		log.Fatal("TRELLO_API_KEY not set, exiting..")
	}

	token := os.Getenv("TRELLO_API_TOKEN")
	if token == "" {
		log.Fatal("TRELLO_API_TOKEN not set, exiting..")
	}
	return api_key, token
}
