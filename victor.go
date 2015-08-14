package main

import (
	"os"
	"os/signal"

	"github.com/brettbuddin/victor"
	"github.com/kyokomi/go-docomo/docomo"
)

func dialogueHandler(robot victor.Robot) {
	robot.HandleFunc(robot.Direct("(.+)"), func(s victor.State) {
		d := docomo.NewClient(os.Getenv("DOCOMO_DIALOGUE_API_KEY"))
		dialogue := docomo.DialogueRequest{
			Utt: &s.Params()[0],
		}

		get, _ := d.Dialogue.Get(dialogue, true)
		s.Chat().Send(s.Message().ChannelID(), get.Utt)
	})
}

func main() {
	bot := victor.New(victor.Config{
		Name:        "victor",
		ChatAdapter: "shell",
	})

	dialogueHandler(bot)
	go bot.Run()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)
	<-sigs

	bot.Stop()
}
