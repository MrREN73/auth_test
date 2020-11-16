package service

import (
	"auth/src/common"
)

type Service struct {
	db common.GormDB
}

func New(db common.GormDB) *Service {
	return &Service{db: db}
}

func (s *Service) GetDB() *common.GormDB {
	return &s.db
}
