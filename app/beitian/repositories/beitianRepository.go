package repositories

import (
	"iot-golang/app/beitian/models"

	"gorm.io/gorm"
)

type dbBeitian struct {
	Conn *gorm.DB
}

// GetNewByToken implements BeitianRepository.
func (db *dbBeitian) GetNewByToken(DeviceToken string) ([]models.Beitian, error) {
	var data []models.Beitian
	result := db.Conn.Where("device_token = ?", DeviceToken).Order("created_at desc").Limit(1).Find(&data)
	return data, result.Error
}

// Create implements BeitianRepository.
func (db *dbBeitian) Create(beitian models.Beitian) error {
	return db.Conn.Create(&beitian).Error
}

// Delete implements BeitianRepository.
func (db *dbBeitian) Delete(Id int64) error {
	return db.Conn.Delete(&models.Beitian{Id: Id}).Error
}

// GetAll implements BeitianRepository.
func (db *dbBeitian) GetAll() ([]models.Beitian, error) {
	var data []models.Beitian
	result := db.Conn.Find(&data)
	return data, result.Error
}

// GetById implements BeitianRepository.
func (db *dbBeitian) GetById(Id int64) (models.Beitian, error) {
	var data models.Beitian
	result := db.Conn.Where("id", Id).First(&data)
	return data, result.Error
}

// GetByToken implements BeitianRepository.
func (db *dbBeitian) GetByToken(DeviceToken string) ([]models.Beitian, error) {
	var data []models.Beitian
	result := db.Conn.Where("device_token", DeviceToken).Find(&data)
	return data, result.Error
}

// Update implements BeitianRepository.
func (db *dbBeitian) Update(Id int64, beitian models.Beitian) error {
	return db.Conn.Where("id", Id).Updates(beitian).Error
}

type BeitianRepository interface {
	Create(beitian models.Beitian) error
	Update(Id int64, beitian models.Beitian) error
	Delete(Id int64) error
	GetById(Id int64) (models.Beitian, error)
	GetByToken(DeviceToken string) ([]models.Beitian, error)
	GetAll() ([]models.Beitian, error)
	GetNewByToken(DeviceToken string) ([]models.Beitian, error)
}

func NewBeitianRepository(Conn *gorm.DB) BeitianRepository {
	return &dbBeitian{Conn: Conn}
}
