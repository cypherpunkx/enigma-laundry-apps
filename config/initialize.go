package config

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

var (
	Cfg      *Config
	TokenCfg TokenConfig
)

func InitiliazeConfig() {

	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&Cfg); err != nil {
		panic(err)
	}

	appTokenExpire, err := strconv.Atoi(Cfg.Token.Expire)

	if err != nil {
		panic(err)
	}

	// accessTokenLifeTime := time.Duration(Cfg.Token.TokenExpire) * time.Minute

	TokenCfg.ApplicationName = Cfg.Token.Name
	TokenCfg.JWTSignatureKey = []byte(Cfg.Token.Key)
	TokenCfg.JWTSigningMethod = jwt.SigningMethodHS256
	TokenCfg.AccessTokenLifeTime = time.Duration(appTokenExpire) * time.Minute

}
