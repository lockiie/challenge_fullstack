package models

import (
	"errors"
	"time"
)

type Developers struct {
	ID     uint32    `json:"id" db:"dev_id"`
	Name   string    `json:"name" validate:"required,lte=50" filter:"lk" db:"dev_name"`
	Gender string    `json:"gender" validate:"oneof=M F" filter:"eq" db:"dev_gender"`
	Birth  time.Time `json:"birth" validate:"required" filter:"bw" db:"dev_birth"`
	Hobby  string    `json:"hobby" validate:"lte=200" filter:"lk" db:"dev_hobby"`
	Levels Levels    `json:"level" validate:"required,nostructlevel"`
}

func (d *Developers) Validators() error {
	if d.Levels.ID == 0 {
		return errors.New("level.ID is required field")
	}
	return translateError(validate.Struct(d))
}
