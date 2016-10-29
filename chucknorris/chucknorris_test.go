package chucknorris

import (
	"testing"

	"github.com/go-chat-bot/bot"
)

func TestChuckNorrisWhenTheTextDoesNotMatchAChuckNorrisName(t *testing.T) {
	cmd := &bot.PassiveCmd{}
	cmd.Raw = "My name is go-bot, I am awesome."
	got, err := chucknorris(cmd)

	if err != nil {
		t.Errorf("Error should be nil => %s", err)
	}
	if got != "" {
		t.Errorf("Test failed. Expected a empty return, got:  '%s'", got)
	}
}

func TestChuckNorrisWhenTheTextMatchChuck(t *testing.T) {
	cmd := &bot.PassiveCmd{}
	cmd.Raw = "My name is chuck"
	got, err := chucknorris(cmd)

	if err != nil {
		t.Errorf("Error should be nil => %s", err)
	}
	if got == "" {
		t.Errorf("Test failed. Should return a Chuck Norris facts")
	}
}

func TestChuckNorrisWhenTheTextMatchNorris(t *testing.T) {
	cmd := &bot.PassiveCmd{}
	cmd.Raw = "Hi, I'm Mr. Norris"
	got, err := chucknorris(cmd)

	if err != nil {
		t.Errorf("Error should be nil => %s", err)
	}
	if got == "" {
		t.Errorf("Test failed. Should return a Chuck Norris facts")
	}
}
