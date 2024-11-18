package gpt

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/jmoiron/sqlx"
	"strings"
)

type Invoice struct {
	title       string
	description string
	payload     string
	currency    string
	label       string
	price       int
}

type Message struct {
	// button params
	Relate    string
	Name      string
	ButtonTag string

	// text message
	text string

	// payment params
	invoice Invoice
}

type Node struct {
	button *Message
	nodes  []*Node
	parent *Node
}

func AddCheckMark(buttons []*Message, name string) int {
	name = strings.Trim(name, " ")
	for index, button := range buttons {
		if button.Name == name {
			button.Name += "✅"
			return index
		}
	}
	return -1
}

func ActionBypass(callBackQuery string, mainRoot *Node, ctx context.Context, bot *bot.Bot, update *models.Update) error {
	for _, node := range mainRoot.nodes {
		if len(node.button.text) > 0 {

		}
	}
}

func Init(db *sqlx.DB) (*Node, error) {
	var buttons []*Message
	err := db.Select(&buttons, "select * from message;")
	if err != nil {
		return nil, err
	}

	var roots []Node
	for _, button := range buttons {
		if len(button.Relate) == 0 {
			roots = append(roots, Node{button: &Message{Relate: button.Relate, Name: button.Name, ButtonTag: button.ButtonTag}})
		}
	}

	var bypassRoots []Node
	mainRoot := &Node{}
	for _, root := range roots {
		root.parent = mainRoot
		mainRoot.nodes = append(mainRoot.nodes, &root)
		r := &root
		nodesInit(&r, buttons)
		bypassRoots = append(bypassRoots, *r)
	}

	return mainRoot, nil
}

func InlineKeyboardMarkUpGenerate(buttons []*Message) models.InlineKeyboardMarkup {
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

func nodesInit(node **Node, buttons []*Message) {
	for _, button := range buttons {
		if (*node).parent.button.Name == button.Relate {
			child := &Node{button: button, parent: *node}
			(*node).nodes = append((*node).nodes, child)
		}
	}

	for _, n := range (*node).nodes {
		nodesInit(&n, buttons)
	}
}
