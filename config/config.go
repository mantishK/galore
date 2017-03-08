package config

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"gopkg.in/yaml.v2"
)

var y map[string]interface{}
var args map[string]interface{}

//initialize the configs by reading them from the config files
func init() {
	initCmdArgs()
	initYaml()

}

func initCmdArgs() {
	args = map[string]interface{}{
		"pg_uname":    nil,
		"pg_pass":     nil,
		"pg_name":     nil,
		"pg_ip":       nil,
		"pg_port":     nil,
		"redis_uname": nil,
		"redis_pass":  nil,
		"redis_port":  nil,
		"redis_ip":    nil,
	}
	for key, _ := range args {
		args[key] = flag.String(key, "", "")
	}
	flag.Parse()
}

func initYaml() {
	var s = string(os.PathSeparator)
	config, err := ioutil.ReadFile(s + "etc" + s + "galore" + s + "config.yaml")
	if err != nil {
		log.Println("Warning: config file /etc/galore/config.yaml not found")
	} else {
		y = make(map[string]interface{})
		err = yaml.Unmarshal([]byte(config), &y)
		if err != nil {
			log.Fatalf("Error reading yaml file: %v", err)
		}
	}
}

func GetString(key string) (val string, ok bool) {
	valPtr, ok := args[key].(*string)
	if ok && len(*valPtr) != 0 {
		val = *valPtr
	}
	val, ok = y[key].(string)
	if ok {
		return
	}
	return
}

func GetInt(key string) (val int, ok bool) {
	var err error
	valPtr, ok := args[key].(*string)
	if ok && len(*valPtr) != 0 {
		val, err = strconv.Atoi(*valPtr)
		if err != nil {
			panic(key + " is not of type int")
		}
	}
	val, ok = y[key].(int)
	if ok {
		return
	}
	return
}

func GetBool(key string) (val bool, ok bool) {
	valPtr, ok := args[key].(*string)
	if ok && len(*valPtr) != 0 {
		if *valPtr == strings.ToUpper("true") {
			val = true
		} else if *valPtr == strings.ToUpper("false") {
			val = false
		} else {
			panic(key + " is not of type bool")
		}
	}
	val, ok = y[key].(bool)
	if ok {
		return
	}
	return
}
