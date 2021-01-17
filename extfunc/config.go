package extfunc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	// "log"
	"errors"
	"os"
	"reflect"
	"time"
)

// Config 包含username 和 password
type Config struct {
	Username string `json:"username"`
	Password string `json:"password"`
	URLCheck string `json:"urlCheck"`
	URLPort  string `json:"urlPort"`
	DNSCheck string `json:"dnsCheck"`
	//Timeout     int       `json:"timeout"`
}

var readTimeout time.Duration

// IsFileExists 判断文件是否存在
func IsFileExists(path string) (bool, error) {
	stat, err := os.Stat(path)
	if err == nil {
		if stat.Mode()&os.ModeType == 0 {
			return true, nil
		}
		return false, errors.New(path + " exists but is not regular file")
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// ParseConfig 读配置
func ParseConfig(path string) (config *Config, err error) {
	file, err := os.Open(path) // For read access.
	if err != nil {
		return
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return
	}

	config = &Config{}
	if err = json.Unmarshal(data, config); err != nil {
		return nil, err
	}
	//readTimeout = time.Duration(config.Timeout) * time.Second
	return
}

// Useful for command line to override options specified in config file
// Debug is not updated.

// UpdateConfig 更新配置
func UpdateConfig(old, new *Config) {
	// Using reflection here is not necessary, but it's a good exercise.
	// For more information on reflections in Go, read "The Laws of Reflection"
	// http://golang.org/doc/articles/laws_of_reflection.html
	newVal := reflect.ValueOf(new).Elem()
	oldVal := reflect.ValueOf(old).Elem()

	// typeOfT := newVal.Type()
	for i := 0; i < newVal.NumField(); i++ {
		newField := newVal.Field(i)
		oldField := oldVal.Field(i)
		// log.Printf("%d: %s %s = %v\n", i,
		// typeOfT.Field(i).Name, newField.Type(), newField.Interface())
		switch newField.Kind() {
		case reflect.Interface:
			if fmt.Sprintf("%v", newField.Interface()) != "" {
				oldField.Set(newField)
			}
		case reflect.String:
			s := newField.String()
			if s != "" {
				oldField.SetString(s)
			}
		case reflect.Int:
			i := newField.Int()
			if i != 0 {
				oldField.SetInt(i)
			}
		}
	}

	//old.Timeout = new.Timeout
	//readTimeout = time.Duration(old.Timeout) * time.Second
}
