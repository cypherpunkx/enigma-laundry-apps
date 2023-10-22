package config

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Server struct {
	Port int `mapstructure:"port"`
}

type Database struct {
	Host     string `mapstructure:"host"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Port     int    `mapstructure:"port"`
	Dbname   string `mapstructure:"dbname"`
	Driver   string `mapstructure:"driver"`
}

type FileConfig struct {
	FilePath string `mapstructure:"filepath"`
}

type TokenConfig struct {
	ApplicationName     string
	JWTSignatureKey     []byte
	JWTSigningMethod    *jwt.SigningMethodHMAC
	AccessTokenLifeTime time.Duration
}

type Token struct {
	Name   string `mapstructure:"name"`
	Key    string `mapstructure:"key"`
	Expire string `mapstructure:"expire"`
}

type Config struct {
	Server     `mapstructure:"server"`
	Database   `mapstructure:"database"`
	FileConfig `mapstructure:"fileconfig"`
	Token      `mapstructure:"token"`
}
