package conf

import (
	"github.com/BurntSushi/toml"
)

type config struct {
	Title  string `toml:"title"`
	Owner  Owner  `toml:"owner"`
	HTTP   HTTP   `toml:"http"`
	Log    Log    `toml:"log"`
	Wechat Wechat `toml:"wechat"`
}
type Owner struct {
	Name string `toml:"name"`
}
type HTTP struct {
	Addr string `toml:"addr"`
}
type Log struct {
	Path string `toml:"path"`
}
type Wechat struct {
	Appid  string `toml:"appid"`
	Secret string `toml:"secret"`
}

var Config *config

func Init(confPath string) error {
	if _, err := toml.DecodeFile(confPath, &Config); err != nil {
		return err
	}
	return nil
}

func Get() *config {
	return Config
}
