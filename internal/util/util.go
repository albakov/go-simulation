package util

import "fmt"

const (
	MessageContinueOrExit     = "[П]ауза/продолжить\n[В]ыйти\n"
	MessageNoReasonToContinue = "Нет причин для продолжения... Завершение работы.\n"
	MessageCounter            = "Количество итераций: %d\n"

	MessageBegin            = "[В]ыйти\n[Н]ачать симуляцию?\n"
	MessageExit             = "Завершение работы. Количество циклов: %v\n"
	MessageIncorrectCommand = "Не правильно введена команда!\n"
	MessageContinue         = "Продолжение...\n"
	MessagePause            = "Пауза...\n"

	MessagePrefix    = "-> "
	MessageBorder    = "===="
	MessageEmptyCell = " .. "
	MessageNewLine   = "\n"
)

func ShowMessage(message string, args ...any) {
	fmt.Printf(message, args...)
}
