package test

import (
	repository "main/repositories"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCompanyApiKey = "1bbd59c5-d75a-4729-93dc-deb35b16547c"

func TestGetCompanyWrongId(t *testing.T) {
	database := repository.DatabaseInstance()
	companyRepository := repository.CompanyRepositoryInstance(repository.CompanyRepositoryDependencies{Database: database})

	result := companyRepository.GetCompany("My-Company")

	assert.Equal(t, result.Id, "")
	assert.Equal(t, result.ApiKey, "")
}

func TestGetCompanyAndGetSuccessful(t *testing.T) {
	database := repository.DatabaseInstance()
	companyRepository := repository.CompanyRepositoryInstance(repository.CompanyRepositoryDependencies{Database: database})

	result := companyRepository.GetCompany(testCompanyApiKey)

	assert.Equal(t, result.ApiKey, testCompanyApiKey)
}
