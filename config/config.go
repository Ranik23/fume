package config

import (
	"io/ioutil"
	"os"
	"github.com/goccy/go-yaml"
)


type DataBase struct {
	Host 		string `yaml:"host"`
	Port 		string `yaml:"port"`
	User 		string `yaml:"user"`
	Password 	string `yaml:"password"`
	Name		string `yaml:"name"`
	Ssl			string `yaml:"ssl"`
}

type ServerConfig struct {
	Host 		string `yaml:"host"`
	Port 		string `yaml:"port"`
}


type Config struct {
	Dictionary 		[]string 		`yaml:"dictionary"`
	Alphabet 		[]string 		`yaml:"alphabet"`
	Env 			string	 		`yaml:"environment"`
	DataBase		DataBase 		`yaml:"database"`
	ServerConfig	ServerConfig	`yaml:"server"`
}


func GetConfig(path string) (*Config, error)  {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var config Config

	err = yaml.Unmarshal([]byte(data), &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}