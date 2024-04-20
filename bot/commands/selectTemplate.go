package commands

import (
	"log"
	//"root/database/models"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (hb *HomeworkBot) addTask(update tgbotapi.Update) {
	state := &hb.state

	if update.Message.Text == "/selectPromt" || update.Message.Text == "Выберите шаблон" {
		keyboard := tgbotapi.NewReplyKeyboard(
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton("Выберите шаблон"),
			),
		)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Выберите шаблон:")
		msg.ReplyMarkup = keyboard

		if _, err := hb.bot.Send(msg); err != nil {
			log.Println("Ошибка: не удалось отправить сообщение. \n", err)
		}
	} else if update.Message.Text == "Домашнее задание" {
		state.IsSubjectInput = true
		state.IsTaskInput = false
		state.SubjectName = ""
		state.Task = ""
		state.IsAddTask = true

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Загрузите фотографию")
		if _, err := hb.bot.Send(msg); err != nil {
			log.Println("Ошибка: не удалось отправит сообщение. \n", err)
		}
	} else if update.Message.Text == "Тест" {
		state.IsSubjectInput = true
		state.IsTaskInput = false
		state.SubjectName = ""
		state.Task = ""
		state.IsAddTask = false

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Загрузите фотографию")
		if _, err := hb.bot.Send(msg); err != nil {
			log.Println("Ошибка: не удалось отправит сообщение. \n", err)
		}
	}
}



