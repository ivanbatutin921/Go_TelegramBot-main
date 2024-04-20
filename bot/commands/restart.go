package commands

import (
	"log"
	//"root/database/models"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)


func (hb *HomeworkBot) restartCommand(update tgbotapi.Update) {
    keyboard := tgbotapi.NewReplyKeyboard(
        tgbotapi.NewKeyboardButtonRow(
            tgbotapi.NewKeyboardButton("restart"),
        ),
    )

    msg := tgbotapi.NewMessage(update.Message.Chat.ID, "restatr")
    msg.ReplyMarkup = keyboard

    if _, err := hb.bot.Send(msg); err != nil {
        log.Println("Ошибка: не удалось отправить сообщение\n", err)
    }

    // Add your logic for restarting the bot here
}