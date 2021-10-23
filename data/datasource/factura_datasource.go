package datasource

type FacturaDataSource interface {
	GetAll() string
}

type FacturaDataSourceImpl struct {
	facturaDataSource FacturaDataSource
}

func NewFacturaDataSourceImpl() FacturaDataSource {
	return FacturaDataSourceImpl{}
}

func (impl FacturaDataSourceImpl) GetAll() string {
	return "FacturaDataSourceImpl"
}
