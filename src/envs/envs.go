package envs

import (
	"fmt"

	"github.com/spf13/viper"
)

type TCfg struct {
	Env      string `mapstructure:"ENV"`
	Host     string `mapstructure:"HOST"`
	Port     string `mapstructure:"PORT"`
	PgHost   string `mapstructure:"PGHOST"`
	PgPort   string `mapstructure:"PGPORT"`
	PgPw     string `mapstructure:"POSTGRES_PASSWORD"`
	PgUn     string `mapstructure:"POSTGRES_USER"`
	PgDbMain string `mapstructure:"DB_MAIN"`
	PgDbTest string `mapstructure:"DB_TEST"`
}

var (
	Cfg    TCfg
	PgConn string
)

const (
	Dev string = "dev"
)

func InitCfg(path string) {
	viper.SetConfigFile(path)

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Failed to read config", err)
		return
	}

	err = viper.Unmarshal(&Cfg)
	if err != nil {
		fmt.Println("Failed to unmarshal env", err)
		return
	}

	PgConn = "postgres://" + Cfg.PgUn + ":" + Cfg.PgPw + "@" + Cfg.PgHost + ":" + Cfg.PgPort + "/" + Cfg.PgDbMain
}
