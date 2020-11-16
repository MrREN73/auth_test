package models

import (
	"auth/src/common"
	models "auth/src/models/role"
	"crypto/sha1"
	"encoding/base64"
	"time"
)

type User struct {
	ID common.BaseID `gorm:"primary_key" json:"id" omit:"update"`
	Auth
	TimeStamps
	CommonInfo
}

type TimeStamps struct {
	CreatedAt time.Time `json:"-" inner:"-" omit:"update"`
	UpdatedAt time.Time `json:"-" inner:"-" omit:"update"`
	DeletedAt time.Time `json:"-" inner:"-" omit:"update"`
}

type CommonInfo struct {
	FirstName  string      `json:"firstName" binding:"required"`
	LastName   string      `json:"lastName" binding:"required"`
	MiddleName string      `json:"middleName"`
	Email      string      `json:"email" sql:"unique;index" binding:"required"`
	Secret     string      `json:"-"  sql:"unique" binding:"required"`
	Active     bool        `json:"active"`
	Role       models.Role `json:"role"`
}

type Auth struct {
	ConfirmToken string
	Confirmed    bool
}

func (c *CommonInfo) Hash() {
	hasher := sha1.New()
	hasher.Write([]byte(c.Secret))
	c.Secret = base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}
