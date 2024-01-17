package repositories

import (
	"iot-golang/internal/pzem/models"

	"gorm.io/gorm"
)

type dbPzem struct {
	Conn *gorm.DB
}

// Create implements PzemRepository.
func (db *dbPzem) Create(pzem models.Pzem) error {
	return db.Conn.Create(&pzem).Error
}

// Delete implements PzemRepository.
func (db *dbPzem) Delete(Id int64) error {
	return db.Conn.Delete(&models.Pzem{Id: Id}).Error
}

// GetAll implements PzemRepository.
func (db *dbPzem) GetAll() ([]models.Pzem, error) {
	var data []models.Pzem
	result := db.Conn.Find(&data)
	return data, result.Error
}

// GetById implements PzemRepository.
func (db *dbPzem) GetById(Id int64) (models.Pzem, error) {
	var data models.Pzem
	result := db.Conn.Where("id", Id).First(&data)
	return data, result.Error
}

// GetByToken implements PzemRepository.
func (db *dbPzem) GetByToken(DeviceToken string) ([]models.Pzem, error) {
	var data []models.Pzem
	result := db.Conn.Where("device_token", DeviceToken).Find(&data)
	return data, result.Error
}

// Update implements PzemRepository.
func (db *dbPzem) Update(Id int64, pzem models.Pzem) error {
	return db.Conn.Where("id", Id).Updates(pzem).Error
}

type PzemRepository interface {
	Create(pzem models.Pzem) error
	Update(Id int64, pzem models.Pzem) error
	Delete(Id int64) error
	GetById(Id int64) (models.Pzem, error)
	GetByToken(DeviceToken string) ([]models.Pzem, error)
	GetAll() ([]models.Pzem, error)
}

func NewPzemRepository(Conn *gorm.DB) PzemRepository {
	return &dbPzem{Conn: Conn}
}
