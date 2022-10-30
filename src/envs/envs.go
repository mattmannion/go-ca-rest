package envs

import (
	"github.com/spf13/viper"
)

type cfg struct {
	Test                    string `mapstructure:"TEST"`
	FirestoreProjectName    string `mapstructure:"FIRESTORE_PROJECTNAME"`
	FirestoreCollectionName string `mapstructure:"FIRESTORE_COLLECTIONNAME"`
	FirestoreJson           string `mapstructure:"FIRESTORE_JSON"`
}

var Cfg cfg

func init() {
	viper.SetConfigFile(".env")

	viper.ReadInConfig()

	viper.Unmarshal(&Cfg)
}
