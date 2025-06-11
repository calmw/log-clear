package main

import (
	"github.com/jasonlvhit/gocron"
	"log-clear/clear"
	"os"
)

func main() {
	clear.LogDir = os.Args[1]
	clear.DeleteLog()
	s := gocron.NewScheduler()
	_ = s.Every(1800).Seconds().From(gocron.NextTick()).Do(clear.DeleteLog)
	<-s.Start()
}
