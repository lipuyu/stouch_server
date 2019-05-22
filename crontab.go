package main

import (
	"fmt"
	"github.com/whiteshtef/clockwork"
)

func main() {
	sched := clockwork.NewScheduler()
	sched.Schedule().Every(10).Seconds().Do(something)
	sched.Run()
}

func something() {
	fmt.Println("foo")
}
