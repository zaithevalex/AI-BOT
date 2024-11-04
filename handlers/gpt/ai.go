package gpt

import (
	"BOOT-BOT/db/manage"
	"BOOT-BOT/db/timers"
	"BOOT-BOT/handlers/general"
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
	"reflect"
)

const (
	googleAIApiKey   = "AIzaSyDgG7MWSRHQr84zlgWyItUuIISrF5HGNAo"
	googleAIVersion  = "gemini-1.5-flash"
	gptDesc          = "GPT-4 - это не просто число, это символ инноваций, оптимизации и бесконечных возможностей. 🚀\n\nGPT-4 представляет собой:\n\n• Переосмысление ограничений: GPT-4 смещает привычные представления, открывая новые горизонты для решения сложных задач. 🤯\n• Слияние традиций и будущего: GPT-4 комбинирует проверенные временем методы с передовыми технологиями. 🕰️🤖\n• Бесконечный потенциал: GPT-4 - это не просто конечная точка, а отправная точка для новых открытий. 💫\n\nGPT-4 - это ваш ключ к успеху в стремительно меняющемся мире. 🔑🌎\n\n\nGoogle GenAI:\n\nGoogle GenAI - ваш личный помощник в мире искусственного интеллекта. 🤖🧠\n\nGoogle GenAI - это:\n\n• Мощь Google в вашем распоряжении: Доступ к непостижимым вычислительным ресурсам Google, позволяющим решать сложные задачи с невероятной скоростью и точностью. ⚡\n• Инновационные алгоритмы: Создавайте контент, генерируйте идеи, анализируйте данные и решайте задачи с помощью передовых алгоритмов машинного обучения. 💡\n• Интуитивный интерфейс: Google GenAI доступен всем: от начинающих до опытных пользователей. 🧑‍💻\n\nGoogle GenAI - ваш путь к беспрецедентным возможностям в мире искусственного интеллекта. 🚀"
	infSubscribeTime = 100
	nilRequestsDesc  = "Вы достигли недельного лимита, предоставляемого бесплатной версией. Чтобы иметь возможность отправлять больше запросов, приобретите подписку. Спасибо за ваш интерес ❤️"
)

func GPTHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.PreCheckoutQuery != nil {
		fmt.Printf("get PreCheckoutQuery for invoce payload: %s\n", update.PreCheckoutQuery.InvoicePayload)

		b.AnswerPreCheckoutQuery(ctx, &bot.AnswerPreCheckoutQueryParams{
			PreCheckoutQueryID: update.PreCheckoutQuery.ID,
			OK:                 true,
			ErrorMessage:       "error: payment stage not finished.",
		})
		return
	}

	db, err := general.Auth(update.Message.Chat.ID)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	if update.Message != nil {
		if update.Message.SuccessfulPayment != nil {
			b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID:    update.Message.Chat.ID,
				Text:      "Payment successful.",
				ParseMode: models.ParseModeMarkdown,
			})
			return
		}
	}

	val, err := manage.GetParam(db, update.Message.Chat.ID, "subscribe_time")
	if err != nil {
		panic(err.Error())
	}
	if reflect.ValueOf(val).Int() < infSubscribeTime {
		err = manage.UpdateParam(db, update.Message.Chat.ID, "subscribe_time", timers.StartWeekUpdate())
		if err != nil {
			panic(err.Error())
		}

		err = manage.UpdateParam(db, update.Message.Chat.ID, "amount_requests", manage.DefaultReqsPerWeek)
		if err != nil {
			panic(err.Error())
		}
	}

	val, err = manage.GetParam(db, update.Message.Chat.ID, "amount_requests")
	amountReqs := reflect.ValueOf(val).Int()
	if amountReqs < 1 {
		_, err = b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   nilRequestsDesc,
			ReplyMarkup: models.InlineKeyboardMarkup{
				InlineKeyboard: [][]models.InlineKeyboardButton{
					{
						{Text: "Приобрести подписку", CallbackData: "button_get_subscribe"},
					},
				},
			},
		})
		if err != nil {
			panic(err.Error())
		}

		return
	}

	err = manage.UpdateParam(db, update.Message.Chat.ID, "amount_requests", amountReqs-1)
	if err != nil {
		panic(err.Error())
	}

	resp, err := googleAISend(ctx, update.Message.Text)
	if err != nil {
		panic(err.Error())
	}

	_, err = b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   resp,
	})
	if err != nil {
		panic(err.Error())
	}
}

func googleAISend(ctx context.Context, req string) (string, error) {
	client, err := genai.NewClient(ctx, option.WithAPIKey(googleAIApiKey))
	if err != nil {
		return "", err
	}
	defer client.Close()

	model := client.GenerativeModel(googleAIVersion)
	resp, err := model.GenerateContent(ctx, genai.Text(req))
	if err != nil {
		return "", err
	}

	return fmt.Sprint(resp.Candidates[0].Content.Parts[0]), nil
}
