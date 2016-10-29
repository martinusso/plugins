package encoding

import (
	"strings"
	"testing"

	"github.com/go-chat-bot/bot"
)

func TestDecode(t *testing.T) {
	bot := &bot.Cmd{
		Command: "decode",
	}
	bot.Args = []string{"base64", "R28gaXMgYW4gb3BlbiBzb3VyY2UgcHJvZ3JhbW1pbmcgbGFuZ3VhZ2U="}
	got, err := decode(bot)
	expected := "Go is an open source programming language"

	if err != nil {
		t.Errorf("Error should be nil => %s", err)
	}
	if got != expected {
		t.Errorf("Test failed. Expected: '%s', got:  '%s'", expected, got)
	}
}
func TestDecodeWhenPassInvalidHash(t *testing.T) {
	bot := &bot.Cmd{
		Command: "decode",
	}
	bot.Args = []string{"base64", "R28gaXMgYW4gb3BlbiBzb3VyY2Ugc", "HJvZ3JhbW1pbmcgbGFuZ3VhZ2U="}
	got, err := decode(bot)

	if err != nil {
		t.Errorf("Error should be nil => %s", err)
	}
	if !strings.HasPrefix(got, "Error:") {
		t.Errorf("Should return a error message when pass a invalid hash")
	}
}
func TestDecodeWhenPassInvalidParam(t *testing.T) {
	bot := &bot.Cmd{
		Command: "decode",
	}
	bot.Args = []string{"invalid_code", "R28gaXMgYW4gb3BlbiBzb3VyY2UgcHJvZ3JhbW1pbmcgbGFuZ3VhZ2U="}
	got, err := decode(bot)

	if err != nil {
		t.Errorf("Error should be nil => %s", err)
	}
	if got != invalidParams {
		t.Errorf("Should return a error message when pass correct amount of params but invalid param")
	}
}
func TestDecodeWhenDontPassAnyParams(t *testing.T) {
	bot := &bot.Cmd{
		Command: "decode",
	}
	got, err := decode(bot)

	if err != nil {
		t.Errorf("Error should be nil => %s", err)
	}
	if got != invalidAmountOfParams {
		t.Errorf("Should return a error message when don't pass any params")
	}
}
