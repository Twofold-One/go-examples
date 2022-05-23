package appconfiguration

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Configuration struct {
	Server   Server `json:"server"`
	Database DB     `json:"database"`
}

type Server struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type DB struct {
	Host     string `json:"host"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func FileBasedConfig() {
	confFile, err := os.Open("/home/twofold_one/GitProjects/go/go-examples/basic/appconfiguration/config.json")
	if err != nil {
		panic(err)
	}
	defer confFile.Close()

	conf, err := ioutil.ReadAll(confFile)
	if err != nil {
		panic(err)
	}

	myConf := Configuration{}
	err = json.Unmarshal(conf, &myConf)
	if err != nil {
		panic(err)
	}
	// +flag adds field names
	fmt.Printf("%+v\n", myConf)
}
