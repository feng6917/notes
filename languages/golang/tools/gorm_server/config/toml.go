package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

type Config struct {
	EnvConfig *EnvConfig           `toml:"env"`
	DBConfigs map[string]*DBConfig `toml:"data"`
}

func (c *Config) String() string {
	return fmt.Sprintf("%+v", *c)
}

type EnvConfig struct {
	Mode  string `toml:"mode"`
	Debug bool   `toml:"debug"`
	Host  string `toml:"host"`
	Port  int    `toml:"port"`
}

func (c *EnvConfig) String() string {
	return fmt.Sprintf("%+v", *c)
}

type DBData struct {
}

type DBConfig struct {
	Driver      string `toml:"driver"`
	Host        string `toml:"host"`
	Port        int    `toml:"port"`
	Name        string `toml:"name"`
	Password    string `toml:"password"`
	Db          string `toml:"data"`
	Charset     string `toml:"charset"`
	Loc         string `toml:"loc"`
	Singular    bool   `toml:"singular"`
	MaxIdConn   int    `toml:"max_id_conn"`
	MaxOpenConn int    `toml:"max_open_conn"`
	TablePeople string `toml:"table_people"`
}

func (c *DBConfig) String() string {
	return fmt.Sprintf("%+v", *c)
}

func Init(file string) (conf *Config, err error) {
	if file != "" {
		_, err = toml.DecodeFile(file, &conf)
		if err != nil {
			return conf, err
		}
	}
	return conf, nil
}
