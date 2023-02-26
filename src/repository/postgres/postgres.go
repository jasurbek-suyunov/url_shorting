package postgres

import (
	"fmt"
	"log"
	"time"

	"github.com/SuyunovJasurbek/url_shorting/config"
	"github.com/SuyunovJasurbek/url_shorting/src/repository"
	"github.com/jmoiron/sqlx"
)

type Storage struct {
	*sqlx.DB
	user repository.UserI
	url  repository.UrlI
}

// Token implements repository.StorageI
func (*Storage) Token() repository.TokenI {
	panic("unimplemented")
}

func NewPostgres(cfg config.Config) (repository.StorageI, error) {
	psqlConnString := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresDatabase,
		cfg.PostgresPassword,
	)
	db, err := sqlx.Open("postgres", psqlConnString)
	if err != nil {
		log.Fatalf("cannot connect to postgresql db: %s", err.Error())
	}

	// setting some settings for db
	db.SetConnMaxLifetime(time.Duration(time.Duration(cfg.PostgresMaxConnections).Minutes()))
	db.SetMaxOpenConns(cfg.PostgresMaxConnections)

	// checking for Ping&Pong
	if err := db.Ping(); err != nil {
		log.Fatalf("ping error %s", err.Error())
	}

	return &Storage{
		DB: db,
	}, nil
}

func (s *Storage) User() repository.UserI {
	if s.user == nil {
		s.user = NewUserRepo(s.DB)
	}
	return s.user
}

func (s *Storage) Url() repository.UrlI {
	if s.url == nil {
		s.url = NewUrlRepo(s.DB)
	}
	return s.url
}
