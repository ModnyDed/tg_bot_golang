package telegram

import telegram "project/clients/Telegram"

type Processor struct {
	tg *telegram.Client
	offset int
	// storage
}

func New(client *telegram.Client) {
	
}