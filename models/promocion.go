package model

import (
	"time"
)

type Promocion struct {
	ID          int32     `gorm:"primary_key;column:id;type:INT4;" json:"id"`
	Descripcion string    `gorm:"column:descripcion;type:VARCHAR;size:100;" json:"descripcion"`
	Porcentaje  float64   `gorm:"column:porcentaje;type:FLOAT8;" json:"porcentaje"`
	FechaInicio time.Time `gorm:"column:fecha_inicio;type:DATE;" json:"fecha_inicio"`
	FechaFin    time.Time `gorm:"column:fecha_fin;type:DATE;" json:"fecha_fin"`
}

func (p *Promocion) TableName() string {
	return "promocion"
}
