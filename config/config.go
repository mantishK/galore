package config

import (
	"flag"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

var y map[string]interface{}

//initialize the configs by reading them from the config files
func init() {
	initYaml()
	initCmdArgs()
}

func initYaml() {
	configFilePath := os.Getenv("GALORE_CONFIG")
	// Use local config file
	if len(configFilePath) == 0 {
		seperator := string(os.PathSeparator)
		configFilePath = "config" + seperator + "config.yaml"
	}
	config, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		log.Println("Warning: config file not found")
		return
	}
	y = make(map[string]interface{})
	err = yaml.Unmarshal([]byte(config), &y)
	if err != nil {
		log.Fatalf("Error reading yaml file: %v", err)
	}
}

func initCmdArgs() {
	t := make(map[string]interface{})
	for key, _ := range y {
		if _, ok := y[key].(string); ok {
			t[key] = flag.String(key, "", "")
		}
		if _, ok := y[key].(int); ok {
			t[key] = flag.Int(key, 0, "")
		}
		if _, ok := y[key].(bool); ok {
			t[key] = flag.Bool(key, false, "")
		}
	}
	flag.Parse()
	for key, _ := range y {
		if _, ok := y[key].(string); ok {
			if x, ok := t[key].(*string); ok && len(*x) > 0 {
				y[key] = *x
			}
		}
		if _, ok := y[key].(int); ok {
			if x, ok := t[key].(*int); ok && *x > 0 {
				y[key] = *x
			}
		}
		if _, ok := y[key].(bool); ok {
			if x, ok := t[key].(*bool); ok && *x != false {
				y[key] = *x
			}
		}
	}
}

func GetString(key string) (val string, ok bool) {
	val, ok = y[key].(string)
	return
}

func GetInt(key string) (val int, ok bool) {
	val, ok = y[key].(int)
	return
}

func GetBool(key string) (val bool, ok bool) {
	val, ok = y[key].(bool)
	return
}
