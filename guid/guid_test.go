package guid

import (
	"strings"
	"testing"

	"github.com/go-chat-bot/bot"
)

const (
	guidSize = 36
)

func TestGUIDShouldReturnValidGUID(t *testing.T) {
	bot := &bot.Cmd{
		Command: "guid",
	}
	got, err := guid(bot)

	if err != nil {
		t.Errorf("Error should be nil => %s", err)
	}
	if len(got) != guidSize {
		t.Errorf("Expected a valid GUID, but it was %s instead.", got)
	}
	if strings.Split(got, "")[14] != "4" {
		t.Errorf("Expected a valid GUID version 4, but it was %s instead.", got)
	}
}

func TestGUIDShouldReturnUpperGUID(t *testing.T) {
	bot := &bot.Cmd{
		Command: "guid",
	}
	bot.Args = []string{"upper"}
	got, err := guid(bot)

	if err != nil {
		t.Errorf("Error should be nil => %s", err)
	}
	if got != strings.ToUpper(got) {
		t.Errorf("Should return a upper GUID, but it was %s instead.", got)
	}
}

func TestGUIDInvalidParam(t *testing.T) {
	bot := &bot.Cmd{
		Command: "guid",
	}
	bot.Args = []string{"lower"}
	got, err := guid(bot)

	if err != nil {
		t.Errorf("Error should be nil => %s", err)
	}
	if got != msgInvalidParam {
		t.Errorf("Test Failed. Expected %s, but it was %s instead.", msgInvalidParam, got)
	}
}

func TestGUIDInvalidAmountOfParams(t *testing.T) {
	bot := &bot.Cmd{
		Command: "guid",
	}
	bot.Args = []string{"upper", "lower"}
	got, err := guid(bot)

	if err != nil {
		t.Errorf("Error should be nil => %s", err)
	}
	if got != msgInvalidAmountOfParams {
		t.Errorf("Test Failed. Expected %s, but it was %s instead.", msgInvalidAmountOfParams, got)
	}
}
