package model

type Medicamento struct {
	ID        int32   `gorm:"primary_key;column:id;type:INT4;" json:"id"`
	Nombre    string  `gorm:"column:nombre;type:VARCHAR;size:50;" json:"nombre"`
	Precio    float64 `gorm:"column:precio;type:FLOAT8;" json:"precio"`
	Ubicacion string  `gorm:"column:ubicacion;type:VARCHAR;size:50;" json:"ubicacion"`
}

func (m *Medicamento) TableName() string {
	return "medicamento"
}
