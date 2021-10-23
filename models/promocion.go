package model

import (
	"database/sql"
	"time"
)

type Promocion struct {
	ID          int32           `gorm:"primary_key;column:id;type:INT4;"`
	Descripcion sql.NullString  `gorm:"column:descripcion;type:VARCHAR;size:100;"`
	Porcentaje  sql.NullFloat64 `gorm:"column:porcentaje;type:FLOAT8;"`
	FechaInicio time.Time       `gorm:"column:fecha_inicio;type:DATE;"`
	FechaFin    time.Time       `gorm:"column:fecha_fin;type:DATE;"`
}

func (p *Promocion) TableName() string {
	return "promocion"
}
