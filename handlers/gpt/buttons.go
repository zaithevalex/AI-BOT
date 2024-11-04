package gpt

import (
	"BOOT-BOT/db/manage"
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"reflect"
)

func GeneralButtonHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	_, err := b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: update.CallbackQuery.ID,
		ShowAlert:       false,
	})
	if err != nil {
		panic(err.Error())
	}

	switch update.CallbackQuery.Data {
	case "button_pick_gpt":
		db, err := manage.Connection()
		if err != nil {
			panic(err.Error())
		}

		chosenParam, err := manage.GetParam(db, update.CallbackQuery.Message.Message.Chat.ID, "ai")
		if err != nil {
			panic(err.Error())
		}

		ais := []*Button{
			{Name: "GPT-3.5", ButtonTag: "button_pick_gpt35"},
			{Name: "GPT-4", ButtonTag: "button_pick_gpt4"},
		}

		AddCheckMark(ais, reflect.ValueOf(chosenParam).String())

		_, err = b.EditMessageReplyMarkup(ctx, &bot.EditMessageReplyMarkupParams{
			ChatID:      update.CallbackQuery.Message.Message.Chat.ID,
			MessageID:   update.CallbackQuery.Message.Message.ID,
			ReplyMarkup: InlineKeyboardMarkUpGenerate(ais),
		})
		if err != nil {
			panic(err.Error())
		}
		break
	case "button_pick_gpt35":
		db, err := manage.Connection()
		if err != nil {
			panic(err.Error())
		}

		err = manage.UpdateParam(db, update.CallbackQuery.Message.Message.Chat.ID, "ai", 1)
		if err != nil {
			panic(err.Error())
		}

		ais := []*Button{
			{Name: "GPT-3.5✅", ButtonTag: "button_pick_gpt35"},
			{Name: "GPT-4", ButtonTag: "button_pick_gpt4"},
			{Name: "🔙Назад", ButtonTag: "button_pick_gpt_back"},
		}

		_, err = b.EditMessageReplyMarkup(ctx, &bot.EditMessageReplyMarkupParams{
			ChatID:      update.CallbackQuery.Message.Message.Chat.ID,
			MessageID:   update.CallbackQuery.Message.Message.ID,
			ReplyMarkup: InlineKeyboardMarkUpGenerate(ais),
		})
		if err != nil {
			panic(err.Error())
		}
		break
	case "button_pick_gpt4":
		db, err := manage.Connection()
		if err != nil {
			panic(err.Error())
		}

		err = manage.UpdateParam(db, update.CallbackQuery.Message.Message.Chat.ID, "ai", 2)
		if err != nil {
			panic(err.Error())
		}

		ais := []*Button{
			{Name: "GPT-3.5", ButtonTag: "button_pick_gpt35"},
			{Name: "GPT-4", ButtonTag: "button_pick_gpt4"},
			{Name: "🔙Назад", ButtonTag: "button_pick_gpt_back"},
		}

		_, err = b.EditMessageReplyMarkup(ctx, &bot.EditMessageReplyMarkupParams{
			ChatID:      update.CallbackQuery.Message.Message.Chat.ID,
			MessageID:   update.CallbackQuery.Message.Message.ID,
			ReplyMarkup: InlineKeyboardMarkUpGenerate(ais),
		})
		if err != nil {
			panic(err.Error())
		}
		break
	case "button_pick_gpt_back":
		db, err := manage.Connection()
		if err != nil {
			panic(err.Error())
		}

		chosenParam, err := manage.GetParam(db, update.CallbackQuery.Message.Message.Chat.ID, "ai")
		if err != nil {
			panic(err.Error())
		}

		ais := []*Button{
			{Name: "GPT", ButtonTag: "button_pick_gpt"},
			{Name: "GoogleAI✅", ButtonTag: "button_pick_gpt_back"},
		}

		SetButtonSelectionStatus(ais, reflect.ValueOf(chosenParam).String())

		_, err = b.EditMessageReplyMarkup(ctx, &bot.EditMessageReplyMarkupParams{
			ChatID:      update.CallbackQuery.Message.Message.Chat.ID,
			MessageID:   update.CallbackQuery.Message.Message.ID,
			ReplyMarkup: InlineKeyboardMarkUpGenerate(ais),
		})
		if err != nil {
			panic(err.Error())
		}
		break
	case "button_pick_googleai":
		db, err := manage.Connection()
		if err != nil {
			panic(err.Error())
		}

		err = manage.UpdateParam(db, update.CallbackQuery.Message.Message.Chat.ID, "ai", 0)
		if err != nil {
			panic(err.Error())
		}

		ais := []*Button{
			{Name: "GPT", ButtonTag: "button_pick_gpt"},
			{Name: "GoogleAI✅", ButtonTag: "button_pick_googleai"},
		}

		_, err = b.EditMessageReplyMarkup(ctx, &bot.EditMessageReplyMarkupParams{
			ChatID:      update.CallbackQuery.Message.Message.Chat.ID,
			MessageID:   update.CallbackQuery.Message.Message.ID,
			ReplyMarkup: InlineKeyboardMarkUpGenerate(ais),
		})
		if err != nil {
			panic(err.Error())
		}
		break
	case "button_get_subscribe":
		_, err = b.EditMessageText(ctx, &bot.EditMessageTextParams{
			ChatID:    update.CallbackQuery.Message.Message.Chat.ID,
			MessageID: update.CallbackQuery.Message.Message.ID,
			Text:      "", // TODO
		})

		aisPrices := []*Button{
			{Name: "GPT-3.5", ButtonTag: "button_get_subscribe_gpt4"},
			{Name: "GPT-4", ButtonTag: "button_get_subscribe_gpt35"},
			{Name: "GoogleAI", ButtonTag: "button_get_subscribe_googleai"},
		}

		_, err = b.EditMessageReplyMarkup(ctx, &bot.EditMessageReplyMarkupParams{
			ChatID:      update.CallbackQuery.Message.Message.Chat.ID,
			MessageID:   update.CallbackQuery.Message.Message.ID,
			ReplyMarkup: InlineKeyboardMarkUpGenerate(aisPrices),
		})
	default:
		break
	}
}
