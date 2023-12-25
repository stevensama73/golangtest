package repositories

import (
	"golangtest/src/models"
	"gorm.io/gorm"
)

type dbItem struct {
	Conn *gorm.DB
}

func (db *dbItem) AddProvider(provider models.Providers) error {
	return db.Conn.Create(provider).Error
}

func (db *dbItem) GetProviders() ([]models.Providers, error) {
	var data []models.Providers
	result := db.Conn.Find(&data)
	return data, result.Error
}

type ProviderRepository interface {
	AddProvider(provider models.Providers) error
	GetProviders() ([]models.Providers, error)
}

func NewProviderRepository(Conn *gorm.DB) ProviderRepository {
	return &dbItem{Conn: Conn}
}
