package gpt

import (
	"BOOT-BOT/db/manage"
	"BOOT-BOT/handlers/general"
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func PickGPTHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	db, err := general.Auth(update.Message.Chat.ID)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	err = deleteLastMessageUser(ctx, b, update)
	if err != nil {
		panic(err.Error())
	}

	ais := []*Message{
		{Name: "GPT", ButtonTag: "button_pick_gpt"},
		{Name: "GoogleAI", ButtonTag: "button_pick_googleai"},
	}

	chosenParam, err := manage.GetParam[string](db, manage.GetUserParam, update.Message.Chat.ID, "ai")
	if err != nil {
		panic(err.Error())
	}
	AddCheckMark(ais, chosenParam)

	_, err = b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        gptDesc,
		ReplyMarkup: InlineKeyboardMarkUpGenerate(ais),
	})
	if err != nil {
		panic(err.Error())
	}
}
