package repositories

import (
	"iot-golang/internal/ina/models"

	"gorm.io/gorm"
)

type dbIna struct {
	Conn *gorm.DB
}

// Create implements InaRepository.
func (db *dbIna) Create(ina models.Ina) error {
	return db.Conn.Create(&ina).Error
}

// Delete implements InaRepository.
func (db *dbIna) Delete(Id int64) error {
	return db.Conn.Delete(&models.Ina{Id: Id}).Error
}

// GetAll implements InaRepository.
func (db *dbIna) GetAll() ([]models.Ina, error) {
	var data []models.Ina
	result := db.Conn.Find(&data)
	return data, result.Error
}

// GetById implements InaRepository.
func (db *dbIna) GetById(Id int64) (models.Ina, error) {
	var data models.Ina
	result := db.Conn.Where("id", Id).First(&data)
	return data, result.Error
}

// GetByToken implements InaRepository.
func (db *dbIna) GetByToken(DeviceToken string) ([]models.Ina, error) {
	var data []models.Ina
	result := db.Conn.Where("device_token", DeviceToken).Find(&data)
	return data, result.Error
}

// Update implements InaRepository.
func (db *dbIna) Update(Id int64, ina models.Ina) error {
	return db.Conn.Where("id", Id).Updates(ina).Error
}

type InaRepository interface {
	Create(ina models.Ina) error
	Update(Id int64, ina models.Ina) error
	Delete(Id int64) error
	GetById(Id int64) (models.Ina, error)
	GetByToken(DeviceToken string) ([]models.Ina, error)
	GetAll() ([]models.Ina, error)
}

func NewInaRepository(Conn *gorm.DB) InaRepository {
	return &dbIna{Conn: Conn}
}
