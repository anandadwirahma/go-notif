package telegram

import (
	"encoding/json"
	"fmt"
	"net/http"
	netUrl "net/url"

	"go-notif/notification"
	"go-notif/shared/config"
)

type Telegram struct {
	token string
}

func NewTelegramChannel(token string) notification.Notifier {
	return &Telegram{
		token: token,
	}
}

func (t *Telegram) SendTextMessage(message string, recipients []string) error {
	data := netUrl.Values{}
	url := fmt.Sprintf("%s%s%s", config.TelegramAddress, t.token, config.TelegramPathSendMessage)

	for _, id := range recipients {
		data.Set("text", message)
		data.Set("chat_id", id)

		resp, err := http.PostForm(url, data)
		if err != nil {
			return err
		}

		if resp.StatusCode != http.StatusOK {
			r := &ErrorResponse{}
			err := json.NewDecoder(resp.Body).Decode(&r)
			if err != nil {
				return err
			} else if !r.Ok {
				err = fmt.Errorf("error send meessage:%s:%d:%s", id, r.ErrorCode, r.Description)
				return err
			}
		}

		err = resp.Body.Close()
		if err != nil {
			return err
		}
	}

	return nil
}
