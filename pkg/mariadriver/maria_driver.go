package mariadriver

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const DefaultPort = "3306"

// MariaDBDriver is the interface
type MariaDBDriver interface {
	Connect() *sqlx.DB
}

// Config is a model for connect MariaDB
type Config struct {
	User         string
	Pass         string
	Host         string
	DatabaseName string
	Port         string
	MaxLifetime  string
	MaxIdleConns string
	MaxOpenConns string
}

type mariaDB struct {
	Conf Config
}

func (db *mariaDB) Connect() *sqlx.DB {
	dsName := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=true", db.Conf.User, db.Conf.Pass, db.Conf.Host, db.Conf.Port, db.Conf.DatabaseName)
	conn, err := sqlx.Connect("mysql", dsName)
	maxOpenConns, _ := strconv.Atoi(db.Conf.MaxOpenConns)
	maxIdleConns, _ := strconv.Atoi(db.Conf.MaxIdleConns)
	maxLifetime, _ := strconv.Atoi(db.Conf.MaxLifetime)
	if maxOpenConns > 0 {
		conn.SetMaxOpenConns(maxOpenConns) // The default is 0 (unlimited), ex: 1000
	}
	if maxIdleConns > 0 {
		conn.SetMaxIdleConns(maxIdleConns) // defaultMaxIdleConns = 2, ex: 10
	}
	conn.SetConnMaxLifetime(time.Duration(maxLifetime)) // 0, Connections are reused forever
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("MySQL Connected!")
	}
	return conn
}

// New for create mariadb driver
func New(config Config) MariaDBDriver {
	return &mariaDB{
		Conf: config,
	}
}

// ConfigEnv for create mariadb driver
func ConfigEnv() Config {
	return Config{
		User:         os.Getenv("MARIA_USER"),
		Pass:         os.Getenv("MARIA_PASS"),
		Host:         os.Getenv("MARIA_HOST"),
		DatabaseName: os.Getenv("MARIA_DATABASE"),
		Port:         os.Getenv("MARIA_PORT"),
		MaxLifetime:  os.Getenv("MARIA_MAX_LIFETIME"),
		MaxIdleConns: os.Getenv("MARIA_MAX_IDLE_CONNS"),
		MaxOpenConns: os.Getenv("MARIA_MAX_OPEN_CONNS"),
	}
}
