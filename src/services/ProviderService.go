package services

import (
	"golangtest/framework/commons"
	"golangtest/framework/helpers"
	"golangtest/src/models"
	"golangtest/src/repositories"

	"gorm.io/gorm"
)

type providerService struct {
	providerRepository repositories.ProviderRepository
}

func (service *providerService) AddProvider(provider *models.Providers) commons.Response {
	var response commons.Response
	provider.NoHp = helpers.EncryptAESCBC(provider.NoHp)
	if err := service.providerRepository.AddProvider(*provider); err != nil {
		response.Status = 500
		response.Message = "Failed to add provider"
	} else {
		response.Status = 201
		response.Message = "Success to add provider"
	}
	return response
}

func (service *providerService) GetProviders() commons.Response {
	var noHpData []string
	var response commons.Response
	data, err := service.providerRepository.GetProviders()
	
	for _, item := range data {
		noHpData = append(noHpData, helpers.DecryptAESCBC(item.NoHp))
	}
	result, error := helpers.GetGanjilGenapNumber(noHpData)
	if error != nil {
		response.Status = 500
		response.Message = "Failed to get ganjil genap"
	}

	if err != nil {
		response.Status = 500
		response.Message = "Failed to get providers"
	} else {
		response.Status = 201
		response.Message = "Success to get providers"
		response.Data = result
	}
	return response
}

func (service *providerService) GeneratePhoneNumber() commons.Response {
	var response commons.Response
	response.Status = 201
	response.Message = "Success to generate phone number"
	response.Data = helpers.GeneratePhoneNumber()
	return response
}

type ProviderService interface {
	AddProvider(provider *models.Providers) commons.Response	
	GetProviders() commons.Response
	GeneratePhoneNumber() commons.Response
}

func NewProviderService(Conn *gorm.DB) ProviderService {
	return &providerService{providerRepository: repositories.NewProviderRepository(Conn)}
}
