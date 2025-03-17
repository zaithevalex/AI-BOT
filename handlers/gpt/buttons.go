package gpt

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

const (
	errRepeatChoice = "bad request, Bad Request: message is not modified: specified new message content and reply markup are exactly the same as a current content and reply markup of the message"
	paymentToken    = "381764678:TEST:99363"
	subDesc         = "GPT3.5/GPT-4/GoogleAI\n✅ 100 запросов ежедневно\n✅ не нужно ждать, чтобы задать следующий вопрос\nСтоимость: \n💰 99 руб     - 2 недели\n💰 169 руб   - 1 месяц (экономия: 32%)\n💰 1599 руб - 1 год (экономия: 61%)"
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
	
		ais := []*Message{
			{Name: "GPT-3.5", ButtonTag: "button_pick_gpt35"},
			{Name: "GPT-4", ButtonTag: "button_pick_gpt4"},
			{Name: "🔙Назад", ButtonTag: "button_pick_gpt_back"},
		}
	
		chosenParam, err := manage.GetParam[string](db, manage.GetUserParam, update.CallbackQuery.Message.Message.Chat.ID, "ai")
		if err != nil {
			panic(err.Error())
		}
		AddCheckMark(ais, chosenParam)
	
		_, err = b.EditMessageReplyMarkup(ctx, &bot.EditMessageReplyMarkupParams{
			ChatID:      update.CallbackQuery.Message.Message.Chat.ID,
			MessageID:   update.CallbackQuery.Message.Message.ID,
			ReplyMarkup: InlineKeyboardMarkUpGenerate(ais),
		})
		if err != nil {
			if err.Error() == errRepeatChoice {
				break
			}
			panic(err.Error())
		}
		break
	case "button_pick_gpt35":
		db, err := manage.Connection()
		if err != nil {
			panic(err.Error())
		}
	
		err = manage.UpdateParam(db, manage.UpdateUserParam, update.CallbackQuery.Message.Message.Chat.ID, "ai", manage.GPT35)
		if err != nil {
			panic(err.Error())
		}
	
		ais := []*Message{
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
			if err.Error() == errRepeatChoice {
				break
			}
			panic(err.Error())
		}
		break
	case "button_pick_gpt4":
		db, err := manage.Connection()
		if err != nil {
			panic(err.Error())
		}
	
		err = manage.UpdateParam(db, manage.UpdateUserParam, update.CallbackQuery.Message.Message.Chat.ID, "ai", manage.GPT4)
		if err != nil {
			panic(err.Error())
		}
	
		ais := []*Message{
			{Name: "GPT-3.5", ButtonTag: "button_pick_gpt35"},
			{Name: "GPT-4✅", ButtonTag: "button_pick_gpt4"},
			{Name: "🔙Назад", ButtonTag: "button_pick_gpt_back"},
		}
	
		_, err = b.EditMessageReplyMarkup(ctx, &bot.EditMessageReplyMarkupParams{
			ChatID:      update.CallbackQuery.Message.Message.Chat.ID,
			MessageID:   update.CallbackQuery.Message.Message.ID,
			ReplyMarkup: InlineKeyboardMarkUpGenerate(ais),
		})
		if err != nil {
			if err.Error() == errRepeatChoice {
				break
			}
			panic(err.Error())
		}
		break
	case "button_pick_gpt_back":
		db, err := manage.Connection()
		if err != nil {
			panic(err.Error())
		}
	
		chosenParam, err := manage.GetParam[string](db, manage.GetUserParam, update.CallbackQuery.Message.Message.Chat.ID, "ai")
		if err != nil {
			panic(err.Error())
		}
	
		ais := []*Message{
			{Name: "GPT", ButtonTag: "button_pick_gpt"},
			{Name: "GoogleAI", ButtonTag: "button_pick_googleai"},
		}
		AddCheckMark(ais, chosenParam)
	
		_, err = b.EditMessageReplyMarkup(ctx, &bot.EditMessageReplyMarkupParams{
			ChatID:      update.CallbackQuery.Message.Message.Chat.ID,
			MessageID:   update.CallbackQuery.Message.Message.ID,
			ReplyMarkup: InlineKeyboardMarkUpGenerate(ais),
		})
		if err != nil {
			if err.Error() == errRepeatChoice {
				break
			}
			panic(err.Error())
		}
		break
	case "button_pick_googleai":
		db, err := manage.Connection()
		if err != nil {
			panic(err.Error())
		}
	
		err = manage.UpdateParam(db, manage.UpdateUserParam, update.CallbackQuery.Message.Message.Chat.ID, "ai", manage.GoogleAI)
		if err != nil {
			panic(err.Error())
		}
	
		ais := []*Message{
			{Name: "GPT", ButtonTag: "button_pick_gpt"},
			{Name: "GoogleAI✅", ButtonTag: "button_pick_googleai"},
		}
	
		_, err = b.EditMessageReplyMarkup(ctx, &bot.EditMessageReplyMarkupParams{
			ChatID:      update.CallbackQuery.Message.Message.Chat.ID,
			MessageID:   update.CallbackQuery.Message.Message.ID,
			ReplyMarkup: InlineKeyboardMarkUpGenerate(ais),
		})
		if err != nil {
			if err.Error() == errRepeatChoice {
				break
			}
			panic(err.Error())
		}
		break
	case "button_get_subscribe":
		_, err = b.EditMessageText(ctx, &bot.EditMessageTextParams{
			ChatID:    update.CallbackQuery.Message.Message.Chat.ID,
			MessageID: update.CallbackQuery.Message.Message.ID,
			Text:      subDesc,
		})
	
		aisPrices := []*Message{
			{Name: "2 недели", ButtonTag: "button_get_subscribe_2weeks"},
			{Name: "1 месяц", ButtonTag: "button_get_subscribe_1month"},
			{Name: "1 год", ButtonTag: "button_get_subscribe_1year"},
		}
	
		_, err = b.EditMessageReplyMarkup(ctx, &bot.EditMessageReplyMarkupParams{
			ChatID:      update.CallbackQuery.Message.Message.Chat.ID,
			MessageID:   update.CallbackQuery.Message.Message.ID,
			ReplyMarkup: InlineKeyboardMarkUpGenerate(aisPrices),
		})
		if err != nil {
			if err.Error() == errRepeatChoice {
				break
			}
			panic(err.Error())
		}
		break
	case "button_get_subscribe_2weeks":
		_, err = b.SendInvoice(ctx, &bot.SendInvoiceParams{
			ChatID:        update.CallbackQuery.Message.Message.Chat.ID,
			Title:         "Подписка на 2 недели",
			Description:   "Вы сможете отправлять неограниченное количество запросов, используя все виды AI в течение двух недель.",
			Payload:       manage.Payload2Weeks,
			ProviderToken: paymentToken,
			Currency:      "RUB",
			Prices: []models.LabeledPrice{
				{Label: "Подписка на 2 недели", Amount: 9900},
			},
		})
		if err != nil {
			panic(err.Error())
		}
		break
	case "button_get_subscribe_1month":
		_, err = b.SendInvoice(ctx, &bot.SendInvoiceParams{
			ChatID:        update.CallbackQuery.Message.Message.Chat.ID,
			Title:         "Подписка на 2 недели",
			Description:   "Вы сможете отправлять неограниченное количество запросов, используя все виды AI в течение одного месяца.",
			Payload:       manage.Payload1Month,
			ProviderToken: paymentToken,
			Currency:      "RUB",
			Prices: []models.LabeledPrice{
				{Label: "Подписка на 1 месяц", Amount: 16900},
			},
		})
		if err != nil {
			panic(err.Error())
		}
		break
	case "button_get_subscribe_1year":
		_, err = b.SendInvoice(ctx, &bot.SendInvoiceParams{
			ChatID:        update.CallbackQuery.Message.Message.Chat.ID,
			Title:         "Подписка на 1 год",
			Description:   "Вы сможете отправлять неограниченное количество запросов, используя все виды AI в течение одного года.",
			Payload:       manage.Payload1Year,
			ProviderToken: paymentToken,
			Currency:      "RUB",
			Prices: []models.LabeledPrice{
				{Label: "Подписка на 1 год", Amount: 159900},
			},
		})
		if err != nil {
			panic(err.Error())
		}
		break
	default:
		break
	}
}
