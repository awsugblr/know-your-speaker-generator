package main

import (
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/adlio/trello"
)

type SpeakerInfo struct {
	SpeakerName string
	Description string
}

func main() {

	api_key, token := validate_and_fetch_keys_tokens()

	client := trello.NewClient(api_key, token)

	list, err := client.GetList(Know_Your_Speaker_List_ID, trello.Defaults())
	if err != nil {
		log.Fatal(err.Error())
	}

	cards, err := list.GetCards(trello.Defaults())
	if err != nil {
		log.Fatal(err.Error())
	}
	speaker_info := SpeakerInfo{}
	t := template.New("template")
	t, _ = template.ParseFiles("template.md")

	for _, card := range cards {
		speakerslug := "posts/2019-07-15-" + strings.Replace(strings.ToLower(card.Name), " ", "-", -1) + ".md"
		f, err := os.Create(speakerslug)
		if err != nil {
			log.Fatal(err.Error())
		}
		speaker_info.SpeakerName = card.Name
		speaker_info.Description = card.Desc
		t.Execute(f, speaker_info)
		f.Close()

	}

}
