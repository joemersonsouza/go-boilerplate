package repository

type CompanyObject struct {
	Id     string
	ApiKey string
}

func GetCompany(apiKey string) CompanyObject {
	ConnectDatabase()
	var company CompanyObject

	rows, err := Db.Query("select * from company where api_key = $1", apiKey)

	if err != nil {
		return company
	}

	rows.Next()
	rows.Scan(&company.Id, &company.ApiKey)

	return company
}
