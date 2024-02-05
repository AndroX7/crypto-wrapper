package config

import (
	"fmt"
	"reflect"
	"regexp"
	"time"

	"github.com/louvri/gob/object"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

func Get(path string) (config Configuration, err error) {
	mapAssign := func(env map[string]interface{}, key string, value interface{}) error {
		var tmp string
		var digitCheck = regexp.MustCompile(`^[0-9]+$`)
		tmp, ok := value.(string)
		if ok {
			n := len(tmp)
			if n > 0 && (tmp[n-1] == 'h' || tmp[n-1] == 's' || tmp[n-1] == 'm') && digitCheck.MatchString(tmp[:n-1]) {
				if result, err := time.ParseDuration(tmp); err != nil {
					return err
				} else {
					env[key] = result
				}

			} else if n > 0 && tmp[n-1] == 's' && (tmp[n-2] == 'n' || tmp[n-2] == 'm' || tmp[n-2] == 'u') && digitCheck.MatchString(tmp[:n-2]) {
				if result, err := time.ParseDuration(tmp); err != nil {
					return err
				} else {
					env[key] = result
				}
			} else {
				env[key] = tmp
			}
		} else {
			env[key] = value
		}
		return nil
	}
	env := make(map[string]interface{})
	var cfg Configuration
	for key, value := range DEFAULTS {
		err = mapAssign(env, key, value)
		if err != nil {
			return cfg, err
		}
	}
	err = mapstructure.WeakDecode(env, &cfg)
	if err != nil {
		return cfg, err
	}
	tmp, err := getViper(path)
	if err == nil {
		merge(&cfg, &tmp)
	} else {
		fmt.Println(err.Error())
	}
	return cfg, nil
}

func getViper(path string) (cfg Configuration, err error) {
	type configData struct {
		Data *Configuration `mapstructure:"data"`
	}
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("json")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return cfg, err
	}
	//var tmp configData
	var tmp configData
	err = viper.Unmarshal(&tmp)
	if err != nil {
		panic(err)
	} else if tmp.Data != nil {
		cfg = *tmp.Data
	} else {
		err = viper.Unmarshal(&cfg)
		if err != nil {
			panic(err)
		}
	}

	return
}

func merge(old, new *Configuration) {
	newEL := reflect.ValueOf(new).Elem()
	oldEL := reflect.ValueOf(old).Elem()
	for i := 0; i < newEL.NumField(); i++ {
		prop := newEL.Type().Field(i).Name
		newRef := newEL.FieldByName(prop)
		oldRef := oldEL.FieldByName(prop)
		if !object.IsEmpty(newRef) {
			err := object.Assign(oldRef, prop, object.Get(newRef))
			if err != nil {
				panic(err)
			}
		}
	}
}
