package db

//DataConfig is a configuration options to get acces to Data Base
type DataConfig struct {
	Driver     string
	DataSource string
}

//NewDataConfig is a DataConfig constructor
func NewDataConfig(driver string, dataSource string) *DataConfig {
	return &DataConfig{
		driver,
		dataSource,
	}
}
