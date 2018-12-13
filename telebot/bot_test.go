package telebot_test

import (
	"testing"

	"github.com/yosh0/telegram"
	"github.com/yosh0/telegram/telebot"
	"golang.org/x/net/context"
	"gopkg.in/stretchr/testify.v1/assert"
)

func TestGetAPI(t *testing.T) {
	bg := context.Background()
	api := telegram.New("")

	ctx := telebot.WithAPI(bg, api)
	assert.NotEqual(t, bg, ctx)
	assert.Equal(t, api, telebot.GetAPI(ctx))

	assert.Panics(t, func() {
		telebot.GetAPI(bg)
	})
}

func TestGetUpdate(t *testing.T) {
	bg := context.Background()
	update := &telegram.Update{}

	ctx := telebot.WithUpdate(bg, update)
	assert.NotEqual(t, bg, ctx)
	assert.Equal(t, update, telebot.GetUpdate(ctx))

	assert.Panics(t, func() {
		telebot.GetUpdate(bg)
	})
}
