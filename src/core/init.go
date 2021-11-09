package core

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

var Config = Configuration{}

func init(){
	env := "test"
	if len(os.Args) >= 2 {
		env = os.Args[1]
	}

	if data, err := ioutil.ReadFile("./resources/" + env + ".yml"); err == nil {
		if err = yaml.Unmarshal(data, &Config); err == nil {
			loadAll(Config)
		} else {
			os.Exit(-1)
		}
	} else {
		os.Exit(-1)
	}
	fmt.Println("config info:", Config)
}

func loadAll(c Configuration) {
	loadRedis(c)
	loadCache()
	loadOrm(c)
	loadClient(c)
}
