package model

import "time"

type URL struct {
	ID        int64      `db:"id" json:"id"`
	Code      string     `db:"code" json:"code"`
	Target    string     `db:"target" json:"target"`
	CreatedAt *time.Time `db:"created_at" json:"created_at,omitempty"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at,omitempty"`
}
