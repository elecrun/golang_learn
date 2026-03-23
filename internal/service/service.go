package service

import (
	"highqps-cache/internal/db"
	"highqps-cache/internal/model"
)

type Service struct {
	db *db.DB
}

func NewService(d *db.DB) *Service {
	return &Service{db: d}
}

func (s *Service) GetItem(id int64) (*model.Item, error) {
	return s.db.GetByID(id)
}