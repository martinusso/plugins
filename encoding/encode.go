package encoding

import (
	"encoding/base64"
	"strings"

	"github.com/go-chat-bot/bot"
)

const (
	invalidAmountOfParams = "Invalid amount of parameters"
	invalidParams         = "Invalid parameters"
)

func encode(command *bot.Cmd) (string, error) {
	if len(command.Args) < 2 {
		return invalidAmountOfParams, nil
	}

	var str string
	switch command.Args[0] {
	case "base64":
		s := strings.Join(command.Args[1:], " ")
		str = encodeBase64(s)
	default:
		return invalidParams, nil
	}

	return str, nil
}

func encodeBase64(str string) string {
	data := []byte(str)
	return base64.StdEncoding.EncodeToString(data)
}

func init() {
	bot.RegisterCommand(
		"encode",
		"Allows you encoding a value",
		"base64 enter here text to encode",
		encode)
}
