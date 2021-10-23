package datasource

type MedicamentoDataSource interface {
	GetAll() string
}

type MedicamentoDataSourceImpl struct {
	MedicamentoDataSource MedicamentoDataSource
}

func NewMedicamentoDataSourceImpl() MedicamentoDataSource {
	return MedicamentoDataSourceImpl{}
}

func (impl MedicamentoDataSourceImpl) GetAll() string {
	return "MedicamentoDataSourceImpl"
}
