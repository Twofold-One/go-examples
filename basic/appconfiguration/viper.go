package appconfiguration

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func ViperConf() {
	viper.SetConfigName("config")
	viper.AddConfigPath("/home/twofold_one/GitProjects/go/go-examples/basic/appconfiguration")
	err := viper.ReadInConfig()
	if err != nil {
		log.Panicf("Fatal error reading config file: %s\n", err)
	}
	// retreive all the existing configuration keys of the app
	fmt.Println(viper.AllKeys())
	// access the value of property username form the object database
	fmt.Println(viper.GetString("database.username"))
}
