package main

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/danryan/hal"
	_ "github.com/danryan/hal/adapter/shell"
	_ "github.com/danryan/hal/store/memory"
	"github.com/kyokomi/go-docomo/docomo"
)

var robot *hal.Robot

func dialogueHandler(res *hal.Response) error {
	d := docomo.NewClient(os.Getenv("DOCOMO_DIALOGUE_API_KEY"))
	dialogue := docomo.DialogueRequest{
		Utt: &res.Match[1],
	}

	get, err := d.Dialogue.Get(dialogue, true)
	if err != nil {
		return err
	}

	return res.Reply(get.Utt)
}

func sayHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	robot.Adapter.Send(nil, string(body))
}

func main() {
	robot, _ = hal.NewRobot()

	robot.Handle(
		hal.Respond(`(.+)`, dialogueHandler),
	)
	hal.Router.HandleFunc("/say", sayHandler)

	robot.Run()
}
