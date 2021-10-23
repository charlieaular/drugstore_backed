package model

import "database/sql"

type Medicamento struct {
	ID        int32           `gorm:"primary_key;column:id;type:INT4;"`
	Nombre    sql.NullString  `gorm:"column:nombre;type:VARCHAR;size:50;"`
	Precio    sql.NullFloat64 `gorm:"column:precio;type:FLOAT8;"`
	Ubicacion sql.NullString  `gorm:"column:ubicacion;type:VARCHAR;size:50;"`
}

func (m *Medicamento) TableName() string {
	return "medicamento"
}
