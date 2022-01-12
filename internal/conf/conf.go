package conf

import (
	"os"

	log "github.com/sirupsen/logrus"
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func init() {
	viper.SetEnvPrefix("USULROSTER")

	flag.StringP(
		"auth-key",
		"k",
		"",
		"Secret key for oauth2 JWT encoding (required)",
	)
	viper.BindPFlag("auth_key", flag.Lookup("auth-key"))
	viper.BindEnv("auth_key")

	flag.StringP(
		"auth-key-alg",
		"a",
		"HS256",
		"Key algorithm for oaut2 JWT encoding",
	)
	viper.BindPFlag("auth_key_alg", flag.Lookup("auth-key-alg"))
	viper.BindEnv("auth_key_alg")

	flag.StringP(
		"mongo-uri",
		"m",
		"mongodb://localhost:27017",
		"MongoDB connection URI (required)",
	)
	viper.BindPFlag("mongo_uri", flag.Lookup("mongo-uri"))
	viper.BindEnv("mongo_uri")

	flag.StringP(
		"port",
		"p",
		"80",
		"Port to expose API on",
	)
	viper.BindPFlag("port", flag.Lookup("port"))
	viper.BindEnv("port")
}

func LoadEnvs() {
	flag.Parse()

	if viper.GetString("mongo_uri") == "" {
		log.Error("No mongo uri set: usulroster -m")
		os.Exit(2)
	}

	if viper.GetString("auth_key") == "" {
		log.Error("No oauth2 key set: usulroster -k")
		os.Exit(2)
	}
}
