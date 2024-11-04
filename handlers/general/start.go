package general

import (
	"BOOT-BOT/db/manage"
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/jmoiron/sqlx"
)

const (
	welcomeDesc = "Привет, %s! 👋\n\nBOOT: Быстрые опросы без ожидания! ⚡️🚀\n\nBOOT - это революционный сервис, который позволяет проводить опросы в считанные секунды! 🤯 \n\nЗабудьте о долгих ожиданиях, когда нужно собрать мнения. С BOOT вы можете отправить опрос прямо сейчас и получить результаты мгновенно. ⏱️\n\nЧто делает BOOT особенным?\n\n• Мгновенная отправка: Отправьте опрос и получите ответы без задержек. 💨\n• Без ограничений: Создавайте опросы любой сложности и длины. 🤩\n• Гибкость: Используйте различные типы вопросов, включая текстовые, множественный выбор и шкалу оценок. 📝📊\n• Простота использования: Интуитивно понятный интерфейс, который не требует специальных навыков. ✨\n• Доступность: Используйте BOOT для любых целей: от маркетинговых исследований до сбора обратной связи от клиентов. 📈👥\n\nBOOT - это доступная роскошь!\n\n• Бесплатный пробный период: Отправьте первые 50 запросов бесплатно! 🎁\n• Низкие цены: С BOOT опросы доступны для любого бюджета. 💰\n• Современные технологии: Мы используем передовые технологии, чтобы обеспечить максимальную скорость и надежность. 💻\n\nBOOT - это идеальный инструмент для тех, кто ценит время и эффективность! ⏰🏆\n"
)

func Auth(id int64) (*sqlx.DB, error) {
	db, err := manage.Connection()
	if err != nil {
		return nil, err
	}

	isExist, err := manage.CheckUser(db, id)
	if err != nil {
		return nil, err
	}

	if !isExist {
		err = manage.AddUser(db, id)
		if err != nil {
			return nil, err
		}
	}
	return db, nil
}

func StartHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	db, err := Auth(update.Message.Chat.ID)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	_, err = b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   fmt.Sprintf(welcomeDesc, update.Message.Chat.Username),
		ReplyMarkup: &models.ReplyKeyboardMarkup{
			Keyboard: [][]models.KeyboardButton{
				{
					{Text: "🧠AI"},
					{Text: "👅Language"},
					{Text: "💰Баланс"},
				},
				{
					{Text: "📋Подписки"},
					{Text: "👤Инфо"},
				},
			},
		},
	})
	if err != nil {
		panic(err.Error())
	}
}
