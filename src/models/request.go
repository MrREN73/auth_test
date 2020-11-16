package models

import (
	"auth/src/common"

	"github.com/jinzhu/gorm"
)

const limit = 10

type RequestUser struct {
	CommonInfo
	PageRequest
}

type PageRequest struct {
	Limit  common.BaseID `json:"limit"`
	Offset common.BaseID `json:"offset"`
}

func (p PageRequest) Append(q *gorm.DB) *gorm.DB {
	if p.Limit == 0 {
		return q.Offset(p.Offset).Limit(limit)
	}

	return q.Offset(p.Offset).Limit(p.Limit)
}
