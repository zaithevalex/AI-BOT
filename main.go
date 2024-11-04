package main

import (
	"BOOT-BOT/db/manage"
	"BOOT-BOT/handlers/general"
	ai "BOOT-BOT/handlers/gpt"
	"context"
	"github.com/go-telegram/bot"
	"os"
	"os/signal"
)

const (
	accessToken = "7858831026:AAEAkK2OF3A0lx5ggT509OvmKOkSvpoyrs0"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	db, err := manage.Connection()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	if _, err = db.Query(manage.CreateTableUsersReq); err != nil {
		panic(err.Error())
	}

	if _, err = db.Query(manage.CreateTableUsersSubscriptions); err != nil {
		panic(err.Error())
	}

	opts := []bot.Option{
		bot.WithDefaultHandler(ai.GPTHandler),
		bot.WithCallbackQueryDataHandler("button", bot.MatchTypePrefix, ai.GeneralButtonHandler),
		bot.WithMessageTextHandler("/start", bot.MatchTypeExact, general.StartHandler),
		bot.WithMessageTextHandler("🧠AI", bot.MatchTypeExact, ai.PickGPTHandler),
		bot.WithMessageTextHandler("📋Подписки", bot.MatchTypeExact, nil),
	}

	b, err := bot.New(accessToken, opts...)
	if err != nil {
		panic(err.Error())
	}

	b.Start(ctx)
}
