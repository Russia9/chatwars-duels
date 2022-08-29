package app

import (
	"gitea.russia9.dev/Russia9/chatwars-duels/messages"
	"github.com/rs/zerolog/log"
	"gopkg.in/tucnak/telebot.v2"
	"strconv"
)

func (a *App) Sender(channel chan messages.DuelMessage) {
	for {
		var message messages.DuelMessage
		message = <-channel
		if message.Winner.Tag != "" {
			message.Winner.Tag = "[" + message.Winner.Tag + "]"
		}
		if message.Loser.Tag != "" {
			message.Loser.Tag = "[" + message.Loser.Tag + "]"
		}

		msgString := "Winner: " +
			message.Winner.Castle +
			message.Winner.Tag +
			message.Winner.Name +
			" 🏅" + strconv.Itoa(message.Winner.Level) +
			" ❤" + strconv.Itoa(message.Winner.Health) + " \n" +
			"Loser: " +
			message.Loser.Castle +
			message.Loser.Tag +
			message.Loser.Name +
			" 🏅" + strconv.Itoa(message.Loser.Level) +
			" ❤" + strconv.Itoa(message.Loser.Health)

		if message.IsChallenge {
			msgString += "\n" + "<b>Friendly Duel</b>"
		}

		if message.IsGuildDuel {
			msgString += "\n" + "<b>Guild Duel</b>"
		}

		_, err := a.Bot.Send(a.Chat, msgString, telebot.ModeHTML)
		if err != nil {
			log.Error().Err(err).Send()
		}
	}
}
