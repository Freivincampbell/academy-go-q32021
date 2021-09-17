package config

import (
	"bytes"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/viper"
	"log"
	"os"
)

type config struct {
	Database struct {
		User                 string
		Password             string
		Net                  string
		Addr                 string
		DBName               string
		AllowNativePasswords bool
		Params               struct {
			ParseTime string
		}
	}
	Server struct {
		Address string
	}
}

var C config

func ReadConfig() {
	Config := &C

	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	var yamlExample = []byte(`
database:
  user: root
  password:
  net: tcp
  addr: 127.0.0.1:3306
  dbname: golang-clean-architecture
  allowNativePasswords: true
  params:
    parseTime: "true"

server:
  address: 8080
`)
	viper.AutomaticEnv()

	if err := viper.ReadConfig(bytes.NewBuffer(yamlExample)); err != nil {
		fmt.Println(err)
		log.Fatalln(err)
	}

	if err := viper.Unmarshal(&Config); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	spew.Dump(C)
}
