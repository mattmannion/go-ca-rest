package envs

import (
	"github.com/spf13/viper"
)

type TCfg struct {
	Env    string `mapstructure:"ENV"`
	Host   string `mapstructure:"HOST"`
	Port   string `mapstructure:"PORT"`
	PgHost string `mapstructure:"PGHOST"`
	PgPort string `mapstructure:"PGPORT"`
	PgPw   string `mapstructure:"POSTGRES_PASSWORD"`
	PgDb   string `mapstructure:"POSTGRES_DB"`
	PgUn   string `mapstructure:"POSTGRES_USER"`
}

var (
	Cfg    TCfg
	PgConn string
)

const (
	Dev string = "dev"
)

func init() {
	viper.SetConfigFile(".env")

	viper.ReadInConfig()

	viper.Unmarshal(&Cfg)

	PgConn = "postgres://" + Cfg.PgUn + ":" + Cfg.PgPw + "@" + Cfg.PgHost + ":" + Cfg.PgPort + "/" + Cfg.PgDb
}
