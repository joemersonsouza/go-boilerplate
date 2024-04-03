package repository

import "go.uber.org/dig"

type ICompanyRepository interface {
	GetCompany(apiKey string) CompanyObject
}

type CompanyRepository struct {
	database IDatabase
}

type CompanyRepositoryDependencies struct {
	dig.In
	Database IDatabase `name:"Database"`
}

type CompanyObject struct {
	Id     string
	ApiKey string
}

func CompanyRepositoryInstance(deps CompanyRepositoryDependencies) *CompanyRepository {
	return &CompanyRepository{
		database: deps.Database,
	}
}

func (instance *CompanyRepository) GetCompany(apiKey string) CompanyObject {
	Db := instance.database.ConnectDatabase()
	var company CompanyObject

	rows, err := Db.Query("select * from company where api_key = $1", apiKey)

	if err != nil {
		return company
	}

	rows.Next()
	rows.Scan(&company.Id, &company.ApiKey)

	return company
}
