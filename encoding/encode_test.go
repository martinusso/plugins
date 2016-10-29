package encoding

import (
	"testing"

	"github.com/go-chat-bot/bot"
)

func TestEncode(t *testing.T) {
	bot := &bot.Cmd{
		Command: "encode",
	}
	bot.Args = []string{"base64", "The Go Programming Language"}
	got, err := encode(bot)

	expected := "VGhlIEdvIFByb2dyYW1taW5nIExhbmd1YWdl"
	if err != nil {
		t.Errorf("Error should be nil => %s", err)
	}
	if got != expected {
		t.Errorf("Test failed. Expected: '%s', got:  '%s'", expected, got)
	}
}

func TestEncodeWithMultipleWords(t *testing.T) {
	bot := &bot.Cmd{
		Command: "encode",
	}
	bot.Args = []string{"base64", "The", "Go", "Programming", "Language"}
	got, err := encode(bot)

	expected := "VGhlIEdvIFByb2dyYW1taW5nIExhbmd1YWdl"
	if err != nil {
		t.Errorf("Error should be nil => %s", err)
	}
	if got != expected {
		t.Errorf("Test failed. Expected: '%s', got:  '%s'", expected, got)
	}
}

func TestEncodeWhenPassInvalidParam(t *testing.T) {
	bot := &bot.Cmd{
		Command: "encode",
	}
	bot.Args = []string{"invalid_code", "R28gaXMgYW4gb3BlbiBzb3VyY2UgcHJvZ3JhbW1pbmcgbGFuZ3VhZ2U="}
	got, err := encode(bot)

	if err != nil {
		t.Errorf("Error should be nil => %s", err)
	}
	if got != invalidParams {
		t.Errorf("Should return a error message when pass correct amount of params but invalid param")
	}
}

func TestEncodeWhenDontPassAnyParams(t *testing.T) {
	bot := &bot.Cmd{
		Command: "encode",
	}
	got, err := encode(bot)

	if err != nil {
		t.Errorf("Error should be nil => %s", err)
	}
	if got != invalidAmountOfParams {
		t.Errorf("Should return a error message when don't pass any params")
	}
}
