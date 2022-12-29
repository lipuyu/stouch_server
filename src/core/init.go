package core

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

var Config = Configuration{}

func init() {
	confFile := "./resources/test.yml"
	if len(os.Args) >= 2 {
		confFile = os.Args[1]
	}

	if data, err := ioutil.ReadFile(confFile); err == nil {
		if err = yaml.Unmarshal(data, &Config); err == nil {
			loadAll(Config)
		} else {
			os.Exit(-1)
		}
	} else {
		os.Exit(-1)
	}
}

func loadAll(c Configuration) {
	fmt.Println("init... config info:", Config)
	loadRedis(c)
	loadCache()
	loadOrm(c)
	loadLogConf(c)
}
