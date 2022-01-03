package models

type Levels struct {
	ID   uint32 `json:"id" db:"lvl_id"`
	Name string `json:"name" validate:"required,lte=30" filter:"lk" db:"lvl_name"`
}

func (l *Levels) Validators() error {
	return translateError(validate.Struct(l))
}
