package envs

import (
	"github.com/spf13/viper"
)

type cfg struct {
	FirestoreProjectName    string `mapstructure:"FIRESTORE_PROJECTNAME"`
	FirestoreCollectionName string `mapstructure:"FIRESTORE_COLLECTIONNAME"`
	FirestoreJson           string `mapstructure:"FIRESTORE_JSON"`
	PgUn                    string `mapstructure:"PGUN"`
	PgPw                    string `mapstructure:"PGPW"`
	PgDb                    string `mapstructure:"PGDB"`
}

var (
	Cfg    cfg
	PgConn string
)

func init() {
	viper.SetConfigFile(".env")

	viper.ReadInConfig()

	viper.Unmarshal(&Cfg)

	PgConn = "postgres://" + Cfg.PgUn + ":" + Cfg.PgPw + "@localhost:5432/" + Cfg.PgDb
}
