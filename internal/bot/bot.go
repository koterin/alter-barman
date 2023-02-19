package bot

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"
	tb "gopkg.in/telebot.v3"

	"alter-barman/config"
	"alter-barman/internal/controller"
)

var Bot = &tb.Bot{}

func Start(ctx context.Context) {
	config.Validate()

	level, err := log.ParseLevel(config.Args.LOG_LEVEL)
	if err != nil {
		log.Fatal(err)
	}

	//log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(level)

	settings := tb.Settings{
		Token: config.Args.TG_BOT_KEY,
		Poller: &tb.LongPoller{
			Timeout: 1 * time.Second,
		},
	}

	Bot, _ = tb.NewBot(settings)

	Bot.Handle("/start", controller.OnStart())
	Bot.Handle(tb.OnText, controller.OnText())

	go func() {
		Bot.Start()
	}()

	log.Info("AlterBarman Bot started")

	<-ctx.Done()

	log.Info("AlterBarman Bot stopped")
	Bot.Stop()
}
