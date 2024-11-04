package gpt

import (
	"BOOT-BOT/db/manage"
	"BOOT-BOT/handlers/general"
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"reflect"
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

	chosenParam, err := manage.GetParam(db, update.Message.Chat.ID, "ai")
	if err != nil {
		panic(err.Error())
	}

	ais := []*Button{
		{Name: "GoogleAI", ButtonTag: "button_pick_googleai"},
		{Name: "GPT", ButtonTag: "button_pick_gpt"},
	}

	SetButtonSelectionStatus(ais, reflect.ValueOf(chosenParam).String())

	_, err = b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        gptDesc,
		ReplyMarkup: InlineKeyboardMarkUpGenerate(ais),
	})
	if err != nil {
		panic(err.Error())
	}
}
