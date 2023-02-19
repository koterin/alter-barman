package controller

import (
	"fmt"
	"strconv"

	log "github.com/sirupsen/logrus"
	tb "gopkg.in/telebot.v3"

	"alter-barman/internal/entity"
	"alter-barman/internal/usecase"
)

func OnStart() tb.HandlerFunc {
	return func(c tb.Context) error {
		if !c.Message().Private() {
			log.Error("Error: chat is not private with user ", c.Sender().Username)

			return c.Send(entity.TextInternalError)
		}

		log.Info("User with name = ", c.Message().Sender.Username, " and ID = ", strconv.Itoa(int(c.Message().Chat.ID)), " started bot")

		return c.Send(fmt.Sprintf(entity.TextInstructions, c.Message().Sender.Username))
	}
}

func OnText() tb.HandlerFunc {
	return func(c tb.Context) error {
		log.Info("User with name = ", c.Message().Sender.Username, " and ID = ", strconv.Itoa(int(c.Message().Chat.ID)), " sent message: ", c.Message().Text)

		usecase.SendNewRecord(int(c.Message().Chat.ID), c.Message().Sender.Username, c.Message().Text)

		return c.Send(entity.TextReceived)
	}
}
