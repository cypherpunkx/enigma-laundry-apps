package manager

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"enigmacamp.com/enigma-laundry-apps/config"
)

type InfraManager interface {
	Conn() *sql.DB
}

type infraManager struct {
	db  *sql.DB
	cfg *config.Config
}

func (i *infraManager) initDB() error {
	var dbConf = i.cfg.Database

	dataSourceName := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", dbConf.Host, dbConf.Port, dbConf.Username, dbConf.Password, dbConf.Dbname)

	db, err := sql.Open(dbConf.Driver, dataSourceName)

	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	i.db = db

	return nil
}

func (i *infraManager) Conn() *sql.DB {
	return i.db
}

func NewInfraManager(configParam *config.Config) (InfraManager, error) {
	infra := &infraManager{
		cfg: configParam,
	}

	err := infra.initDB()

	if err != nil {
		return nil, err
	}

	return infra, nil
}
