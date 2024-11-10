package general

import (
	"BOOT-BOT/handlers/gpt"
	"github.com/jmoiron/sqlx"
)

type Action struct {
	tag     string
	db      *sqlx.DB
	buttons []*gpt.Button
	body    func() error
}

type ListAction []Action

func (l *ListAction) Create() {

}
