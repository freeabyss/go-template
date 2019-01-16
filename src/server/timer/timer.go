package timer

import (
	"github.com/robfig/cron"
)

func  SetupTimer () {
	c := cron.New()

	c.AddFunc("0 0/1 * * * *", HealthCheck)

	c.Start()

}
