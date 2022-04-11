package models

import (
	"time"

	"github.com/satori/go.uuid"
	"github.com/uptrace/bun"
)

type Sex string

const (
	Male   Sex = "male"
	Female Sex = "female"
	Other  Sex = "other"
)

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	UUID       uuid.NullUUID `bun:",type:uuid,default:uuid_generate_v4()"`
	FirstName  string
	LastName   string
	MiddleName string
	FullName   string
	Sex        Sex
	BirthDate  time.Time
}
