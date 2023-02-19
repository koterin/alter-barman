package usecase

import (
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"

	"alter-barman/config"
)

var TGClient = &http.Client{Timeout: 10 * time.Second}

func SendNewRecord(chatID int, user string, msg string) {
	rawURL := config.Args.TG_API_URL + "bot" + config.Args.TG_BOT_KEY + "/sendMessage"

	reqURL, err := url.Parse(rawURL)
	if err != nil {
		log.Error("Error parsing URL in .SendNewRecord: ", err)
	}

	alert := "**НОВАЯ ЗАПИСЬ:**\n" + user + " внес запись:\n" + msg

	q := reqURL.Query()
	q.Add("chat_id", strconv.Itoa(chatID))
	q.Add("text", alert)
	q.Add("parse_mode", "markdown")
	reqURL.RawQuery = q.Encode()

	resp, err := TGClient.Get(reqURL.String())
	if err != nil {
		log.Error("Error sending request to admin: ", err)

		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error("Error reading TG response body to admin: ", err)

		return
	}

	// TODO: parse into normal struct.
	sb := string(body)[6:10]

	if sb != "true" {
		log.Error("Error send status from Telegram to admin: ", string(body))

		return
	}

	return
}
