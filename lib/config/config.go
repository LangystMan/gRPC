package config

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/rs/zerolog/log"
	"gopkg.in/ini.v1"
)

func init() {
	log.Logger = LoadLogger()
}

type Setup struct {
	cfg  *ini.File
	gorm *gorm.DB
}

type Sql struct {
	Type     string `ini:"type"`
	Host     string `ini:"host"`
	Port     string `ini:"port"`
	User     string `ini:"user"`
	Password string `ini:"password"`
	Database string `ini:"database"`
}

func LoadSetup(file string) (*Setup, error) {

	set := Setup{}
	err := set.LoadConfig(file)
	if err != nil {
		return nil, err
	}

	return &set, nil
}

func (obj *Setup) LoadConfig(file string) error {

	var err error
	if obj.cfg != nil {
		return errors.New("configuration file already initialized")
	}

	obj.cfg, err = ini.Load(file)
	if err != nil {
		return err
	}

	log.Info().Msg("SUCCESS LOAD CONFIGURATION FILE")

	return nil
}

func (obj *Setup) GetAddress() (string, error) {

	host, err := obj.cfg.Section("api").GetKey("host")
	if err != nil {
		return "", err
	}

	port, err := obj.cfg.Section("api").GetKey("port")
	if err != nil {
		return "", err
	}

	return host.String() + ":" + port.String(), nil
}

func (obj *Setup) ConnectGORM() error {

	if obj.gorm == nil {

		sql := Sql{}
		err := obj.cfg.Section("sql").MapTo(&sql)
		if err != nil {
			return err
		}

		if len(sql.Type) == 0 {
			return errors.New("missing config parameter sql/type")
		}

		var connect string
		switch sql.Type {
		case "postgres":
			if len(sql.Host) == 0 {
				return errors.New("missing config parameter sql/host")
			}

			if len(sql.Port) == 0 {
				return errors.New("missing config parameter sql/host")
			}

			if len(sql.User) == 0 {
				return errors.New("missing config parameter sql/user")
			}

			if len(sql.Database) == 0 {
				return errors.New("missing config parameter sql/database")
			}

			if len(sql.Password) == 0 {
				return errors.New("missing config parameter sql/password")
			}

			connect = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", sql.Host, sql.Port, sql.User, sql.Database, sql.Password)
		default:
			return errors.New("unknown database type parameter sql/type")
		}

		obj.gorm, err = gorm.Open(sql.Type, connect)
		if err != nil || obj.gorm == nil {
			return errors.New("unable to connect to database")
		}

		if err = obj.gorm.DB().Ping(); err != nil {
			return errors.New("unable to ping to database")
		}

		obj.gorm.SingularTable(true)

		log.Info().Msg("SUCCESS CONNECTED TO DATABASE")

	}

	return nil
}
