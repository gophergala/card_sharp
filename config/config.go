package config

import (
	"flag"
	"log"
	"os"

	"git.andrewcsellers.com/acsellers/chat/store"
	"github.com/BurntSushi/toml"
)

var (
	Conn     *store.Conn
	Dev      = flag.Bool("dev", false, "Development mode, load local config")
	ConfPath = flag.String("conf", "/etc/card_party/solo.conf", "Production Config file location")
	Config   Settings
	Cookie   *securecookie.Cookie
)

var Settings struct {
	WebPort      int
	ResourcePath string
	SQLAddr      string
	SQLType      string
}

func init() {
	flag.Parse()

	var confPath string
	if Dev {
		if _, err := os.Stat("solo.conf"); err != nil {
			log.Fatal("Missing solo.conf for config settings")
		}
	} else {
		confPath = *ConfPath
	}
	if _, err := toml.DecodeFile(confPath, &Config); err != nil {
		log.Fatal("Parse Config File", err)
	}
	var err error
	if Conn, err = store.Open(Config.SQLType, Config.SQLAddr); err != nil {
		log.Fatal("Open SQL Connection", Config, err)
	}

	Cookie = securecookie.New(
		securecookie.GenerateRandomKey(32),
		securecookie.GenerateRandomKey(32),
	)
}
