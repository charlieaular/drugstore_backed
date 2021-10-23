package model

import (
	"github.com/charlieaular/drugstore_backend/utils"
)

type Promocion struct {
	ID          int32            `gorm:"primary_key;column:id;type:INT4;" json:"id"`
	Descripcion string           `gorm:"column:descripcion;type:VARCHAR;size:100;" json:"descripcion" binding:"required"`
	Porcentaje  float64          `gorm:"column:porcentaje;type:FLOAT8;" json:"porcentaje" binding:"required,lte=70"`
	FechaInicio utils.CustomDate `gorm:"column:fecha_inicio;type:DATE;" json:"fecha_inicio,string" time_format:"2006-01-02" time_utc:"1"`
	FechaFin    utils.CustomDate `gorm:"column:fecha_fin;type:DATE;" json:"fecha_fin,string" time_format:"2006-01-02" time_utc:"1"`
}

func (p *Promocion) TableName() string {
	return "promocion"
}
