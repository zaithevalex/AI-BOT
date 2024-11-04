package gpt

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"strings"
)

type Button struct {
	Name      string
	ButtonTag string
}

func AddCheckMark(buttons []*Button, name string) int {
	name = strings.Trim(name, " ")
	for index, button := range buttons {
		if button.Name == name {
			button.Name += "✅"
			return index
		}
	}
	return -1
}

func InlineKeyboardMarkUpGenerate(buttons []*Button) models.InlineKeyboardMarkup {
	var replyMarkup [][]models.InlineKeyboardButton
	var rowKeys []models.InlineKeyboardButton
	for index, button := range buttons {
		if index%2 == 0 && len(rowKeys) > 0 {
			replyMarkup = append(replyMarkup, rowKeys)
			rowKeys = nil
		}
		rowKeys = append(rowKeys, models.InlineKeyboardButton{Text: button.Name, CallbackData: button.ButtonTag})
	}
	replyMarkup = append(replyMarkup, rowKeys)

	return models.InlineKeyboardMarkup{
		InlineKeyboard: replyMarkup,
	}
}

func deleteLastMessageUser(ctx context.Context, b *bot.Bot, update *models.Update) error {
	_, err := b.DeleteMessage(ctx, &bot.DeleteMessageParams{
		ChatID:    update.Message.Chat.ID,
		MessageID: update.Message.ID,
	})
	if err != nil {
		return err
	}
	return nil
}
