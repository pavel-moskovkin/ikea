package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Sex int

const (
	Male Sex = iota
	Female
	Other
)

func (s Sex) String() string {
	return []string{"male", "female", "other"}[s]
}

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	UUID       uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()"`
	FirstName  string
	LastName   string
	MiddleName string
	FullName   string
	Sex        Sex
	BirthDate  time.Time
}
