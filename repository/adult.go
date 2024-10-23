package repository

import (
	"database/sql"

	"github.com/AidanFarhi/sitlog/model"
)

type AdultRepository interface {
	CreateAdult(adult model.Adult) error
	UpdateAdult(adult model.Adult) error
}

type SimpleAdultRepository struct {
	DB *sql.DB
}

func NewSimpleAdultRepository(db *sql.DB) SimpleAdultRepository {
	return SimpleAdultRepository{DB: db}
}

func (sar SimpleAdultRepository) GetAdult(id int) (model.Adult, error) {
	adult := model.Adult{}
	query := "SELECT name, email FROM adult WHERE id = ?"
	row := sar.DB.QueryRow(query, id)
	err := row.Scan(&adult.Name, &adult.Email)
	if err != nil {
		return adult, err
	}
	return adult, nil
}

func (sar SimpleAdultRepository) CreateAdult(adult model.Adult) error {
	query := "INSERT INTO adult (name, email, password) VALUES (?, ?, ?)"
	_, err := sar.DB.Exec(query, adult.Name, adult.Email, adult.Password)
	if err != nil {
		return err
	}
	return nil
}
