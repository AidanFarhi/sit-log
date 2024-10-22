package repository

import "github.com/AidanFarhi/sitlog/model"

type AdultRepository interface {
	GetAdult(id int) (model.Adult, error)
	CreateAdult(adult model.Adult) error
	UpdateAdult(adult model.Adult) error
}
