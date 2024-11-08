package feedcast_database

import (
	"cmp"
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	gorm_mysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log/slog"
	"net/url"
	"os"
	"strings"
	"time"
)

var ErrInvalidDsnProtocol = errors.New("dsn-invalid-scheme")

type Connection struct {
	db   *sqlx.DB
	Gorm *gorm.DB
}

func (c *Connection) Close() error {
	return c.db.Close()
}

func WithDbDsn(dsn string) error {
	u, e := url.Parse(dsn)

	if e == nil {
		if "mysql" != u.Scheme {
			return ErrInvalidDsnProtocol
		}

		pwd, _ := u.User.Password()

		os.Setenv("DB_HOST", u.Host)
		os.Setenv("DB_DATABASE", strings.TrimLeft(u.Path, "/"))
		os.Setenv("DB_USER", u.User.Username())
		os.Setenv("DB_PASSWORD", pwd)
	}

	return e
}

func GetConnection() *Connection {
	timeout := 60 * time.Second
	cfg := mysql.Config{
		User:                 cmp.Or(os.Getenv("DB_USER"), "feedcast"),
		Passwd:               cmp.Or(os.Getenv("DB_PASSWORD"), "feedcast"),
		Net:                  "tcp",
		Addr:                 cmp.Or(os.Getenv("DB_HOST"), "localhost"),
		DBName:               cmp.Or(os.Getenv("DB_DATABASE"), "feedcast"),
		AllowNativePasswords: true,
		ReadTimeout:          timeout,
		Timeout:              timeout,
	}

	params := url.Values{
		"parseTime": []string{"true"},
		"loc":       []string{cmp.Or(os.Getenv("DB_TIMEZONE"), "Europe/Paris")},
	}

	dsnConnect := fmt.Sprintf("%s&%s", cfg.FormatDSN(), params.Encode())
	db, err := sqlx.Connect("mysql", dsnConnect)

	if nil != err {
		slog.Error("Database connection error", "mess", err.Error())
		os.Exit(1)
	}

	g, _ := gorm.Open(
		gorm_mysql.New(gorm_mysql.Config{
			Conn: db,
		}),
		&gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
			SkipDefaultTransaction: true,
			Logger:                 logger.Default.LogMode(logger.Error),
		},
	)

	return &Connection{
		db:   db,
		Gorm: g,
	}
}
