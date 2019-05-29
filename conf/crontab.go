package conf

import (
	"fmt"
	"github.com/whiteshtef/clockwork"
)

func Run() {
	sched := clockwork.NewScheduler()
	sched.Schedule().Every(10).Seconds().Do(something)
	sched.Run()
}

func something() {
	res, err := Cache.Value("user")
	fmt.Println(res, err)
}
