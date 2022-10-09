package service

import (
	"context"

	"github.com/nikoksr/notify"
	"github.com/nikoksr/notify/service/telegram"
)

type Notifier interface {
	Notify(channel string, subject string, message string)
}

type notifier struct {
}

func NewNotifier() Notifier {
	return &notifier{}
}

func (n *notifier) Notify(channel string, subject string, message string) {
	telegramService, _ := telegram.New("<api-key>")
	telegramService.AddReceivers(-123) // chat id

	notify.UseServices(telegramService)

	err := notify.Send(
		context.Background(),
		subject,
		message,
	)

	if err != nil {
		panic(err)
	}
}
