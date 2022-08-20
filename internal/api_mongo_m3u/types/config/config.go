package config

import "time"

type Config struct {
	MongoDB  *MongoDB  `yaml:"mongodb"`
	HTTP     *HTTP     `yaml:"http"`
	Telegram *Telegram `yaml:"telegram"`
}

type MongoDB struct {
	URI      string `yaml:"uri"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DataBase string `yaml:"data_base"`
}

type HTTP struct {
	Port               string        `yaml:"server_port"`
	ReadTimeout        time.Duration `yaml:"read_timeout"`
	WriteTimeout       time.Duration `yaml:"write_timeout"`
	MaxHeaderMegabytes int           `yaml:"max_header_megabytes"`
}

type Telegram struct {
	TelegramToken string `yaml:"telegram_token"`
	ChatID        string `yaml:"chat_id"`
}
