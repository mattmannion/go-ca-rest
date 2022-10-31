package envs

import (
	"github.com/spf13/viper"
)

type cfg struct {
	Env    string `mapstructure:"ENV"`
	PgHost string `mapstructure:"PGHOST"`
	PgPort string `mapstructure:"PGPORT"`
	PgUn   string `mapstructure:"PGUN"`
	PgPw   string `mapstructure:"PGPW"`
	PgDb   string `mapstructure:"PGDB"`
}

var (
	Cfg    cfg
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
