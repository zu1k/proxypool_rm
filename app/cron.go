package app

import (
	"github.com/jasonlvhit/gocron"
)

func Cron() {
	_ = gocron.Every(15).Minutes().Do(CrawlGo)
	<-gocron.Start()
}
