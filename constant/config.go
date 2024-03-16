package constant

import (
	"sync"

	"github.com/BurntSushi/toml"
)

type Cfg struct {
	AccessToken string `toml:"access_token"`
	Proxies     string `toml:"proxies"`
	Addr        string `toml:"addr" default:":8080"`
}

var (
	cfg  *Cfg
	once sync.Once
)

func Config() *Cfg {
	once.Do(func() {
		filepath := "constant/config.toml"
		_, err := toml.DecodeFile(filepath, &cfg)
		if err != nil {
			panic(err)
		}
	})
	return cfg
}
