package main

import (
	"os"

	"github.com/danryan/hal"
	_ "github.com/danryan/hal/adapter/shell"
	_ "github.com/danryan/hal/store/memory"
	"github.com/kyokomi/go-docomo/docomo"
)

var dialogueHandler = hal.Respond(`(.+)`, func(res *hal.Response) error {
	d := docomo.NewClient(os.Getenv("DOCOMO_DIALOGUE_API_KEY"))
	dialogue := docomo.DialogueRequest{
		Utt: &res.Match[1],
	}

	get, err := d.Dialogue.Get(dialogue, true)
	if err != nil {
		return err
	}

	return res.Reply(get.Utt)
})

func main() {
	robot, err := hal.NewRobot()
	if err != nil {
		hal.Logger.Error(err)
	}

	robot.Handle(
		dialogueHandler,
	)

	if err := robot.Run(); err != nil {
		hal.Logger.Error(err)
	}
}
