package main

import (
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"

	"github.com/brettbuddin/victor"
	"github.com/kyokomi/go-docomo/docomo"
)

var bot victor.Robot

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

func sayHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	bot.Chat().Send("0", string(body))
}

func main() {
	bot = victor.New(victor.Config{
		Name:        "victor",
		ChatAdapter: "shell",
	})

	dialogueHandler(bot)
	go bot.Run()

	http.HandleFunc("/say", sayHandler)
	go http.ListenAndServe(":9000", nil)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)
	<-sigs

	bot.Stop()
}
