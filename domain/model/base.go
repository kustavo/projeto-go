package model

import "time"

type Entity interface {
	SetId(id uint64)
}

type Base struct {
	ID        uint64    `json:"id"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
