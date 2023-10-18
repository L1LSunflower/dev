package factory

import (
	"errors"
	"fmt"
)

type Messenger interface {
	Send() string
}

const (
	DEFAULT_MESSAGE = "message send by"

	TELEGRAM_MESSENGER = "telegram"
	WHATSAPP_MESSENGER = "whatsapp"
	VIBER_MESSENGER    = "viber"
)

var ErrMessengerNotFound = errors.New("messenger not found")
var messengerMap = map[string]Messenger{
	WHATSAPP_MESSENGER: new(WhatsApp),
	TELEGRAM_MESSENGER: new(Telegram),
	VIBER_MESSENGER:    new(Viber),
}

type WhatsApp struct {
	// Может быть драйвер для Whatsapp
}

func (w *WhatsApp) Send() string {
	// Реализация отправки сообщение по WhatsApp
	return fmt.Sprintf("%s %s", DEFAULT_MESSAGE, WHATSAPP_MESSENGER)
}

type Telegram struct {
	// Может быть драйвер для Telegram
}

func (t *Telegram) Send() string {
	// Реализация отправки сообщение по Telegram
	return fmt.Sprintf("%s %s", DEFAULT_MESSAGE, TELEGRAM_MESSENGER)
}

type Viber struct {
	// Может быть драйвер для Viber
}

func (v *Viber) Send() string {
	// Реализация отправки сообщение по Viber
	return fmt.Sprintf("%s %s", DEFAULT_MESSAGE, VIBER_MESSENGER)
}

func MessengerFactory(messengerType string) (Messenger, error) {
	messenger, ok := messengerMap[messengerType]
	if !ok {
		return nil, ErrMessengerNotFound
	}

	return messenger, nil
}
