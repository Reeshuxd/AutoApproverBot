//   Approver Bot
//   Copyright (C) 2021 Reeshuxd (@reeshuxd)

package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
)

func main() {
	bot, err := gotgbot.NewBot(
		os.Getenv("TOKEN"),
		&gotgbot.BotOpts{
			APIURL:      "",
			Client:      http.Client{},
			GetTimeout:  gotgbot.DefaultGetTimeout,
			PostTimeout: gotgbot.DefaultPostTimeout,
		},
	)
	if err != nil {
		fmt.Println("Failed to create bot:", err.Error())
	}
	updater := ext.NewUpdater(
		&ext.UpdaterOpts{
			ErrorLog: nil,
			DispatcherOpts: ext.DispatcherOpts{
				Error: func(b *gotgbot.Bot, ctx *ext.Context, err error) ext.DispatcherAction {
					fmt.Println("an error occurred while handling update:", err.Error())
					return ext.DispatcherActionNoop
				},
				Panic:       nil,
				ErrorLog:    nil,
				MaxRoutines: 0,
			},
		})
	dp := updater.Dispatcher

	// Commands
	dp.AddHandler(handlers.NewCommand("start", Start))
	dp.AddHandler(handlers.NewChatJoinRequest(nil, Approve))

	// Start Polling()
	poll := updater.StartPolling(bot, &ext.PollingOpts{DropPendingUpdates: true})
	if poll != nil {
		fmt.Println("Failed to start bot:", poll.Error())
	}

	fmt.Printf("@%s has been sucesfully started\n💝Made by @MW_Dump\n", bot.Username)
	updater.Idle()
}

func Start(bot *gotgbot.Bot, ctx *ext.Context) error {
	if ctx.EffectiveChat.Type != "private" {
		return nil
	}

	user := ctx.EffectiveSender.User
	text := `
<b>Hello <a href="tg://user?id=%v">%v</a></b>
I am a bot made for accepting newly coming join requests at the time they comes.
I am made with <a href="go.dev">golang</a> to give a better performance!

Bot made with 💝 by <a href="t.me/MoviesWorld_Chan_nel">𝕄𝕆𝕍𝕀𝔼𝕊 𝕎𝕆ℝ𝕃𝔻</a> for you!
<b>Movie Request Group </b> @MoviesWorld_Group2
	`
	ctx.EffectiveMessage.Reply(
		bot,
		fmt.Sprintf(text, user.Id, user.FirstName),
		&gotgbot.SendMessageOpts{
			ReplyMarkup: gotgbot.InlineKeyboardMarkup{
				InlineKeyboard: [][]gotgbot.InlineKeyboardButton{{
					{Text: "♻️ 𝙲𝙷𝙰𝙽𝙽𝙴𝙻 ", Url: "https://t.me/+R2m54zJe33wxMGRl"},
				}},
			},
			ParseMode:             "html",
			DisableWebPagePreview: true,
		},
	)
	return nil
}

func Approve(bot *gotgbot.Bot, ctx *ext.Context) error {
	_, err := bot.ApproveChatJoinRequest(ctx.EffectiveChat.Id, ctx.EffectiveSender.User.Id)
	if err != nil {
		fmt.Println("Error while approving:", err.Error())
	}
	return nil
}
