package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/line/line-bot-sdk-go/v7/linebot/httphandler"
)

func main() {
	handler, err := httphandler.New(
		"af42dc438bbcf308b5b0d274b4e1846e",
		"CX7kjGwq6ASjy3wd2SRihDD4XhlEzVKbTQ07JIUqGhNhXHuQwJ1L9NdP80uvSpqFz7qpmsdSQO0r9HmvEITCUGoy4j/zJWxwx09+5P8Mklzbo1H2FBnrrPXYx3iFhl+iZU74LMu0q8HEpQCj/vk1DgdB04t89/1O/w1cDnyilFU=",
	)
	if err != nil {
		log.Fatal(err)
	}

	// Setup HTTP Server for receiving requests from LINE platform
	handler.HandleEvents(func(events []*linebot.Event, r *http.Request) {
		bot, err := handler.NewClient()
		if err != nil {
			log.Print(err)
			return
		}
		for _, event := range events {
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					msg := fmt.Sprintf("reMsg: %v \ngroupID: %v\n", message.Text, event.Source.GroupID)
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(msg)).Do(); err != nil {
						log.Print(err)
					}
				}
			}
		}
	})
	http.Handle("/callback", handler)
	// This is just a sample code.
	// For actually use, you must support HTTPS by using `ListenAndServeTLS`, reverse proxy or etc.
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		log.Fatal(err)
	}
}
