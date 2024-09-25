package main

import (
	"github.com/jasonlvhit/gocron"
	"log-clear/clear"
)

func main() {

	s := gocron.NewScheduler()
	_ = s.Every(1800).Seconds().From(gocron.NextTick()).Do(clear.DeleteLog)
	<-s.Start()
}
