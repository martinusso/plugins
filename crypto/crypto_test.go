package crypto

import (
	"testing"

	"github.com/go-chat-bot/bot"
)

func TestCryptoShouldReturnErrorMessageWhenDontPassAnyParams(t *testing.T) {
	bot := &bot.Cmd{
		Command: "crypto",
	}
	got, err := crypto(bot)

	if err != nil {
		t.Errorf("Error should be nil => %s", err)
	}
	if got != invalidAmountOfParams {
		t.Errorf("Test failed. Should return a error message when don't pass any params.")
	}
}

func TestCryptoShouldReturnErrorMessageWhenPassAnInvalidAlgorithm(t *testing.T) {
	bot := &bot.Cmd{
		Command: "crypto",
	}
	bot.Args = []string{"invalidAlgorithm", "input data"}
	got, err := crypto(bot)

	if err != nil {
		t.Errorf("Error should be nil => %s", err)
	}
	if got != invalidParams {
		t.Errorf("Should return a error message when pass an invalid algorithm")
	}
}

func TestCryptoShouldEncryptUsingMD5(t *testing.T) {
	bot := &bot.Cmd{
		Command: "crypto",
	}
	bot.Args = []string{"md5", "go-chat-bot"}
	got, err := crypto(bot)
	expected := "1120d1df84fec8a0557e8737ac021651"

	if err != nil {
		t.Errorf("Error should be nil => %s", err)
	}
	if got != expected {
		t.Errorf("Test failed. Expected: '%s', got:  '%s'", expected, got)
	}
}

func TestCryptoShouldEncryptMultipleWordsUsingMD5(t *testing.T) {
	bot := &bot.Cmd{
		Command: "crypto",
	}
	bot.Args = []string{"md5", "The", "Go", "Programming", "Language"}
	got, err := crypto(bot)
	expected := "adb505803d3502f2f00c88365ab85bf0"

	if err != nil {
		t.Errorf("Error should be nil => %s", err)
	}
	if got != expected {
		t.Errorf("Test failed. Expected: '%s', got:  '%s'", expected, got)
	}
}

func TestCryptoShouldEncryptUsingSHA1(t *testing.T) {
	bot := &bot.Cmd{
		Command: "crypto",
	}
	bot.Args = []string{"sha1", "go-chat-bot"}
	got, err := crypto(bot)
	expected := "385ca248ffebb5ed7f62d1ea2b0545cff80ac18e"

	if err != nil {
		t.Errorf("Error should be nil => %s", err)
	}
	if got != expected {
		t.Errorf("Test failed. Expected: '%s', got:  '%s'", expected, got)
	}
}

func TestCryptoShouldEncryptMultipleWordsUsingSHA1(t *testing.T) {
	bot := &bot.Cmd{
		Command: "crypto",
	}
	bot.Args = []string{"sha-1", "The", "Go", "Programming", "Language"}
	got, err := crypto(bot)
	expected := "88a93e668044877a845097aaf620532a232bfd34"

	if err != nil {
		t.Errorf("Error should be nil => %s", err)
	}
	if got != expected {
		t.Errorf("Test failed. Expected: '%s', got:  '%s'", expected, got)
	}
}
