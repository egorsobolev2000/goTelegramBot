package goTelegramBot

import (
	"bytes"
	"encoding/json"
	gtb "github.com/egorsobolev2000/goTelegramBot"
	"net/http"
)

func SendMessage(update gtb.Update, replayText string, botUrl string) error {
	var botMessage gtb.BotMessage
	botMessage.ChatId = update.Message.Chat.ChatId
	botMessage.Text = replayText
	buf, err := json.Marshal(botMessage)
	if err != nil {
		return err
	}

	_, err = http.Post(botUrl + "/sendMessage", "application/json", bytes.NewBuffer(buf))
	if err != nil {
		return err
	}

	return nil
}