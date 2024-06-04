package postgres

import (
	"database/sql"
	"fmt"
	"root/config"
	"root/storage"

	_ "github.com/lib/pq"
)

type Storage struct {
	Db      *sql.DB
	Products storage.Product
	Users    storage.User
	Baskets  storage.Election
}

func NewPostgresStorage() (storage.InitRoor, error) {
	config:=config.Load()
	con := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
						config.PostgresUser, config.PostgresPassword, 
						config.PostgresHost, config.PostgresPort, 
						config.PostgresDatabase)
	db, err := sql.Open("postgres", con)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &Storage{Db: db}, err

}

func (s *Storage) Product() storage.Product {
	if s.Products == nil {
		s.Products = &ProductStorage{s.Db}
	}
	return s.Products
}

func (s *Storage) User() storage.User {
	if s.Users == nil {
		s.Users =&UserStorage{s.Db}
	}
	return s.Users
}



func (s *Storage) Basket() storage.Election {
	if s.Baskets == nil {
		s.Baskets = &ElectionStorage{s.Db}
	}
	return s.Baskets
}


